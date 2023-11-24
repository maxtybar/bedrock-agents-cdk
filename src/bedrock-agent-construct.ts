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
import { PolicyStatement } from 'aws-cdk-lib/aws-iam';
import { Construct } from 'constructs';
import {
  KnowledgeBase,
  ActionGroup,
  BedrockAgentProps,
} from './interfaces';

// Assign default props
const defaultProps: Partial<BedrockAgentProps> = {
  foundationModel: 'anthropic.claude-v2',
  idleSessionTTLInSeconds: 3600,
};


// Create a BedrockAgent construct
export class BedrockAgent extends Construct {

  readonly region: string;
  readonly agentName: string;
  readonly instruction: string;
  readonly foundationModel: string;
  readonly agentResourceRoleArn: string;
  readonly description: string | undefined;
  readonly idleSessionTTLInSeconds: number;
  readonly actionGroup: ActionGroup;
  readonly knowledgeBase: KnowledgeBase;
  readonly bedrockAgentCustomResourceRole: iam.Role;

  constructor(scope: Construct, name: string, props: BedrockAgentProps) {
    super(scope, name);

    props = { ...defaultProps, ...props };

    this.region = Stack.of(this).region;
    this.agentName = props.agentName;
    this.instruction = props.instruction;
    this.foundationModel = props.foundationModel as string;
    this.agentResourceRoleArn = props.agentResourceRoleArn ?? this.getDefaultAgentResourceRoleArn();
    this.description = props.description;
    this.idleSessionTTLInSeconds = props.idleSessionTTLInSeconds as number;
    this.actionGroup = props.actionGroup ?? this.getUndefinedActionGroup();
    this.knowledgeBase = this.handleKnowledgeBaseProps(props.knowledgeBase);

    this.bedrockAgentCustomResourceRole = new iam.Role(this, 'bedrockAgentCustomResourceRole', {
      assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com'),
    });

    /**
    * Check if IAM role was provided for knowledge base and if it was
    * then add iam:PassRole permission on that role to custom resource's
    * IAM role.
    */
    const resources = [this.agentResourceRoleArn];
    if (this.knowledgeBase.roleArn !== 'Undefined') {
      resources.push(this.knowledgeBase.roleArn);
    }

    this.bedrockAgentCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['iam:PassRole'],
      resources: resources,
    }));

    this.bedrockAgentCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['*'],
      resources: ['arn:aws:bedrock:*'],
    }));

    /**
    * If agent group was provided then attach lambda:AllowInvoke policy
    * by an agent to the provided Lambda.
    */
    this.actionGroup.actionGroupExecutor != 'Undefined' ? this.lambdaAttachResourceBasedPolicy(): 'Undefined';

    /**
    * If agent group was provided but IAM role for the agent
    * was not then attach s3 read access to the bucket to the default role.
    */
    props.actionGroup?.s3BucketName && !props.agentResourceRoleArn && this.attachS3BucketReadOnlyPolicy();

    const layer = new lambda.LayerVersion(this, 'BedrockAgentLayer', {
      code: lambda.Code.fromAsset(path.join(__dirname, '../assets/lambda-layer/bedrock-agent-layer.zip')),
      compatibleRuntimes: [lambda.Runtime.PYTHON_3_10],
    });

    const onEvent = new lambda.Function(this, 'BedrockAgentCustomResourceFunction', {
      runtime: lambda.Runtime.PYTHON_3_10,
      handler: 'bedrock_agent_custom_resource.on_event',
      code: lambda.Code.fromAsset(path.join(__dirname, '../assets/function')),
      architecture: lambda.Architecture.X86_64,
      layers: [layer],
      timeout: Duration.seconds(600),
      environment: {
        AGENT_NAME: this.agentName,
        INSTRUCTION: this.instruction,
        FOUNDATION_MODEL: this.foundationModel,
        AGENT_RESOURCE_ROLE_ARN: this.agentResourceRoleArn,
        DESCRIPTION: this.description ?? 'Undefined',
        IDLE_SESSION_TTL_IN_SECONDS: this.idleSessionTTLInSeconds.toString(),
        ACTION_GROUP: JSON.stringify(this.actionGroup),
        KNOWLEDGE_BASE: JSON.stringify(this.knowledgeBase),
      },
      role: this.bedrockAgentCustomResourceRole,
    });

    const bedrockAgentCustomResourceProvider = new custom_resources.Provider(this, 'BedrockCustomResourceProvider', {
      onEventHandler: onEvent,
      logRetention: logs.RetentionDays.ONE_DAY,
    });

    new CustomResource(this, 'BedrockCustomResource', {
      serviceToken: bedrockAgentCustomResourceProvider.serviceToken,
    });
  }

  private getDefaultAgentResourceRoleArn(): string {
    return new iam.Role(this, 'AgentIamRole', {
      roleName: 'AmazonBedrockExecutionRoleForAgents_' + this.agentName,
      assumedBy: new iam.ServicePrincipal('bedrock.amazonaws.com'),
      description: 'Agent role created by CDK.',
    }).roleArn;
  }

  private attachS3BucketReadOnlyPolicy(): void {
    const agentRole = iam.Role.fromRoleArn(this, 'AgentIamRoleArn', this.agentResourceRoleArn);
    const policyStatement =
      new PolicyStatement({
        actions: ['s3:GetObject*', 's3:ListBucket'],
        resources: [`arn:aws:s3:::${this.actionGroup.s3BucketName}/*`,
          `arn:aws:s3:::${this.actionGroup.s3BucketName}`],
      });
    agentRole.addToPrincipalPolicy(policyStatement);
  }

  private lambdaAttachResourceBasedPolicy(): void {
    this.bedrockAgentCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['lambda:AddPermission', 'lambda:GetFunction', 'lambda:RemovePermission'],
      resources: [this.actionGroup.actionGroupExecutor],
    }));
  }

  private handleKnowledgeBaseProps(knowledgeBase?: KnowledgeBase): KnowledgeBase {
    if (!knowledgeBase) {
      return this.getUndefinedKnowledgeBase();
    } else {
      return this.setKnowledgeBaseDefaultValues(knowledgeBase);
    }
  }

  private getUndefinedActionGroup(): ActionGroup {
    return {
      actionGroupName: 'Undefined',
      actionGroupExecutor: 'Undefined',
      s3BucketName: 'Undefined',
      s3ObjectKey: 'Undefined',
    };
  }

  private getUndefinedKnowledgeBase(): KnowledgeBase {
    return {
      name: 'Undefined',
      roleArn: 'Undefined',
      instruction: 'Undefined',
      knowledgeBaseConfiguration: {
        type: 'Undefined',
        vectorKnowledgeBaseConfiguration: {
          embeddingModelArn: 'Undefined',
        },
      },
      storageConfiguration: {
        opensearchServerlessConfiguration: {
          vectorIndexName: 'Undefined',
          collectionArn: 'Undefined',
          fieldMapping: {
            vectorField: 'Undefined',
            metadataField: 'Undefined',
            textField: 'Undefined',
          },
        },
        type: 'OPENSEARCH_SERVERLESS',
      },
      dataSource: {
        dataSourceConfiguration: {
          s3Configuration: {
            bucketArn: 'Undefined',
          },
        },
      },
    };
  }

  private setKnowledgeBaseDefaultValues(knowledgeBase: KnowledgeBase): KnowledgeBase {
    return {
      name: knowledgeBase.name,
      roleArn: knowledgeBase.roleArn,
      instruction: knowledgeBase.instruction,
      knowledgeBaseConfiguration: {
        type: knowledgeBase.knowledgeBaseConfiguration?.type ?? 'VECTOR',
        vectorKnowledgeBaseConfiguration: {
          embeddingModelArn: knowledgeBase.knowledgeBaseConfiguration?.vectorKnowledgeBaseConfiguration.embeddingModelArn ??
          `arn:aws:bedrock:${this.region}::foundation-model/amazon.titan-embed-text-v1`,
        },
      },
      storageConfiguration: knowledgeBase.storageConfiguration,
      dataSource: {
        name: knowledgeBase.dataSource.name ?? `MyDataSource-${this.agentName}`,
        dataSourceConfiguration: {
          s3Configuration: {
            bucketArn: knowledgeBase.dataSource.dataSourceConfiguration.s3Configuration.bucketArn,
          },
          type: knowledgeBase.dataSource.dataSourceConfiguration.type ?? 'S3',
        },
      },
      removalPolicy: knowledgeBase.removalPolicy ?? RemovalPolicy.DESTROY,
    };
  }
}

