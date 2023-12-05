import * as cdk from 'aws-cdk-lib';

export interface ActionGroup {
  /**
     * Required. Action group name.
     */
  readonly actionGroupName: string;
  /**
     * Required. Lambda function ARN that will be used in this action group.
     * The provided Lambda function will be assigned a resource-based policy
     * to allow access from the newly created agent.
     */
  readonly actionGroupExecutor: string;
  /**
     * Required. S3 bucket name where Open API schema is stored.
     * Bucket must be in the same region where agent is created.
     */
  readonly s3BucketName: string;
  /**
     * Required. S3 bucket key for the actual schema.
     * Must be either JSON or yaml file.
     */
  readonly s3ObjectKey: string;
  /**
     * Optional. Description for the action group.
     */
  readonly desription?: string;

}

export interface BedrockKnowledgeBaseProps {
  /**
    * Required. Name of the KnowledgeBase.
    */
  readonly name: string;
  /**
     * Required. Resource role ARN for a knowledge base.
     * Role name must start with ``AmazonBedrockExecutionRoleForKnowledgeBase_`` prefix and assumed by ``bedrock.amazonaws.com``.
     * Role must have access to the S3 bucket used as a data source as a knowledge base.
     * If you use OpenSearch serverless, the role must have ``aoss:APIAccessAll`` policy attached to it
     * allowing it to make API calls against your collection's data plane. Your collection
     * must also allow data access from KnowledgeBase role.
     * See more here @see https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html
     */
  readonly roleArn: string;
  /**
    * Optional. KnowledgeBase configuration.
    * @default
    * type: "VECTOR",
    * vectorKnowledgeBaseConfiguration: {
    *     embeddingModelArn: "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"
    * }
    */
  readonly knowledgeBaseConfiguration?: KnowledgeBaseConfiguration;
  /**
    * Required. KnowledgeBase storage configuration.
    * Has to be either ``opensearchServerlessConfiguration`` or
    * ``pineconeConfiguration`` or ``redisEnterpriseCloudConfiguration``
    * and respective type field mapping.
    */
  readonly storageConfiguration: OpenSearchServerlessStorageConfiguration |
  RedisEnterpriseCloudStorageConfiguration | PineconeStorageConfiguration;
  /**
     * Required. Data source configuration.
     */
  readonly dataSource: DataSource;
  /**
     * Optional. Description for the knowledge base.
     */
  readonly description?: string;
  /**
     * Optional. Removal Policy. If you want to keep your Knowledge Base intact
     * in case you destroy this CDK make sure you set removalPolicy to
     * ``cdk.RemovalPolicy.RETAIN``. By default your Knowledge Base will be
     * deleted along with the agent.
     * @default - cdk.RemovalPolicy.DESTROY
     */
  readonly removalPolicy?: cdk.RemovalPolicy;

}

export interface DataSource {
  /**
    * Optional. Name for your data source.
    * @default - `MyDataSource-${agentName}` or `MyDataSource-${knowledgeBaseName}`
    */
  readonly name?: string;
  /**
      * Required. Data Source Configuration.
      * @example
      dataSourceConfiguration = {
        s3Configuration: {
          bucketArn: "yourS3BucketArn",
        },
          "type": "S3"
      }
      */
  readonly dataSourceConfiguration: DataSourceConfiguration;
}

export interface DataSourceConfiguration {
  /**
      * Required. S3 Configuration.
      * @example
          s3Configuration: {
            bucketArn: "yourS3BucketArn"
          }
      */
  readonly s3Configuration: S3Configuration;
  /**
      * Optional. Type of configuration.
      * @default - "S3"
      */
  readonly type?: string;

}

export interface S3Configuration {
  /**
      * Required. S3 bucket with files that you want to create embeddings
      * on for agent to make search on.
      */
  readonly bucketArn: string;
  /**
      * Optional. Prefix for a bucket if your files located in a folder.
      * If you have a folder ``files``inside the bucket,
      * and the folder contains files you want to perform
      * search on, then use ``[files/]`` as an ``inclusionPrefixes``.
      */
  readonly inclusionPrefixes?: string[];

}

export interface KnowledgeBaseConfiguration {
  /**
    * Required. Type of configuration.
    * @default - "VECTOR"
    */
  readonly type: string;
  /**
    * Required. Embeddings model to convert your data into an embedding.
    * @default - vectorKnowledgeBaseConfiguration: {
    *     embeddingModelArn: "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"
    * }
    */
  readonly vectorKnowledgeBaseConfiguration: VectorKnowledgeBaseConfiguration;

}

export interface VectorKnowledgeBaseConfiguration {
  /**
    * Required. Embeddings model to convert your data into an embedding.
    * @default - "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"
    */
  readonly embeddingModelArn: string;

}

export interface OpenSearchServerlessConfiguration {
  /**
    * Required. ARN of your OpenSearch Serverless Collection.
    */
  readonly collectionArn: string;
  /**
    * Required. Vector index name of your OpenSearch Serverless Collection.
    */
  readonly vectorIndexName: string;
  /**
    * Required. Field mapping consisting of ``vectorField``,
    * ``textField`` and ``metadataField``.
    */
  readonly fieldMapping: OpenSearchFieldMapping;

}

export interface PineconeConfiguration {
  /**
    * Required. Connection string for your Pinecone index management page.
    */
  readonly connectionString: string;
  /**
    * Required. ARN of the secret containing the API Key to use when connecting to the Pinecone database.
    * Learn more in the link below.
    * @see https://www.pinecone.io/blog/amazon-bedrock-integration/
    */
  readonly credentialsSecretArn: string;
  /**
    * Optional. Name space that will be used for writing new data to your Pinecone database.
    */
  readonly namespace?: string;
  /**
    * Required. Field mapping consisting of ``textField`` and ``metadataField``.
    */
  readonly fieldMapping: PineconeFieldMapping;

}

export interface RedisEnterpriseCloudConfiguration {
  /**
    * Required. The endpoint URL for your Redis Enterprise Cloud database.
    */
  readonly endpointUrl: string;
  /**
    * Required. ARN of the secret defining the username, password, serverCertificate,
    * clientCertificate and clientPrivateKey to use when connecting to the Redis Enterprise Cloud database.
    * Learn more in the link below.
    *  @see https://docs.redis.com/latest/rc/cloud-integrations/aws-marketplace/aws-bedrock/set-up-redis/
    */
  readonly credentialsSecretArn: string;
  /**
    * Required. Vector index name of your Redis Enterprise Cloud.
    */
  readonly vectorIndexName: string;
  /**
    * Required. Field mapping consisting of ``vectorField``,
    * ``textField`` and ``metadataField``.
    */
  readonly fieldMapping: RedisFieldMapping;

}

export interface BaseFieldMapping {
  /**
    * Required. Metadata field that you configured in your Vector DB.
    * Bedrock will store metadata that is required to carry out souce attribution
    * and enable data ingestion and querying for OpenSearch and Pinecone and
    * will store raw text from your data in chunks in this field for Redis Enterprise Cloud.
    */
  readonly metadataField: string;
  /**
    * Required. Text field that you configured in your Vector DB.
    * Bedrock will store raw text from your data in chunks in this field.
    */
  readonly textField: string;
}

export interface OpenSearchFieldMapping extends BaseFieldMapping {
  /**
    * Required. Vector field that you configured in OpenSearch.
    * Bedrock will store the vector embeddings in this field.
    */
  readonly vectorField: string;
}

export interface PineconeFieldMapping extends BaseFieldMapping {}

export interface RedisFieldMapping extends BaseFieldMapping {
  /**
    * Required. Vector field that you configured in Redis Enterprise Cloud.
    * Bedrock will store the vector embeddings in this field.
    */
  readonly vectorField: string;
}

export interface OpenSearchServerlessStorageConfiguration {
  /**
     * Required. OpenSearch Serverless Configuration.
     * @example
     * opensearchServerlessConfiguration: {
     *     collectionArn: "arn:aws:opensearchserverless:us-east-1:123456789012:collection/my-collection",
     *     fieldMapping: {
     *         textField: "text",
     *         metadataField: "metadata",
     *         vectorField: "vector"
     *     },
     *     vectorIndexName: "my-index",
     * },
     */
  readonly opensearchServerlessConfiguration: OpenSearchServerlessConfiguration;
  /**
    * Required. Has to be ``"OPENSEARCH_SERVERLESS"`` for Opensearch Serverless Configuration.
    */
  readonly type: 'OPENSEARCH_SERVERLESS';
}

export interface PineconeStorageConfiguration {
  /**
    * Required. Pinecone Configuration.
    * @example
    * pineconeConfiguration: {
    *     credentialsSecretArn: 'arn:aws:secretsmanager:your-region:123456789098:secret:apiKey';
    *     connectionString: 'https://your-connection-string.pinecone.io';
    *     fieldMapping: {
    *         metadataField: 'metadata-field',
    *         textField: 'text-field'
    *     },
    * },
    */
  readonly pineconeConfiguration: PineconeConfiguration;
  /**
    * Required. Has to be ``"PINECONE"`` for Pinecone Configuration.
    */
  readonly type: 'PINECONE';
}

export interface RedisEnterpriseCloudStorageConfiguration {
  /**
    * Required. Redis Enterprise Cloud Configuration.
    * @example
    * redisEnterpriseCloudConfiguration: {
    *     credentialsSecretArn: 'arn:aws:secretsmanager:your-region:123456789098:secret:apiKey';
    *     endpointUrl: 'your-endpoint-url';
    *     fieldMapping: {
    *         metadataField: 'metadata-field',
    *         textField: 'text-field',
    *         vectorField: "vector"
    *     },
    *     vectorIndexName: 'your-vector-index-name',
    * },
    */
  readonly redisEnterpriseCloudConfiguration: RedisEnterpriseCloudConfiguration;
  /**
    * Required. Has to be ``"REDIS_ENTERPRISE_CLOUD"`` for Redis Enterprise Cloud Configuration.
    */
  readonly type: 'REDIS_ENTERPRISE_CLOUD';
}

export interface KnowledgeBaseAssociation {
  /**
     * Required. Name of the existing Knowledge Base that
     * you want to associate with the agent.
     */
  readonly knowledgeBaseName: string;
  /**
     * Required. Instruction based on the design and type of information of the knowledge base.
     * This will impact how the knowledge base interacts with the agent.
     * @max 150 characters
     */
  readonly instruction: string;
}

// Create a BedrockAgentProps interface for BedrockAgent
export interface BedrockAgentProps {
  /**
    * Required. Name of the agent.
    */
  readonly agentName: string;
  /**
    * Required. Instruction for the agent.
    * Characters length:
    * @min 40
    * @max 1200
    */
  readonly instruction: string;
  /**
    * Optional. Foundation model.
    * @example - "anthropic.claude-instant-v1" or "anthropic.claude-v2"
    * @default - "anthropic.claude-v2"
    */
  readonly foundationModel?: string;
  /**
    * Optional. Resource role ARN for an agent.
    * Role name must start with ``AmazonBedrockExecutionRoleForAgents_`` prefix and assumed by ``bedrock.amazonaws.com``.
    * If role is not specified the default one will be created with no attached policies to it.
    * If actionGroup is specified and the role is not, then the default created role will have an attached S3 read access
    * to the bucket provided in the actionGroup.
    */
  readonly agentResourceRoleArn?: string;
  /**
    * Optional. Description for the agent.
    */
  readonly description?: string;
  /**
    * Optional. Max Session Time in seconds.
    * @type number
    * @min 60
    * @max 3600
    * @default - 3600
    */
  readonly idleSessionTTLInSeconds?: number;
  /**
    * Optional. Action group for the agent. If specified must contain ``s3BucketName`` and ``s3ObjectKey`` with JSON
    * or yaml Open API schema, as well as Lambda ARN for the action group executor,
    * and action group name.
    */
  readonly actionGroups?: ActionGroup[];
  /**
    * Optional. A list of knowledge base association objects
    * consisting of name and instruction for the associated knowledge base.
    * @example
    * knowledgeBaseAssociations: [
    *   {
    *     knowledgeBaseName: "knowledge-base-name",
    *     instruction: "instruction-for-knowledge-base"
    *   }
    */
  readonly knowledgeBaseAssociations?: KnowledgeBaseAssociation[];
}
