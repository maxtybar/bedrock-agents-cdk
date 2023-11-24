# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### BedrockAgent <a name="BedrockAgent" id="bedrock-agents-cdk.BedrockAgent"></a>

#### Initializers <a name="Initializers" id="bedrock-agents-cdk.BedrockAgent.Initializer"></a>

```typescript
import { BedrockAgent } from 'bedrock-agents-cdk'

new BedrockAgent(scope: Construct, name: string, props: BedrockAgentProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.Initializer.parameter.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.Initializer.parameter.props">props</a></code> | <code><a href="#bedrock-agents-cdk.BedrockAgentProps">BedrockAgentProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="bedrock-agents-cdk.BedrockAgent.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `name`<sup>Required</sup> <a name="name" id="bedrock-agents-cdk.BedrockAgent.Initializer.parameter.name"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="bedrock-agents-cdk.BedrockAgent.Initializer.parameter.props"></a>

- *Type:* <a href="#bedrock-agents-cdk.BedrockAgentProps">BedrockAgentProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="bedrock-agents-cdk.BedrockAgent.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="bedrock-agents-cdk.BedrockAgent.isConstruct"></a>

```typescript
import { BedrockAgent } from 'bedrock-agents-cdk'

BedrockAgent.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="bedrock-agents-cdk.BedrockAgent.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.actionGroup">actionGroup</a></code> | <code><a href="#bedrock-agents-cdk.ActionGroup">ActionGroup</a></code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.agentName">agentName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.agentResourceRoleArn">agentResourceRoleArn</a></code> | <code>string</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.bedrockAgentCustomResourceRole">bedrockAgentCustomResourceRole</a></code> | <code>aws-cdk-lib.aws_iam.Role</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.foundationModel">foundationModel</a></code> | <code>string</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.idleSessionTTLInSeconds">idleSessionTTLInSeconds</a></code> | <code>number</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.instruction">instruction</a></code> | <code>string</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.knowledgeBase">knowledgeBase</a></code> | <code><a href="#bedrock-agents-cdk.KnowledgeBase">KnowledgeBase</a></code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.region">region</a></code> | <code>string</code> | *No description.* |
| <code><a href="#bedrock-agents-cdk.BedrockAgent.property.description">description</a></code> | <code>string</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="bedrock-agents-cdk.BedrockAgent.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `actionGroup`<sup>Required</sup> <a name="actionGroup" id="bedrock-agents-cdk.BedrockAgent.property.actionGroup"></a>

```typescript
public readonly actionGroup: ActionGroup;
```

- *Type:* <a href="#bedrock-agents-cdk.ActionGroup">ActionGroup</a>

---

##### `agentName`<sup>Required</sup> <a name="agentName" id="bedrock-agents-cdk.BedrockAgent.property.agentName"></a>

```typescript
public readonly agentName: string;
```

- *Type:* string

---

##### `agentResourceRoleArn`<sup>Required</sup> <a name="agentResourceRoleArn" id="bedrock-agents-cdk.BedrockAgent.property.agentResourceRoleArn"></a>

```typescript
public readonly agentResourceRoleArn: string;
```

- *Type:* string

---

##### `bedrockAgentCustomResourceRole`<sup>Required</sup> <a name="bedrockAgentCustomResourceRole" id="bedrock-agents-cdk.BedrockAgent.property.bedrockAgentCustomResourceRole"></a>

```typescript
public readonly bedrockAgentCustomResourceRole: Role;
```

- *Type:* aws-cdk-lib.aws_iam.Role

---

##### `foundationModel`<sup>Required</sup> <a name="foundationModel" id="bedrock-agents-cdk.BedrockAgent.property.foundationModel"></a>

```typescript
public readonly foundationModel: string;
```

- *Type:* string

---

##### `idleSessionTTLInSeconds`<sup>Required</sup> <a name="idleSessionTTLInSeconds" id="bedrock-agents-cdk.BedrockAgent.property.idleSessionTTLInSeconds"></a>

```typescript
public readonly idleSessionTTLInSeconds: number;
```

- *Type:* number

---

##### `instruction`<sup>Required</sup> <a name="instruction" id="bedrock-agents-cdk.BedrockAgent.property.instruction"></a>

```typescript
public readonly instruction: string;
```

- *Type:* string

---

##### `knowledgeBase`<sup>Required</sup> <a name="knowledgeBase" id="bedrock-agents-cdk.BedrockAgent.property.knowledgeBase"></a>

```typescript
public readonly knowledgeBase: KnowledgeBase;
```

- *Type:* <a href="#bedrock-agents-cdk.KnowledgeBase">KnowledgeBase</a>

---

##### `region`<sup>Required</sup> <a name="region" id="bedrock-agents-cdk.BedrockAgent.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

---

##### `description`<sup>Optional</sup> <a name="description" id="bedrock-agents-cdk.BedrockAgent.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string

---


## Structs <a name="Structs" id="Structs"></a>

### ActionGroup <a name="ActionGroup" id="bedrock-agents-cdk.ActionGroup"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.ActionGroup.Initializer"></a>

```typescript
import { ActionGroup } from 'bedrock-agents-cdk'

const actionGroup: ActionGroup = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.ActionGroup.property.actionGroupExecutor">actionGroupExecutor</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.ActionGroup.property.actionGroupName">actionGroupName</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.ActionGroup.property.s3BucketName">s3BucketName</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.ActionGroup.property.s3ObjectKey">s3ObjectKey</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.ActionGroup.property.desription">desription</a></code> | <code>string</code> | Optional. |

---

##### `actionGroupExecutor`<sup>Required</sup> <a name="actionGroupExecutor" id="bedrock-agents-cdk.ActionGroup.property.actionGroupExecutor"></a>

```typescript
public readonly actionGroupExecutor: string;
```

- *Type:* string

Required.

Lambda function ARN that will be used in this action group.
The provided Lambda function will be assigned a resource-based policy
to allow access from the newly created agent.

---

##### `actionGroupName`<sup>Required</sup> <a name="actionGroupName" id="bedrock-agents-cdk.ActionGroup.property.actionGroupName"></a>

```typescript
public readonly actionGroupName: string;
```

- *Type:* string

Required.

Action group name.

---

##### `s3BucketName`<sup>Required</sup> <a name="s3BucketName" id="bedrock-agents-cdk.ActionGroup.property.s3BucketName"></a>

```typescript
public readonly s3BucketName: string;
```

- *Type:* string

Required.

S3 bucket name where Open API schema is stored.
Bucket must be in the same region where agent is created.

---

##### `s3ObjectKey`<sup>Required</sup> <a name="s3ObjectKey" id="bedrock-agents-cdk.ActionGroup.property.s3ObjectKey"></a>

```typescript
public readonly s3ObjectKey: string;
```

- *Type:* string

Required.

S3 bucket key for the actual schema.
Must be either JSON or yaml file.

---

##### `desription`<sup>Optional</sup> <a name="desription" id="bedrock-agents-cdk.ActionGroup.property.desription"></a>

```typescript
public readonly desription: string;
```

- *Type:* string

Optional.

Description for the action group.

---

### BaseFieldMapping <a name="BaseFieldMapping" id="bedrock-agents-cdk.BaseFieldMapping"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.BaseFieldMapping.Initializer"></a>

```typescript
import { BaseFieldMapping } from 'bedrock-agents-cdk'

const baseFieldMapping: BaseFieldMapping = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.BaseFieldMapping.property.metadataField">metadataField</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.BaseFieldMapping.property.textField">textField</a></code> | <code>string</code> | Required. |

---

##### `metadataField`<sup>Required</sup> <a name="metadataField" id="bedrock-agents-cdk.BaseFieldMapping.property.metadataField"></a>

```typescript
public readonly metadataField: string;
```

- *Type:* string

Required.

Metadata field that you configured in your Vector DB.
Bedrock will store metadata that is required to carry out souce attribution
and enable data ingestion and querying for OpenSearch and Pinecone and
will store raw text from your data in chunks in this field for Redis Enterprise Cloud.

---

##### `textField`<sup>Required</sup> <a name="textField" id="bedrock-agents-cdk.BaseFieldMapping.property.textField"></a>

```typescript
public readonly textField: string;
```

- *Type:* string

Required.

Text field that you configured in your Vector DB.
Bedrock will store raw text from your data in chunks in this field.

---

### BedrockAgentProps <a name="BedrockAgentProps" id="bedrock-agents-cdk.BedrockAgentProps"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.BedrockAgentProps.Initializer"></a>

```typescript
import { BedrockAgentProps } from 'bedrock-agents-cdk'

const bedrockAgentProps: BedrockAgentProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.agentName">agentName</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.instruction">instruction</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.actionGroup">actionGroup</a></code> | <code><a href="#bedrock-agents-cdk.ActionGroup">ActionGroup</a></code> | Optional. |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.agentResourceRoleArn">agentResourceRoleArn</a></code> | <code>string</code> | Optional. |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.description">description</a></code> | <code>string</code> | Optional. |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.foundationModel">foundationModel</a></code> | <code>string</code> | Optional. |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.idleSessionTTLInSeconds">idleSessionTTLInSeconds</a></code> | <code>number</code> | Optional. |
| <code><a href="#bedrock-agents-cdk.BedrockAgentProps.property.knowledgeBase">knowledgeBase</a></code> | <code><a href="#bedrock-agents-cdk.KnowledgeBase">KnowledgeBase</a></code> | Optional. |

---

##### `agentName`<sup>Required</sup> <a name="agentName" id="bedrock-agents-cdk.BedrockAgentProps.property.agentName"></a>

```typescript
public readonly agentName: string;
```

- *Type:* string

Required.

Name of the agent.

---

##### `instruction`<sup>Required</sup> <a name="instruction" id="bedrock-agents-cdk.BedrockAgentProps.property.instruction"></a>

```typescript
public readonly instruction: string;
```

- *Type:* string

Required.

Instruction for the agent.
Characters length:

---

##### `actionGroup`<sup>Optional</sup> <a name="actionGroup" id="bedrock-agents-cdk.BedrockAgentProps.property.actionGroup"></a>

```typescript
public readonly actionGroup: ActionGroup;
```

- *Type:* <a href="#bedrock-agents-cdk.ActionGroup">ActionGroup</a>

Optional.

Action group for the agent. If specified must contain ``s3BucketName`` and ``s3ObjectKey`` with JSON
or yaml Open API schema, as well as Lambda ARN for the action group executor,
and action group name.

---

##### `agentResourceRoleArn`<sup>Optional</sup> <a name="agentResourceRoleArn" id="bedrock-agents-cdk.BedrockAgentProps.property.agentResourceRoleArn"></a>

```typescript
public readonly agentResourceRoleArn: string;
```

- *Type:* string

Optional.

Resource role ARN for an agent.
Role name must start with ``AmazonBedrockExecutionRoleForAgents_`` prefix and assumed by ``bedrock.amazonaws.com``.
If role is not specified the default one will be created with no attached policies to it.
If actionGroup is specified and the role is not, then the default created role will have an attached S3 read access
to the bucket provided in the actionGroup.

---

##### `description`<sup>Optional</sup> <a name="description" id="bedrock-agents-cdk.BedrockAgentProps.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string

Optional.

Description for the agent.

---

##### `foundationModel`<sup>Optional</sup> <a name="foundationModel" id="bedrock-agents-cdk.BedrockAgentProps.property.foundationModel"></a>

```typescript
public readonly foundationModel: string;
```

- *Type:* string
- *Default:* "anthropic.claude-v2"

Optional.

Foundation model.

---

*Example*

```typescript
- "anthropic.claude-instant-v1" or "anthropic.claude-v2"
```


##### `idleSessionTTLInSeconds`<sup>Optional</sup> <a name="idleSessionTTLInSeconds" id="bedrock-agents-cdk.BedrockAgentProps.property.idleSessionTTLInSeconds"></a>

```typescript
public readonly idleSessionTTLInSeconds: number;
```

- *Type:* number
- *Default:* 3600

Optional.

Max Session Time in seconds.

---

##### `knowledgeBase`<sup>Optional</sup> <a name="knowledgeBase" id="bedrock-agents-cdk.BedrockAgentProps.property.knowledgeBase"></a>

```typescript
public readonly knowledgeBase: KnowledgeBase;
```

- *Type:* <a href="#bedrock-agents-cdk.KnowledgeBase">KnowledgeBase</a>

Optional.

Knowledge base. If provided must be either of type ``openSearchServerlessConfiguration``
or ``pineconeConfiguration`` or ``redisEnterpriseCloudConfiguration``.

> [https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-create.html](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-create.html)

---

### DataSource <a name="DataSource" id="bedrock-agents-cdk.DataSource"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.DataSource.Initializer"></a>

```typescript
import { DataSource } from 'bedrock-agents-cdk'

const dataSource: DataSource = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.DataSource.property.dataSourceConfiguration">dataSourceConfiguration</a></code> | <code><a href="#bedrock-agents-cdk.DataSourceConfiguration">DataSourceConfiguration</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.DataSource.property.name">name</a></code> | <code>string</code> | Optional. |

---

##### `dataSourceConfiguration`<sup>Required</sup> <a name="dataSourceConfiguration" id="bedrock-agents-cdk.DataSource.property.dataSourceConfiguration"></a>

```typescript
public readonly dataSourceConfiguration: DataSourceConfiguration;
```

- *Type:* <a href="#bedrock-agents-cdk.DataSourceConfiguration">DataSourceConfiguration</a>

Required.

Data Source Configuration.

---

*Example*

```typescript
dataSourceConfiguration = {
s3Configuration: {
  bucketArn: "yourS3BucketArn",
},
  "type": "S3"
}
```


##### `name`<sup>Optional</sup> <a name="name" id="bedrock-agents-cdk.DataSource.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string
- *Default:* `MyDataSource-${agentName}`

Optional.

Name for your data source.

---

### DataSourceConfiguration <a name="DataSourceConfiguration" id="bedrock-agents-cdk.DataSourceConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.DataSourceConfiguration.Initializer"></a>

```typescript
import { DataSourceConfiguration } from 'bedrock-agents-cdk'

const dataSourceConfiguration: DataSourceConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.DataSourceConfiguration.property.s3Configuration">s3Configuration</a></code> | <code><a href="#bedrock-agents-cdk.S3Configuration">S3Configuration</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.DataSourceConfiguration.property.type">type</a></code> | <code>string</code> | Optional. |

---

##### `s3Configuration`<sup>Required</sup> <a name="s3Configuration" id="bedrock-agents-cdk.DataSourceConfiguration.property.s3Configuration"></a>

```typescript
public readonly s3Configuration: S3Configuration;
```

- *Type:* <a href="#bedrock-agents-cdk.S3Configuration">S3Configuration</a>

Required.

S3 Configuration.

---

*Example*

```typescript
  s3Configuration: {
    bucketArn: "yourS3BucketArn"
  }
```


##### `type`<sup>Optional</sup> <a name="type" id="bedrock-agents-cdk.DataSourceConfiguration.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string
- *Default:* "S3"

Optional.

Type of configuration.

---

### KnowledgeBase <a name="KnowledgeBase" id="bedrock-agents-cdk.KnowledgeBase"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.KnowledgeBase.Initializer"></a>

```typescript
import { KnowledgeBase } from 'bedrock-agents-cdk'

const knowledgeBase: KnowledgeBase = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.dataSource">dataSource</a></code> | <code><a href="#bedrock-agents-cdk.DataSource">DataSource</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.instruction">instruction</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.name">name</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.roleArn">roleArn</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.storageConfiguration">storageConfiguration</a></code> | <code><a href="#bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration">OpenSearchServerlessStorageConfiguration</a> \| <a href="#bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration">RedisEnterpriseCloudStorageConfiguration</a> \| <a href="#bedrock-agents-cdk.PineconeStorageConfiguration">PineconeStorageConfiguration</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.description">description</a></code> | <code>string</code> | Optional. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.knowledgeBaseConfiguration">knowledgeBaseConfiguration</a></code> | <code><a href="#bedrock-agents-cdk.KnowledgeBaseConfiguration">KnowledgeBaseConfiguration</a></code> | Optional. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBase.property.removalPolicy">removalPolicy</a></code> | <code>aws-cdk-lib.RemovalPolicy</code> | Optional. |

---

##### `dataSource`<sup>Required</sup> <a name="dataSource" id="bedrock-agents-cdk.KnowledgeBase.property.dataSource"></a>

```typescript
public readonly dataSource: DataSource;
```

- *Type:* <a href="#bedrock-agents-cdk.DataSource">DataSource</a>

Required.

Data source configuration.

---

##### `instruction`<sup>Required</sup> <a name="instruction" id="bedrock-agents-cdk.KnowledgeBase.property.instruction"></a>

```typescript
public readonly instruction: string;
```

- *Type:* string

Required.

Instruction based on the design and type of information of the knowledge base.
This will impact how the knowledge base interacts with the agent.

---

##### `name`<sup>Required</sup> <a name="name" id="bedrock-agents-cdk.KnowledgeBase.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Required.

Name of the KnowledgeBase.

---

##### `roleArn`<sup>Required</sup> <a name="roleArn" id="bedrock-agents-cdk.KnowledgeBase.property.roleArn"></a>

```typescript
public readonly roleArn: string;
```

- *Type:* string

Required.

Resource role ARN for a knowledge base.
Role name must start with ``AmazonBedrockExecutionRoleForKnowledgeBase_`` prefix and assumed by ``bedrock.amazonaws.com``.
Role must have access to the S3 bucket used as a data source as a knowledge base.
If you use OpenSearch serverless, the role must have ``aoss:APIAccessAll`` policy attached to it
allowing it to make API calls against your collection's data plane. Your collection
must also allow data access from KnowledgeBase role.
See more here @see https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html

---

##### `storageConfiguration`<sup>Required</sup> <a name="storageConfiguration" id="bedrock-agents-cdk.KnowledgeBase.property.storageConfiguration"></a>

```typescript
public readonly storageConfiguration: OpenSearchServerlessStorageConfiguration | RedisEnterpriseCloudStorageConfiguration | PineconeStorageConfiguration;
```

- *Type:* <a href="#bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration">OpenSearchServerlessStorageConfiguration</a> | <a href="#bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration">RedisEnterpriseCloudStorageConfiguration</a> | <a href="#bedrock-agents-cdk.PineconeStorageConfiguration">PineconeStorageConfiguration</a>

Required.

KnowledgeBase storage configuration.
Has to be either ``opensearchServerlessConfiguration`` or
``pineconeConfiguration`` or ``redisEnterpriseCloudConfiguration``
and respective type field mapping.

---

##### `description`<sup>Optional</sup> <a name="description" id="bedrock-agents-cdk.KnowledgeBase.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string

Optional.

Description for the knowledge base.

---

##### `knowledgeBaseConfiguration`<sup>Optional</sup> <a name="knowledgeBaseConfiguration" id="bedrock-agents-cdk.KnowledgeBase.property.knowledgeBaseConfiguration"></a>

```typescript
public readonly knowledgeBaseConfiguration: KnowledgeBaseConfiguration;
```

- *Type:* <a href="#bedrock-agents-cdk.KnowledgeBaseConfiguration">KnowledgeBaseConfiguration</a>
- *Default:* type: "VECTOR", vectorKnowledgeBaseConfiguration: {     embeddingModelArn: "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1" }

Optional.

KnowledgeBase configuration.

---

##### `removalPolicy`<sup>Optional</sup> <a name="removalPolicy" id="bedrock-agents-cdk.KnowledgeBase.property.removalPolicy"></a>

```typescript
public readonly removalPolicy: RemovalPolicy;
```

- *Type:* aws-cdk-lib.RemovalPolicy
- *Default:* cdk.RemovalPolicy.DESTROY

Optional.

Removal Policy. If you want to keep your Knowledge Base intact
in case you destroy this CDK make sure you set removalPolicy to
``cdk.RemovalPolicy.RETAIN``. By default your Knowledge Base will be
deleted along with the agent.

---

### KnowledgeBaseConfiguration <a name="KnowledgeBaseConfiguration" id="bedrock-agents-cdk.KnowledgeBaseConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.KnowledgeBaseConfiguration.Initializer"></a>

```typescript
import { KnowledgeBaseConfiguration } from 'bedrock-agents-cdk'

const knowledgeBaseConfiguration: KnowledgeBaseConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.KnowledgeBaseConfiguration.property.type">type</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.KnowledgeBaseConfiguration.property.vectorKnowledgeBaseConfiguration">vectorKnowledgeBaseConfiguration</a></code> | <code><a href="#bedrock-agents-cdk.VectorKnowledgeBaseConfiguration">VectorKnowledgeBaseConfiguration</a></code> | Required. |

---

##### `type`<sup>Required</sup> <a name="type" id="bedrock-agents-cdk.KnowledgeBaseConfiguration.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string
- *Default:* "VECTOR"

Required.

Type of configuration.

---

##### `vectorKnowledgeBaseConfiguration`<sup>Required</sup> <a name="vectorKnowledgeBaseConfiguration" id="bedrock-agents-cdk.KnowledgeBaseConfiguration.property.vectorKnowledgeBaseConfiguration"></a>

```typescript
public readonly vectorKnowledgeBaseConfiguration: VectorKnowledgeBaseConfiguration;
```

- *Type:* <a href="#bedrock-agents-cdk.VectorKnowledgeBaseConfiguration">VectorKnowledgeBaseConfiguration</a>
- *Default:* vectorKnowledgeBaseConfiguration: { embeddingModelArn: "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1" }

Required.

Embeddings model to convert your data into an embedding.

---

### OpenSearchFieldMapping <a name="OpenSearchFieldMapping" id="bedrock-agents-cdk.OpenSearchFieldMapping"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.OpenSearchFieldMapping.Initializer"></a>

```typescript
import { OpenSearchFieldMapping } from 'bedrock-agents-cdk'

const openSearchFieldMapping: OpenSearchFieldMapping = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.OpenSearchFieldMapping.property.metadataField">metadataField</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.OpenSearchFieldMapping.property.textField">textField</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.OpenSearchFieldMapping.property.vectorField">vectorField</a></code> | <code>string</code> | Required. |

---

##### `metadataField`<sup>Required</sup> <a name="metadataField" id="bedrock-agents-cdk.OpenSearchFieldMapping.property.metadataField"></a>

```typescript
public readonly metadataField: string;
```

- *Type:* string

Required.

Metadata field that you configured in your Vector DB.
Bedrock will store metadata that is required to carry out souce attribution
and enable data ingestion and querying for OpenSearch and Pinecone and
will store raw text from your data in chunks in this field for Redis Enterprise Cloud.

---

##### `textField`<sup>Required</sup> <a name="textField" id="bedrock-agents-cdk.OpenSearchFieldMapping.property.textField"></a>

```typescript
public readonly textField: string;
```

- *Type:* string

Required.

Text field that you configured in your Vector DB.
Bedrock will store raw text from your data in chunks in this field.

---

##### `vectorField`<sup>Required</sup> <a name="vectorField" id="bedrock-agents-cdk.OpenSearchFieldMapping.property.vectorField"></a>

```typescript
public readonly vectorField: string;
```

- *Type:* string

Required.

Vector field that you configured in OpenSearch.
Bedrock will store the vector embeddings in this field.

---

### OpenSearchServerlessConfiguration <a name="OpenSearchServerlessConfiguration" id="bedrock-agents-cdk.OpenSearchServerlessConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.OpenSearchServerlessConfiguration.Initializer"></a>

```typescript
import { OpenSearchServerlessConfiguration } from 'bedrock-agents-cdk'

const openSearchServerlessConfiguration: OpenSearchServerlessConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.OpenSearchServerlessConfiguration.property.collectionArn">collectionArn</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.OpenSearchServerlessConfiguration.property.fieldMapping">fieldMapping</a></code> | <code><a href="#bedrock-agents-cdk.OpenSearchFieldMapping">OpenSearchFieldMapping</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.OpenSearchServerlessConfiguration.property.vectorIndexName">vectorIndexName</a></code> | <code>string</code> | Required. |

---

##### `collectionArn`<sup>Required</sup> <a name="collectionArn" id="bedrock-agents-cdk.OpenSearchServerlessConfiguration.property.collectionArn"></a>

```typescript
public readonly collectionArn: string;
```

- *Type:* string

Required.

ARN of your OpenSearch Serverless Collection.

---

##### `fieldMapping`<sup>Required</sup> <a name="fieldMapping" id="bedrock-agents-cdk.OpenSearchServerlessConfiguration.property.fieldMapping"></a>

```typescript
public readonly fieldMapping: OpenSearchFieldMapping;
```

- *Type:* <a href="#bedrock-agents-cdk.OpenSearchFieldMapping">OpenSearchFieldMapping</a>

Required.

Field mapping consisting of ``vectorField``,
``textField`` and ``metadataField``.

---

##### `vectorIndexName`<sup>Required</sup> <a name="vectorIndexName" id="bedrock-agents-cdk.OpenSearchServerlessConfiguration.property.vectorIndexName"></a>

```typescript
public readonly vectorIndexName: string;
```

- *Type:* string

Required.

Vector index name of your OpenSearch Serverless Collection.

---

### OpenSearchServerlessStorageConfiguration <a name="OpenSearchServerlessStorageConfiguration" id="bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration.Initializer"></a>

```typescript
import { OpenSearchServerlessStorageConfiguration } from 'bedrock-agents-cdk'

const openSearchServerlessStorageConfiguration: OpenSearchServerlessStorageConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration.property.opensearchServerlessConfiguration">opensearchServerlessConfiguration</a></code> | <code><a href="#bedrock-agents-cdk.OpenSearchServerlessConfiguration">OpenSearchServerlessConfiguration</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration.property.type">type</a></code> | <code>string</code> | Required. |

---

##### `opensearchServerlessConfiguration`<sup>Required</sup> <a name="opensearchServerlessConfiguration" id="bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration.property.opensearchServerlessConfiguration"></a>

```typescript
public readonly opensearchServerlessConfiguration: OpenSearchServerlessConfiguration;
```

- *Type:* <a href="#bedrock-agents-cdk.OpenSearchServerlessConfiguration">OpenSearchServerlessConfiguration</a>

Required.

OpenSearch Serverless Configuration.

---

*Example*

```typescript
opensearchServerlessConfiguration: {
    collectionArn: "arn:aws:opensearchserverless:us-east-1:123456789012:collection/my-collection",
    fieldMapping: {
        textField: "text",
        metadataField: "metadata",
        vectorField: "vector"
    },
    vectorIndexName: "my-index",
},
```


##### `type`<sup>Required</sup> <a name="type" id="bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string

Required.

Has to be ``"OPENSEARCH_SERVERLESS"`` for Opensearch Serverless Configuration.

---

### PineconeConfiguration <a name="PineconeConfiguration" id="bedrock-agents-cdk.PineconeConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.PineconeConfiguration.Initializer"></a>

```typescript
import { PineconeConfiguration } from 'bedrock-agents-cdk'

const pineconeConfiguration: PineconeConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.PineconeConfiguration.property.connectionString">connectionString</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.PineconeConfiguration.property.credentialsSecretArn">credentialsSecretArn</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.PineconeConfiguration.property.fieldMapping">fieldMapping</a></code> | <code><a href="#bedrock-agents-cdk.PineconeFieldMapping">PineconeFieldMapping</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.PineconeConfiguration.property.namespace">namespace</a></code> | <code>string</code> | Optional. |

---

##### `connectionString`<sup>Required</sup> <a name="connectionString" id="bedrock-agents-cdk.PineconeConfiguration.property.connectionString"></a>

```typescript
public readonly connectionString: string;
```

- *Type:* string

Required.

Connection string for your Pinecone index management page.

---

##### `credentialsSecretArn`<sup>Required</sup> <a name="credentialsSecretArn" id="bedrock-agents-cdk.PineconeConfiguration.property.credentialsSecretArn"></a>

```typescript
public readonly credentialsSecretArn: string;
```

- *Type:* string

Required.

ARN of the secret containing the API Key to use when connecting to the Pinecone database.
Learn more in the link below.

> [https://www.pinecone.io/blog/amazon-bedrock-integration/](https://www.pinecone.io/blog/amazon-bedrock-integration/)

---

##### `fieldMapping`<sup>Required</sup> <a name="fieldMapping" id="bedrock-agents-cdk.PineconeConfiguration.property.fieldMapping"></a>

```typescript
public readonly fieldMapping: PineconeFieldMapping;
```

- *Type:* <a href="#bedrock-agents-cdk.PineconeFieldMapping">PineconeFieldMapping</a>

Required.

Field mapping consisting of``textField`` and ``metadataField``.

---

##### `namespace`<sup>Optional</sup> <a name="namespace" id="bedrock-agents-cdk.PineconeConfiguration.property.namespace"></a>

```typescript
public readonly namespace: string;
```

- *Type:* string

Optional.

Name space that will be used for writing new data to your Pinecone database.

---

### PineconeFieldMapping <a name="PineconeFieldMapping" id="bedrock-agents-cdk.PineconeFieldMapping"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.PineconeFieldMapping.Initializer"></a>

```typescript
import { PineconeFieldMapping } from 'bedrock-agents-cdk'

const pineconeFieldMapping: PineconeFieldMapping = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.PineconeFieldMapping.property.metadataField">metadataField</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.PineconeFieldMapping.property.textField">textField</a></code> | <code>string</code> | Required. |

---

##### `metadataField`<sup>Required</sup> <a name="metadataField" id="bedrock-agents-cdk.PineconeFieldMapping.property.metadataField"></a>

```typescript
public readonly metadataField: string;
```

- *Type:* string

Required.

Metadata field that you configured in your Vector DB.
Bedrock will store metadata that is required to carry out souce attribution
and enable data ingestion and querying for OpenSearch and Pinecone and
will store raw text from your data in chunks in this field for Redis Enterprise Cloud.

---

##### `textField`<sup>Required</sup> <a name="textField" id="bedrock-agents-cdk.PineconeFieldMapping.property.textField"></a>

```typescript
public readonly textField: string;
```

- *Type:* string

Required.

Text field that you configured in your Vector DB.
Bedrock will store raw text from your data in chunks in this field.

---

### PineconeStorageConfiguration <a name="PineconeStorageConfiguration" id="bedrock-agents-cdk.PineconeStorageConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.PineconeStorageConfiguration.Initializer"></a>

```typescript
import { PineconeStorageConfiguration } from 'bedrock-agents-cdk'

const pineconeStorageConfiguration: PineconeStorageConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.PineconeStorageConfiguration.property.pineconeConfiguration">pineconeConfiguration</a></code> | <code><a href="#bedrock-agents-cdk.PineconeConfiguration">PineconeConfiguration</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.PineconeStorageConfiguration.property.type">type</a></code> | <code>string</code> | Required. |

---

##### `pineconeConfiguration`<sup>Required</sup> <a name="pineconeConfiguration" id="bedrock-agents-cdk.PineconeStorageConfiguration.property.pineconeConfiguration"></a>

```typescript
public readonly pineconeConfiguration: PineconeConfiguration;
```

- *Type:* <a href="#bedrock-agents-cdk.PineconeConfiguration">PineconeConfiguration</a>

Required.

Pinecone Configuration.

---

*Example*

```typescript
pineconeConfiguration: {
    credentialsSecretArn: 'arn:aws:secretsmanager:your-region:123456789098:secret:apiKey';
    connectionString: 'https://your-connection-string.pinecone.io';
    fieldMapping: {
        metadataField: 'metadata-field',
        textField: 'text-field'
    },
},
```


##### `type`<sup>Required</sup> <a name="type" id="bedrock-agents-cdk.PineconeStorageConfiguration.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string

Required.

Has to be ``"PINECONE"`` for Pinecone Configuration.

---

### RedisEnterpriseCloudConfiguration <a name="RedisEnterpriseCloudConfiguration" id="bedrock-agents-cdk.RedisEnterpriseCloudConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.Initializer"></a>

```typescript
import { RedisEnterpriseCloudConfiguration } from 'bedrock-agents-cdk'

const redisEnterpriseCloudConfiguration: RedisEnterpriseCloudConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.credentialsSecretArn">credentialsSecretArn</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.endpointUrl">endpointUrl</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.fieldMapping">fieldMapping</a></code> | <code><a href="#bedrock-agents-cdk.RedisFieldMapping">RedisFieldMapping</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.vectorIndexName">vectorIndexName</a></code> | <code>string</code> | Required. |

---

##### `credentialsSecretArn`<sup>Required</sup> <a name="credentialsSecretArn" id="bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.credentialsSecretArn"></a>

```typescript
public readonly credentialsSecretArn: string;
```

- *Type:* string

Required.

ARN of the secret defining the username, password, serverCertificate,
clientCertificate and clientPrivateKey to use when connecting to the Redis Enterprise Cloud database.
Learn more in the link below.

> [https://docs.redis.com/latest/rc/cloud-integrations/aws-marketplace/aws-bedrock/set-up-redis/](https://docs.redis.com/latest/rc/cloud-integrations/aws-marketplace/aws-bedrock/set-up-redis/)

---

##### `endpointUrl`<sup>Required</sup> <a name="endpointUrl" id="bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.endpointUrl"></a>

```typescript
public readonly endpointUrl: string;
```

- *Type:* string

Required.

The endpoint URL for your Redis Enterprise Cloud database.

---

##### `fieldMapping`<sup>Required</sup> <a name="fieldMapping" id="bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.fieldMapping"></a>

```typescript
public readonly fieldMapping: RedisFieldMapping;
```

- *Type:* <a href="#bedrock-agents-cdk.RedisFieldMapping">RedisFieldMapping</a>

Required.

Field mapping consisting of ``vectorField``,
``textField`` and ``metadataField``.

---

##### `vectorIndexName`<sup>Required</sup> <a name="vectorIndexName" id="bedrock-agents-cdk.RedisEnterpriseCloudConfiguration.property.vectorIndexName"></a>

```typescript
public readonly vectorIndexName: string;
```

- *Type:* string

Required.

Vector index name of your Redis Enterprise Cloud.

---

### RedisEnterpriseCloudStorageConfiguration <a name="RedisEnterpriseCloudStorageConfiguration" id="bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration.Initializer"></a>

```typescript
import { RedisEnterpriseCloudStorageConfiguration } from 'bedrock-agents-cdk'

const redisEnterpriseCloudStorageConfiguration: RedisEnterpriseCloudStorageConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration.property.redisEnterpriseCloudConfiguration">redisEnterpriseCloudConfiguration</a></code> | <code><a href="#bedrock-agents-cdk.RedisEnterpriseCloudConfiguration">RedisEnterpriseCloudConfiguration</a></code> | Required. |
| <code><a href="#bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration.property.type">type</a></code> | <code>string</code> | Required. |

---

##### `redisEnterpriseCloudConfiguration`<sup>Required</sup> <a name="redisEnterpriseCloudConfiguration" id="bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration.property.redisEnterpriseCloudConfiguration"></a>

```typescript
public readonly redisEnterpriseCloudConfiguration: RedisEnterpriseCloudConfiguration;
```

- *Type:* <a href="#bedrock-agents-cdk.RedisEnterpriseCloudConfiguration">RedisEnterpriseCloudConfiguration</a>

Required.

Redis Enterprise Cloud Configuration.

---

*Example*

```typescript
redisEnterpriseCloudConfiguration: {
    credentialsSecretArn: 'arn:aws:secretsmanager:your-region:123456789098:secret:apiKey';
    endpointUrl: 'your-endpoint-url';
    fieldMapping: {
        metadataField: 'metadata-field',
        textField: 'text-field',
        vectorField: "vector"
    },
    vectorIndexName: 'your-vector-index-name',
},
```


##### `type`<sup>Required</sup> <a name="type" id="bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string

Required.

Has to be ``"REDIS_ENTERPRISE_CLOUD"`` for Redis Enterprise Cloud Configuration.

---

### RedisFieldMapping <a name="RedisFieldMapping" id="bedrock-agents-cdk.RedisFieldMapping"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.RedisFieldMapping.Initializer"></a>

```typescript
import { RedisFieldMapping } from 'bedrock-agents-cdk'

const redisFieldMapping: RedisFieldMapping = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.RedisFieldMapping.property.metadataField">metadataField</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.RedisFieldMapping.property.textField">textField</a></code> | <code>string</code> | Required. |
| <code><a href="#bedrock-agents-cdk.RedisFieldMapping.property.vectorField">vectorField</a></code> | <code>string</code> | Required. |

---

##### `metadataField`<sup>Required</sup> <a name="metadataField" id="bedrock-agents-cdk.RedisFieldMapping.property.metadataField"></a>

```typescript
public readonly metadataField: string;
```

- *Type:* string

Required.

Metadata field that you configured in your Vector DB.
Bedrock will store metadata that is required to carry out souce attribution
and enable data ingestion and querying for OpenSearch and Pinecone and
will store raw text from your data in chunks in this field for Redis Enterprise Cloud.

---

##### `textField`<sup>Required</sup> <a name="textField" id="bedrock-agents-cdk.RedisFieldMapping.property.textField"></a>

```typescript
public readonly textField: string;
```

- *Type:* string

Required.

Text field that you configured in your Vector DB.
Bedrock will store raw text from your data in chunks in this field.

---

##### `vectorField`<sup>Required</sup> <a name="vectorField" id="bedrock-agents-cdk.RedisFieldMapping.property.vectorField"></a>

```typescript
public readonly vectorField: string;
```

- *Type:* string

Required.

Vector field that you configured in Redis Enterprise Cloud.
Bedrock will store the vector embeddings in this field.

---

### S3Configuration <a name="S3Configuration" id="bedrock-agents-cdk.S3Configuration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.S3Configuration.Initializer"></a>

```typescript
import { S3Configuration } from 'bedrock-agents-cdk'

const s3Configuration: S3Configuration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.S3Configuration.property.bucketArn">bucketArn</a></code> | <code>string</code> | Required. |

---

##### `bucketArn`<sup>Required</sup> <a name="bucketArn" id="bedrock-agents-cdk.S3Configuration.property.bucketArn"></a>

```typescript
public readonly bucketArn: string;
```

- *Type:* string

Required.

S3 bucket with files that you want to create embeddings
on for agent to make search on.

---

### VectorKnowledgeBaseConfiguration <a name="VectorKnowledgeBaseConfiguration" id="bedrock-agents-cdk.VectorKnowledgeBaseConfiguration"></a>

#### Initializer <a name="Initializer" id="bedrock-agents-cdk.VectorKnowledgeBaseConfiguration.Initializer"></a>

```typescript
import { VectorKnowledgeBaseConfiguration } from 'bedrock-agents-cdk'

const vectorKnowledgeBaseConfiguration: VectorKnowledgeBaseConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#bedrock-agents-cdk.VectorKnowledgeBaseConfiguration.property.embeddingModelArn">embeddingModelArn</a></code> | <code>string</code> | Required. |

---

##### `embeddingModelArn`<sup>Required</sup> <a name="embeddingModelArn" id="bedrock-agents-cdk.VectorKnowledgeBaseConfiguration.property.embeddingModelArn"></a>

```typescript
public readonly embeddingModelArn: string;
```

- *Type:* string
- *Default:* "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"

Required.

Embeddings model to convert your data into an embedding.

---


