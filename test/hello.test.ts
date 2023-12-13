import {
  App, Stack, assertions,
} from 'aws-cdk-lib';
import { BedrockAgent, BedrockKnowledgeBase } from '../src';


test('Match snapshot', () => {
  const app = new App();
  const stack = new Stack(app, 'BedrockAgentStack');
  const kbName = 'MyTestKnowledgeBase';
  const kbRoleArn = 'yourKnowledgeBaseRoleArn';
  const dataSourceName = 'MyDataSource';
  const vectorIndexName = 'my-test-index';
  const vectorFieldName = 'my-test-vector';
  const textField = 'text-field';
  const metadataField = 'metadata-field';
  const storageConfigurationType = 'OPENSEARCH_SERVERLESS';
  const dataSourceType = 'S3';
  const dataSourceBucketArn = 'yourDataSourceBucketArn';
  const opensearchServerlessCollectionArn = 'yourOpensearchServerlessCollectionArn';
  const dataSourceBucketPrefix = 'yourDataSourceBucketPrefix';

  new BedrockAgent(stack, 'BedrockAgentStack', {
    agentName: 'MyAgentTestName',
    instruction: 'This is a template instruction for my agent. You were created by AWS TypeScript CDK.',
    foundationModel: 'anthropic.claude-instant-v1',
    description: 'My custom description for an agent.',
  });

  new BedrockKnowledgeBase(stack, 'BedrockOpenSearchKnowledgeBase', {
    name: kbName,
    roleArn: kbRoleArn,
    storageConfiguration: {
      opensearchServerlessConfiguration: {
        collectionArn: opensearchServerlessCollectionArn,
        fieldMapping: {
          metadataField: metadataField,
          textField: textField,
          vectorField: vectorFieldName,
        },
        vectorIndexName: vectorIndexName,
      },
      type: storageConfigurationType,
    },
    dataSource: {
      name: dataSourceName,
      dataSourceConfiguration: {
        s3Configuration: {
          bucketArn: dataSourceBucketArn,
          inclusionPrefixes: [dataSourceBucketPrefix],
        },
        type: dataSourceType,
      },
    },
  });

  const template = assertions.Template.fromStack(stack);
  expect(template).toMatchSnapshot();
});

