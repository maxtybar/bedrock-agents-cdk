import * as path from 'path';
import {
  aws_lambda as lambda,
  aws_logs as logs,
  aws_iam as iam,
  Duration,
  CustomResource,
  RemovalPolicy,
  Stack,
  custom_resources,
} from 'aws-cdk-lib';
import { Construct } from 'constructs';
import {
  BedrockKnowledgeBaseProps,
  KnowledgeBaseConfiguration,
  DataSource,
  OpenSearchServerlessStorageConfiguration,
  RedisEnterpriseCloudStorageConfiguration,
  PineconeStorageConfiguration,
} from './interfaces';

// Assign default props
const defaultProps: Partial<BedrockKnowledgeBaseProps> = {
  removalPolicy: RemovalPolicy.DESTROY,
};


// Create a BedrockAgent construct
export class BedrockKnowledgeBase extends Construct {

  readonly region: string;
  readonly name: string;
  readonly roleArn: string;
  readonly knowledgeBaseConfiguration?: KnowledgeBaseConfiguration;
  readonly storageConfiguration: OpenSearchServerlessStorageConfiguration | RedisEnterpriseCloudStorageConfiguration | PineconeStorageConfiguration;
  readonly dataSource: DataSource;
  readonly description?: string;
  readonly removalPolicy?: RemovalPolicy;
  readonly bedrockKnowledgeBaseCustomResourceRole: iam.Role;

  constructor(scope: Construct, name: string, props: BedrockKnowledgeBaseProps) {
    super(scope, name);

    props = { ...defaultProps, ...props };

    this.region = Stack.of(this).region;
    this.name = props.name;
    this.roleArn = props.roleArn;
    this.knowledgeBaseConfiguration = this.setKnowledgeBaseConfigurationDefaultValues(props.knowledgeBaseConfiguration);
    this.storageConfiguration = props.storageConfiguration;
    this.dataSource = this.setDataSourceDefaultValues(props.dataSource);
    this.description = props.description;
    this.removalPolicy = props.removalPolicy as RemovalPolicy;

    this.bedrockKnowledgeBaseCustomResourceRole = new iam.Role(this, 'BedrockKnowledgeBaseCustomResourceRole', {
      assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com'),
    });

    this.bedrockKnowledgeBaseCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['iam:PassRole'],
      resources: [this.roleArn],
    }));

    this.bedrockKnowledgeBaseCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['*'],
      resources: ['arn:aws:bedrock:*'],
    }));

    const layer = new lambda.LayerVersion(this, 'BedrockAgentLayer', {
      code: lambda.Code.fromAsset(path.join(__dirname, '../assets/lambda-layer/bedrock-agent-layer.zip')),
      compatibleRuntimes: [lambda.Runtime.PYTHON_3_10],
    });

    const onEvent = new lambda.Function(this, 'BedrockAgentCustomResourceFunction', {
      runtime: lambda.Runtime.PYTHON_3_10,
      handler: 'bedrock_knowledgebase_custom_resource.on_event',
      code: lambda.Code.fromAsset(path.join(__dirname, '../assets/functions')),
      architecture: lambda.Architecture.X86_64,
      layers: [layer],
      timeout: Duration.seconds(600),
      environment: {
        KNOWLEDGEBASE_NAME: this.name,
        ROLE_ARN: this.roleArn,
        KNOWLEDGEBASE_CONFIGURATION: JSON.stringify(this.knowledgeBaseConfiguration),
        STORAGE_CONFIGURATION: JSON.stringify(this.storageConfiguration),
        DATA_SOURCE: JSON.stringify(this.dataSource),
        DESCRIPTION: this.description ?? 'Undefined',
        REMOVAL_POLICY: this.removalPolicy.toString(),
      },
      role: this.bedrockKnowledgeBaseCustomResourceRole,
    });

    const bedrockKnowledgeBaseCustomResourceProvider = new custom_resources.Provider(this, 'BedrockKnowledgeBaseCustomResourceProvider', {
      onEventHandler: onEvent,
      logRetention: logs.RetentionDays.ONE_DAY,
    });

    new CustomResource(this, 'BedrockKnowledgeBaseCustomResource', {
      serviceToken: bedrockKnowledgeBaseCustomResourceProvider.serviceToken,
    });
  }

  private setKnowledgeBaseConfigurationDefaultValues(knowledgeBaseConfiguration?: KnowledgeBaseConfiguration): KnowledgeBaseConfiguration {
    return {
      type: knowledgeBaseConfiguration?.type ?? 'VECTOR',
      vectorKnowledgeBaseConfiguration: {
        embeddingModelArn: knowledgeBaseConfiguration?.vectorKnowledgeBaseConfiguration.embeddingModelArn ??
        `arn:aws:bedrock:${this.region}::foundation-model/amazon.titan-embed-text-v1`,
      },
    };
  }

  private setDataSourceDefaultValues(dataSource: DataSource): DataSource {
    return {
      name: dataSource.name ?? `MyDataSource-${this.name}`,
      dataSourceConfiguration: {
        s3Configuration: {
          bucketArn: dataSource.dataSourceConfiguration.s3Configuration.bucketArn,
          ...(dataSource.dataSourceConfiguration.s3Configuration.inclusionPrefixes && {
            inclusionPrefixes: dataSource.dataSourceConfiguration.s3Configuration.inclusionPrefixes,
          }),
        },
        type: dataSource.dataSourceConfiguration.type ?? 'S3',
      },
    };
  }
}

