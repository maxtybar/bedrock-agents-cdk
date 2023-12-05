# Bedrock Agent and Bedrock Knowledge Base Constructs

See [API.md](API.md) for more information about construct.

Also see TypeScript deployment examples below. All of the examples assume you have appropriate IAM permissions for provisioning resources in AWS.

# Example - deploy agent only

This example will create an agent without Action Group and with default IAM role. If ``agentResourceRoleArn`` is not specified a default IAM role with no policies attached to it will be provisioned for your agent.

``` typescript
import * as cdk from 'aws-cdk-lib';
import { BedrockAgent } from 'bedrock-agents-cdk';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'BedrockAgentStack');

new BedrockAgent(stack, "BedrockAgent", {
  agentName: "BedrockAgent",
  instruction: "This is a test instruction. You were built by AWS CDK construct to answer all questions.",
})
```

# Example - deploy knowledge base only

This example will create a knowledge base backed by Amazon OpenSearch Serverless collection that could be used without an agent. This example assumes you have already created Amazon OpenSearch Serverless collection (`collectionArn`).

**Note: The IAM role creation in this example is for illustration only. Always provion IAM roles with the least required priviliges.**

``` typescript
import * as cdk from 'aws-cdk-lib';
import { BedrockKnowledgeBase } from 'bedrock-agents-cdk';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'BedrockKnowledgeBaseStack');

const kbName = 'MyTestKnowledgeBase';
const dataSourceName = 'MyDataSource';
const collectionArn = 'arn:aws:aoss:yourCollectionRegion:yourAWSAccountId:collection/yourCollectionId';
const vectorIndexName = 'my-test-index';
const vectorFieldName = 'my-test-vector';
const textField = 'text-field';
const metadataField = 'metadata-field';
const storageConfigurationType = 'OPENSEARCH_SERVERLESS';
const dataSourceType = 'S3';
const dataSourceBucketArn = 'yourDataSourceBucketArn';

// Bedrock Knowledge Base IAM role
const kbRoleArn = new iam.Role(stack, 'BedrockKnowledgeBaseRole', {
  roleName: 'AmazonBedrockExecutionRoleForKnowledgeBase_kb_test',
  assumedBy: new iam.ServicePrincipal('bedrock.amazonaws.com'),
  managedPolicies: [iam.ManagedPolicy.fromAwsManagedPolicyName('AdministratorAccess')],
}).roleArn;

// Create Bedrock Knowledge Base backed by OpenSearch Servereless
new BedrockKnowledgeBase(stack, 'BedrockOpenSearchKnowledgeBase', {
  name: kbName,
  roleArn: kbRoleArn,
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
    type: storageConfigurationType,
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
});
```

# Example - deploy agent with a single action group

This example will create an agent with an Action Group and with your IAM role (`agentResourceRoleArn`). It assumes that you already have an S3 bucket and a stored JSON or yml Open API schema file that will be included in your action group. Additionaly, `pathToLambdaFile` should contain path to your function code file inside your cdk project that you want to be attached to your agent's action group. Resource-based policy statement will be attached to your Lambda function allowing Bedrock Agent to invoke it.

**Note: The IAM role creation in this example is for illustration only. Always provion IAM roles with the least required priviliges.**

```typescript
import * as path from 'path';
import * as cdk from 'aws-cdk-lib';
import { BedrockAgent } from 'bedrock-agents-cdk';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'BedrockAgentStack');

const pathToLambdaFile = 'pathToYourLambdaFunctionFile';
const s3BucketName = 'nameOfYourS3Bucket';
const s3ObjectKey = 'nameOfYourOpenAPISchemaFileInS3Bucket';

const agentResourceRoleArn = new cdk.aws_iam.Role(stack, 'BedrockAgentRole', {
  roleName: 'AmazonBedrockExecutionRoleForAgents_agent_test',
  assumedBy: new cdk.aws_iam.ServicePrincipal('bedrock.amazonaws.com'),
  managedPolicies: [cdk.aws_iam.ManagedPolicy.fromAwsManagedPolicyName('AdministratorAccess')],
}).roleArn;

const lambdaFunctionRole = new cdk.aws_iam.Role(stack, 'BedrockAgentLambdaFunctionRole', {
  roleName: 'BedrockAgentLambdaFunctionRole',
  assumedBy: new cdk.aws_iam.ServicePrincipal('lambda.amazonaws.com'),
  managedPolicies: [cdk.aws_iam.ManagedPolicy.fromAwsManagedPolicyName('AdministratorAccess')],
});

const actionGroupExecutor = new cdk.aws_lambda.Function(stack, 'BedrockAgentActionGroupExecutor', {
  runtime: cdk.aws_lambda.Runtime.PYTHON_3_10,
  handler: 'youLambdaFileName.nameOfYourHandler',
  code: cdk.aws_lambda.Code.fromAsset(path.join(__dirname, pathToLambdaFile)),
  timeout: cdk.Duration.seconds(600),
  role: lambdaFunctionRole,
}); 

new BedrockAgent(stack, "BedrockAgentStack", {
  agentName: "BedrockAgent",
  instruction: "This is a test instruction. You were built by AWS CDK construct to answer all questions.",
  agentResourceRoleArn: agentResourceRoleArn,
  actionGroups: [{
    actionGroupName: "BedrockAgentActionGroup",
    actionGroupExecutor: actionGroupExecutor.functionArn,
    s3BucketName: s3BucketName,
    s3ObjectKey: s3ObjectKey,
  }]
})
```

# Example - create an agent with 2 action groups, and 2 knowledge bases that you associate with the agent 

Below is an example of how you can create 2 Amazon Bedrock Knowledge Bases (1 backed by OpenSearch Serverless and 1 backed by Pinecone), Amazon Bedrock Agent with 2 action groups, and associate 2 knowledge bases to the agent.

This example assumes you have AWS Lambda function(s) ARN (`actionGroupLambdaArn`), Amazon S3 bucket name (`actionGroupS3BucketName`) with Open API json or yaml files (`actionGroupS3ObjectKey1` and `actionGroupS3ObjectKey1`) that you want your agent to use, as well as Amazon S3 bucket ARN (`dataSourceBucketArn`) where you have files that you want the Knowledge Base to perform ebeddings on. It also assumes that you already have OpenSearch Serverless ARN collection (`openSearchCollectionArn`), Pinecone index created and copied it's host name (`pineconeConnectionString`) and created and stored an API key in AWS Secrets Manager (`pineconeCredentialsSecretArn`). 

Using 2 Knowledge Bases and/or 2 agent action groups is not required, this is just an example. Feel free to experiment with as many Knowledge Bases or Action Groups as you'd like.

**Note: The IAM role creation in this example is for illustration only. Always provion IAM roles with the least required priviliges.**

```typescript
import * as cdk from 'aws-cdk-lib';
import { BedrockAgent, BedrockKnowledgeBase } from 'bedrock-agents-cdk';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'BedrockKnowledgeBaseStack');

const agentName = 'MyTestAgent';
const openSearchKbName = 'MyTestOpenSearchKnowledgeBase';
const pineconeKbName = 'MyTestPineconeKnowledgeBase';
const actionGroupName1 = 'MyTestActionGroup1';
const actionGroupName2 = 'MyTestActionGroup2';
const foundationModel = 'anthropic.claude-instant-v1';
const agentInstruction = 'This is a template instruction for my agent. You were created by AWS CDK.';
const kbInstruction = 'This is a template instruction for my knowledge base. You were created by AWS CDK.';
const openSearchCollectionArn = 'arn:aws:aoss:yourCollectionRegion:yourAWSAccountId:collection/yourCollectionId';
const pineconeCredentialsSecretArn = 'arn:aws:secretsmanager:yourSecretRegion:yourAWSAccountId:secret:yourSecretName';
const pineconeConnectionString = 'https://name-index-aa12345.svc.pineconeRegion.pinecone.io';
const vectorIndexName = 'my-test-index';
const vectorFieldName = 'my-test-vector';
const textField = 'text-field';
const metadataField = 'metadata-field';
const openSearchStorageConfigurationType = 'OPENSEARCH_SERVERLESS';
const pineconeStorageConfigurationType = 'PINECONE';
const dataSourceBucketArn = 'yourDataSourceBucketArn';
const inclusionPrefix = 'yourFolder/';
const actionGroupLambdaArn = 'yourActionGroupLambdaArn';
const actionGroupS3BucketName = 'yourActionGroupApiSchemaBucketName';
const actionGroupS3ObjectKey1 = 'yourActionGroupApiSchemaKey1';
const actionGroupS3ObjectKey2 = 'yourActionGroupApiSchemaKey2';

// Bedrock Agent IAM role
const agentResourceRoleArn = new cdk.aws_iam.Role(stack, 'BedrockAgentRole', {
  roleName: 'AmazonBedrockExecutionRoleForAgents_agent_test',
  assumedBy: new cdk.aws_iam.ServicePrincipal('bedrock.amazonaws.com'),
  managedPolicies: [cdk.aws_iam.ManagedPolicy.fromAwsManagedPolicyName('AdministratorAccess')],
}).roleArn;

// Bedrock Knowledge Base IAM role
const knowledgeBaseRoleArn = new iam.Role(stack, 'BedrockKnowledgeBaseRole', {
  roleName: 'AmazonBedrockExecutionRoleForKnowledgeBase_kb_test',
  assumedBy: new iam.ServicePrincipal('bedrock.amazonaws.com'),
  managedPolicies: [iam.ManagedPolicy.fromAwsManagedPolicyName('AdministratorAccess')],
}).roleArn;

const myOpenSearchKb = new BedrockKnowledgeBase(stack, 'BedrockOpenSearchKnowledgeBase', {
  name: openSearchKbName,
  roleArn: knowledgeBaseRoleArn,
  storageConfiguration: {
    opensearchServerlessConfiguration: {
      collectionArn: openSearchCollectionArn,
      fieldMapping: {
        metadataField: metadataField,
        textField: textField,
        vectorField: vectorFieldName,
      },
      vectorIndexName: vectorIndexName,
    },
    type: openSearchStorageConfigurationType,
  },
  dataSource: {
    dataSourceConfiguration: {
      s3Configuration: {
        bucketArn: dataSourceBucketArn,
      },
    },
  },
});

const myPineconeKb = new BedrockKnowledgeBase(stack, 'BedrockPineconeKnowledgeBase', {
  name: pineconeKbName,
  roleArn: knowledgeBaseRoleArn,
  storageConfiguration: {
    pineconeConfiguration: {
      credentialsSecretArn: pineconeCredentialsSecretArn,
      connectionString: pineconeConnectionString,
      fieldMapping: {
        metadataField: metadataField,
        textField: textField,
      },
    },
    type: pineconeStorageConfigurationType,
  },
  dataSource: {
    dataSourceConfiguration: {
      s3Configuration: {
        bucketArn: dataSourceBucketArn,
        inclusionPrefixes: [inclusionPrefix],
      },
    },
  },
});

const agent = new BedrockAgent(stack, 'BedrockAgent', {
  agentName: agentName,
  instruction: agentInstruction,
  foundationModel: foundationModel,
  agentResourceRoleArn: agentResourceRoleArn,
  actionGroups: [{
    actionGroupName: actionGroupName1,
    actionGroupExecutor: actionGroupLambdaArn,
    s3BucketName: actionGroupS3BucketName,
    s3ObjectKey: actionGroupS3ObjectKey1,
    desription: 'This is a test action group 1 description.',
  },
  {
    actionGroupName: actionGroupName2,
    actionGroupExecutor: actionGroupLambdaArn,
    s3BucketName: actionGroupS3BucketName,
    s3ObjectKey: actionGroupS3ObjectKey2,
  }],
  knowledgeBaseAssociations: [{
    knowledgeBaseName: openSearchKbName,
    instruction: kbInstruction,
  },
  {
    knowledgeBaseName: pineconeKbName,
    instruction: kbInstruction,
  }],
});

agent.node.addDependency(myOpenSearchKb);
agent.node.addDependency(myPineconeKb);
```

# Example - deploy agent, create Amazon OpenSearch Serverless collection and knowledge base backed by it

Below is an example of how you can provision an AWS OpenSearch Serverless collection, Amazon Bedrock Agent and Amazon Bedrock Knowledge using these constructs.

This example assumes you have an AWS Lambda function ARN (`actionGroupLambdaArn`), Amazon S3 bucket name (`actionGroupS3BucketName`) with Open API json or yaml file (`actionGroupS3ObjectKey`) that you want your agent to use, as well as Amazon S3 bucket ARN (`dataSourceBucketArn`) where you have files that you want the Knowledge Base to perform ebeddings on.

You can substitute the other variables (such as `collectionName`, `vectorIndexName`, etc.) as you'd like. It is also important to download [custom_resource](https://github.com/maxtybar/bedrock-agents-cdk/tree/main/custom_resource) and [lambda_layer](https://github.com/maxtybar/bedrock-agents-cdk/tree/main/lambda_layer) folders and include it in your cdk deployment. Substitute `lambdaLayerZipFilePath` and `customResourcePythonFilePath` respectively depending on how you structure your project. This custom resource insures provisioning of OpenSearch indices. 

**Note: The IAM role creation in this example is for illustration only. Always provion IAM roles with the least required priviliges.**

``` typescript
import * as path from 'path';
import {
  App,
  Stack,
  aws_iam as iam,
  aws_opensearchserverless as oss,
  aws_lambda as lambda,
  Duration,
  CustomResource,
  aws_logs as logs,
  custom_resources,
} from 'aws-cdk-lib';
import { BedrockAgent, BedrockKnowledgeBase } from 'bedrock-agents-cdk';

const app = new App();

const stack = new Stack(app, 'BedrockAgentStack');

const agentName = 'MyTestAgent';
const kbName = 'MyTestKnowledgeBase';
const actionGroupName = 'MyTestActionGroup';
const dataSourceName = 'MyDataSource';
const foundationModel = 'anthropic.claude-instant-v1';
const agentInstruction = 'This is a template instruction for my agent. You were created by AWS CDK.';
const kbInstruction = 'This is a template instruction for my knowledge base. You were created by AWS CDK.';
const collectionName = 'my-test-collection';
const vectorIndexName = 'my-test-index';
const vectorFieldName = 'my-test-vector';
const textField = 'text-field';
const metadataField = 'metadata-field';
const storageConfigurationType = 'OPENSEARCH_SERVERLESS';
const dataSourceType = 'S3';
const dataSourceBucketArn = 'yourDataSourceBucketArn';
const actionGroupLambdaArn = 'yourActionGroupLambdaArn';
const actionGroupS3BucketName = 'yourActionGroupApiSchemaBucketName';
const actionGroupS3ObjectKey = 'yourActionGroupApiSchemaKey';
const lambdaLayerZipFilePath = '../lambda_layer/bedrock-agent-layer.zip';
const customResourcePythonFilePath = '../custom_resource';

// Bedrock Agent IAM role
const agentRoleArn = new iam.Role(stack, 'BedrockAgentRole', {
  roleName: 'AmazonBedrockExecutionRoleForAgents_agent_test',
  assumedBy: new iam.ServicePrincipal('bedrock.amazonaws.com'),
  managedPolicies: [iam.ManagedPolicy.fromAwsManagedPolicyName('AdministratorAccess')],
}).roleArn;

// Bedrock Knowledge Base IAM role
const kbRoleArn = new iam.Role(stack, 'BedrockKnowledgeBaseRole', {
  roleName: 'AmazonBedrockExecutionRoleForKnowledgeBase_kb_test',
  assumedBy: new iam.ServicePrincipal('bedrock.amazonaws.com'),
  managedPolicies: [iam.ManagedPolicy.fromAwsManagedPolicyName('AdministratorAccess')],
}).roleArn;

// Lambda IAM role
const customResourceRole = new iam.Role(stack, 'CustomResourceRole', {
  assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com'),
  managedPolicies: [iam.ManagedPolicy.fromAwsManagedPolicyName('service-role/AWSLambdaBasicExecutionRole')],
});

// Opensearch encryption policy
const encryptionPolicy = new oss.CfnSecurityPolicy(stack, 'EncryptionPolicy', {
  name: 'embeddings-encryption-policy',
  type: 'encryption',
  description: `Encryption policy for ${collectionName} collection.`,
  policy: `
  {
    "Rules": [
      {
        "ResourceType": "collection",
        "Resource": ["collection/${collectionName}*"]
      }
    ],
    "AWSOwnedKey": true
  }
  `,
});

// Opensearch network policy
const networkPolicy = new oss.CfnSecurityPolicy(stack, 'NetworkPolicy', {
  name: 'embeddings-network-policy',
  type: 'network',
  description: `Network policy for ${collectionName} collection.`,
  policy: `
    [
      {
        "Rules": [
          {
            "ResourceType": "collection",
            "Resource": ["collection/${collectionName}*"]
          },
          {
            "ResourceType": "dashboard",
            "Resource": ["collection/${collectionName}*"]
          }
        ],
        "AllowFromPublic": true
      }
    ]
  `,
});

// Opensearch data access policy
const dataAccessPolicy = new oss.CfnAccessPolicy(stack, 'DataAccessPolicy', {
  name: 'embeddings-access-policy',
  type: 'data',
  description: `Data access policy for ${collectionName} collection.`,
  policy: `
    [
      {
        "Rules": [
          {
            "ResourceType": "collection",
            "Resource": ["collection/${collectionName}*"],
            "Permission": [
              "aoss:CreateCollectionItems",
              "aoss:DescribeCollectionItems",
              "aoss:DeleteCollectionItems",
              "aoss:UpdateCollectionItems"
            ]
          },
          {
            "ResourceType": "index",
            "Resource": ["index/${collectionName}*/*"],
            "Permission": [
              "aoss:CreateIndex",
              "aoss:DeleteIndex",
              "aoss:UpdateIndex",
              "aoss:DescribeIndex",
              "aoss:ReadDocument",
              "aoss:WriteDocument"
            ]
          }
        ],
        "Principal": [
          "${customResourceRole.roleArn}",
          "${kbRoleArn}"
        ]
      }
    ]
  `,
});

// Opensearch servelrless collection
const opensearchServerlessCollection = new oss.CfnCollection(stack, 'OpenSearchServerlessCollection', {
  name: collectionName,
  description: 'Collection created by CDK to explore vector embeddings and Bedrock Agents.',
  type: 'VECTORSEARCH',
});

// Allow Lambda access to OpenSearch data plane
customResourceRole.addToPolicy(
  new iam.PolicyStatement({
    resources: [opensearchServerlessCollection.attrArn],
    actions: ['aoss:APIAccessAll'],
  }),
);

// Lambda layer
const layer = new lambda.LayerVersion(stack, 'OpenSearchCustomResourceLayer', {
  code: lambda.Code.fromAsset(path.join(__dirname, lambdaLayerZipFilePath)),
  compatibleRuntimes: [lambda.Runtime.PYTHON_3_10],
  description: 'Required dependencies for Lambda',
});

// Lambda function
const onEvent = new lambda.Function(stack, 'OpenSearchCustomResourceFunction', {
  runtime: lambda.Runtime.PYTHON_3_10,
  handler: 'indices_custom_resource.on_event',
  code: lambda.Code.fromAsset(path.join(__dirname, customResourcePythonFilePath)),
  layers: [layer],
  timeout: Duration.seconds(600),
  environment: {
    COLLECTION_ENDPOINT: opensearchServerlessCollection.attrCollectionEndpoint,
    VECTOR_FIELD_NAME: vectorFieldName,
    VECTOR_INDEX_NAME: vectorIndexName,
    TEXT_FIELD: textField,
    METADATA_FIELD: metadataField,
  },
  role: customResourceRole,
});

// Custom resource provider
const provider = new custom_resources.Provider(stack, 'CustomResourceProvider', {
  onEventHandler: onEvent,
  logRetention: logs.RetentionDays.ONE_DAY,
});

// Custom resource
new CustomResource(stack, 'CustomResource', {
  serviceToken: provider.serviceToken,
});

// Create Bedrock Knowledge Base backed by OpenSearch Servereless
const myOpenSearchKb = new BedrockKnowledgeBase(stack, 'BedrockOpenSearchKnowledgeBase', {
  name: kbName,
  roleArn: kbRoleArn,
  storageConfiguration: {
    opensearchServerlessConfiguration: {
      collectionArn: opensearchServerlessCollection.attrArn,
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
      },
      type: dataSourceType,
    },
  },
});

// Amazon Bedrock Agent and Knowledge Base backed by Opensearch Serverless
const bedrockAgent = new BedrockAgent(stack, 'BedrockAgent', {
  agentName: agentName,
  instruction: agentInstruction,
  foundationModel: foundationModel,
  agentResourceRoleArn: agentRoleArn,
  actionGroups: [{
    actionGroupName: actionGroupName,
    actionGroupExecutor: actionGroupLambdaArn,
    s3BucketName: actionGroupS3BucketName,
    s3ObjectKey: actionGroupS3ObjectKey,
  }],
  knowledgeBaseAssociations: [{
    knowledgeBaseName: kbName,
    instruction: kbInstruction,
  }],
});

opensearchServerlessCollection.node.addDependency(encryptionPolicy);
opensearchServerlessCollection.node.addDependency(networkPolicy);
opensearchServerlessCollection.node.addDependency(dataAccessPolicy);
onEvent.node.addDependency(opensearchServerlessCollection);
bedrockAgent.node.addDependency(onEvent);
bedrockAgent.node.addDependency(myOpenSearchKb);
```