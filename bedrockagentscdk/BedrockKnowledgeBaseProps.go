package bedrockagentscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type BedrockKnowledgeBaseProps struct {
	// Required.
	//
	// Data source configuration.
	DataSource *DataSource `field:"required" json:"dataSource" yaml:"dataSource"`
	// Required.
	//
	// Name of the KnowledgeBase.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required.
	//
	// Resource role ARN for a knowledge base.
	// Role name must start with ``AmazonBedrockExecutionRoleForKnowledgeBase_`` prefix and assumed by ``bedrock.amazonaws.com``.
	// Role must have access to the S3 bucket used as a data source as a knowledge base.
	// If you use OpenSearch serverless, the role must have ``aoss:APIAccessAll`` policy attached to it
	// allowing it to make API calls against your collection's data plane. Your collection
	// must also allow data access from KnowledgeBase role.
	// See more here @see https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Required.
	//
	// KnowledgeBase storage configuration.
	// Has to be either ``opensearchServerlessConfiguration``,
	// ``pineconeConfiguration``, ``redisEnterpriseCloudConfiguration`` or `rdsConfiguration`
	// and respective type field mapping.
	StorageConfiguration interface{} `field:"required" json:"storageConfiguration" yaml:"storageConfiguration"`
	// Optional.
	//
	// Description for the knowledge base.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// KnowledgeBase configuration.
	// Default: type: "VECTOR",
	// vectorKnowledgeBaseConfiguration: {
	//     embeddingModelArn: "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"
	// }.
	//
	KnowledgeBaseConfiguration *KnowledgeBaseConfiguration `field:"optional" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// Optional.
	//
	// Removal Policy. If you want to keep your Knowledge Base intact
	// in case you destroy this CDK make sure you set removalPolicy to
	// ``cdk.RemovalPolicy.RETAIN``. By default your Knowledge Base will be
	// deleted along with the agent.
	// Default: - cdk.RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

