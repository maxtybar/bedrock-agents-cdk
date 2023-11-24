package bedrockagentscdk


type OpenSearchServerlessStorageConfiguration struct {
	// Required.
	//
	// OpenSearch Serverless Configuration.
	//
	// Example:
	//   opensearchServerlessConfiguration: {
	//       collectionArn: "arn:aws:opensearchserverless:us-east-1:123456789012:collection/my-collection",
	//       fieldMapping: {
	//           textField: "text",
	//           metadataField: "metadata",
	//           vectorField: "vector"
	//       },
	//       vectorIndexName: "my-index",
	//   },
	//
	OpensearchServerlessConfiguration *OpenSearchServerlessConfiguration `field:"required" json:"opensearchServerlessConfiguration" yaml:"opensearchServerlessConfiguration"`
	// Required.
	//
	// Has to be ``"OPENSEARCH_SERVERLESS"`` for Opensearch Serverless Configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
}

