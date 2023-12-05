import * as path from 'path';
import {
  aws_lambda as lambda,
  aws_logs as logs,
  aws_iam as iam,
  Duration,
  CustomResource,
  Stack,
  custom_resources,
} from 'aws-cdk-lib';
import { PolicyStatement } from 'aws-cdk-lib/aws-iam';
import { Construct } from 'constructs';
import {
  ActionGroup,
  BedrockAgentProps,
  KnowledgeBaseAssociation,
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
  readonly actionGroups: ActionGroup[];
  readonly knowledgeBaseAssociations: KnowledgeBaseAssociation[];
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
    this.actionGroups = props.actionGroups ?? this.getUndefinedActionGroup();
    this.knowledgeBaseAssociations = props.knowledgeBaseAssociations ?? this.getUndefinedKnowledgeBase();

    this.bedrockAgentCustomResourceRole = new iam.Role(this, 'BedrockAgentCustomResourceRole', {
      assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com'),
    });

    // /**
    // * Check if IAM role was provided for knowledge base and if it was
    // * then add iam:PassRole permission on that role to custom resource's
    // * IAM role.
    // */
    // const resources = [this.agentResourceRoleArn];
    // if (this.knowledgeBase.roleArn !== 'Undefined') {
    //   resources.push(this.knowledgeBase.roleArn);
    // }

    this.bedrockAgentCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['iam:PassRole'],
      resources: [this.agentResourceRoleArn],
    }));

    this.bedrockAgentCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['*'],
      resources: ['arn:aws:bedrock:*'],
    }));

    /**
    * If agent group was provided then attach lambda:AllowInvoke policy
    * by an agent to the provided Lambda.
    */
    // this.actionGroups.actionGroupExecutor != 'Undefined' ? this.lambdaAttachResourceBasedPolicy(): 'Undefined';
    this.actionGroups.forEach(actionGroup => actionGroup.actionGroupExecutor != 'Undefined' ?
      this.lambdaAttachResourceBasedPolicy(actionGroup.actionGroupExecutor) : 'Undefined');

    /**
    * If agent group was provided but IAM role for the agent
    * was not then attach s3 read access to the bucket to the default role.
    */
    // props.actionGroups?.s3BucketName && !props.agentResourceRoleArn && this.attachS3BucketReadOnlyPolicy();
    props.actionGroups?.forEach(actionGroup =>
      actionGroup.s3BucketName && !props.agentResourceRoleArn && this.attachS3BucketReadOnlyPolicy(actionGroup.s3BucketName),
    );

    const layer = new lambda.LayerVersion(this, 'BedrockAgentLayer', {
      code: lambda.Code.fromAsset(path.join(__dirname, '../assets/lambda-layer/bedrock-agent-layer.zip')),
      compatibleRuntimes: [lambda.Runtime.PYTHON_3_10],
    });

    const onEvent = new lambda.Function(this, 'BedrockAgentCustomResourceFunction', {
      runtime: lambda.Runtime.PYTHON_3_10,
      handler: 'bedrock_agent_custom_resource.on_event',
      code: lambda.Code.fromAsset(path.join(__dirname, '../assets/functions')),
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
        ACTION_GROUPS: JSON.stringify(this.actionGroups),
        KNOWLEDGE_BASE_ASSOCIATIONS: JSON.stringify(this.knowledgeBaseAssociations),
      },
      role: this.bedrockAgentCustomResourceRole,
    });

    const bedrockAgentCustomResourceProvider = new custom_resources.Provider(this, 'BedrockAgentCustomResourceProvider', {
      onEventHandler: onEvent,
      logRetention: logs.RetentionDays.ONE_DAY,
    });

    new CustomResource(this, 'BedrockAgentCustomResource', {
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

  private attachS3BucketReadOnlyPolicy(s3BucketName: string): void {
    const agentRole = iam.Role.fromRoleArn(this, 'AgentIamRoleArn', this.agentResourceRoleArn);
    const policyStatement =
      new PolicyStatement({
        actions: ['s3:GetObject*', 's3:ListBucket'],
        resources: [`arn:aws:s3:::${s3BucketName}/*`,
          `arn:aws:s3:::${s3BucketName}`],
      });
    agentRole.addToPrincipalPolicy(policyStatement);
  }

  private lambdaAttachResourceBasedPolicy(actionGroupExecutor: string): void {
    this.bedrockAgentCustomResourceRole.addToPolicy(new iam.PolicyStatement({
      actions: ['lambda:AddPermission', 'lambda:GetFunction', 'lambda:RemovePermission'],
      resources: [actionGroupExecutor],
    }));
  }

  private getUndefinedActionGroup(): ActionGroup[] {
    return [{
      actionGroupName: 'Undefined',
      actionGroupExecutor: 'Undefined',
      s3BucketName: 'Undefined',
      s3ObjectKey: 'Undefined',
    }];
  }

  private getUndefinedKnowledgeBase(): KnowledgeBaseAssociation[] {
    return [{
      knowledgeBaseName: 'Undefined',
      instruction: 'Undefined',
    }];
  }
}

