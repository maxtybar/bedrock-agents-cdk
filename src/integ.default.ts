import {
  App,
  Stack,
} from 'aws-cdk-lib';
import { BedrockAgent } from './bedrock-agent-construct';

const app = new App();

const stack = new Stack(app, 'BedrockAgentStack');

const agentName = 'MyTestAgent';
const kbName = 'MyTestKnowledgeBase';
const agentRoleArn = 'yourAgentRoleArn';
const kbRoleArn = 'yourKnowledgeBaseRolArn';
const actionGroupName = 'MyTestActionGroup';
const dataSourceName = 'MyDataSource';
const foundationModel = 'anthropic.claude-instant-v1';
const agentInstruction = 'This is a template instruction for my agent. You were created by AWS CDK.';
const kbInstruction = 'This is a template instruction for my knowledge base. You were created by AWS CDK.';
const collectionArn = 'yourOpenSearchServerlessCollectionArn';
const vectorIndexName = 'vector-index-name';
const vectorFieldName = 'vector-field-name';
const textField = 'text-field';
const metadataField = 'metadata-field';
const opensearchStorageConfigurationType = 'OPENSEARCH_SERVERLESS';
const dataSourceType = 'S3';
const dataSourceBucketArn = 'yourDataSourceBucketArn';
const actionGroupLambdaArn = 'yourActionGroupLambdaArn';
const actionGroupS3BucketName = 'yourActionGroupApiSchemaBucketName';
const actionGroupS3ObjectKey = 'yourActionGroupApiSchemaKey';


// Amazon Bedrock Agent and Knowledge Base backed by Opensearch Serverless
new BedrockAgent(stack, 'BedrockAgentAndKbConstruct', {
  agentName: agentName,
  instruction: agentInstruction,
  foundationModel: foundationModel,
  agentResourceRoleArn: agentRoleArn,
  actionGroup: {
    actionGroupName: actionGroupName,
    actionGroupExecutor: actionGroupLambdaArn,
    s3BucketName: actionGroupS3BucketName,
    s3ObjectKey: actionGroupS3ObjectKey,
  },
  knowledgeBase: {
    name: kbName,
    roleArn: kbRoleArn,
    instruction: kbInstruction,
    storageConfiguration: {
      opensearchServerlessConfiguration: {
        collectionArn: collectionArn,
        fieldMapping: {
          metadataField: metadataField,
          textField: textField,
          vectorField: vectorFieldName,
        },
        vectorIndexName: vectorIndexName,
      },
      type: opensearchStorageConfigurationType,
    },
    dataSource: {
      name: dataSourceName,
      dataSourceConfiguration: {
        s3Configuration: {
          bucketArn: dataSourceBucketArn,
        },
        type: dataSourceType,
      },
    },
  },
});
