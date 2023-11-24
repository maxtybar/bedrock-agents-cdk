package bedrockagentscdk


type OpenSearchServerlessConfiguration struct {
	// Required.
	//
	// ARN of your OpenSearch Serverless Collection.
	CollectionArn *string `field:"required" json:"collectionArn" yaml:"collectionArn"`
	// Required.
	//
	// Field mapping consisting of ``vectorField``,
	// ``textField`` and ``metadataField``.
	FieldMapping *OpenSearchFieldMapping `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// Required.
	//
	// Vector index name of your OpenSearch Serverless Collection.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
}

