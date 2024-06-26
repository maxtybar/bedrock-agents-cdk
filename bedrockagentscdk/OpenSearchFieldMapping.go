package bedrockagentscdk


type OpenSearchFieldMapping struct {
	// Required.
	//
	// Metadata field that you configured in your Vector DB.
	// Bedrock will store metadata that is required to carry out source attribution
	// and enable data ingestion and querying.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// Required.
	//
	// Text field that you configured in your Vector DB.
	// Bedrock will store raw text from your data in chunks in this field.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
	// Required.
	//
	// Vector field that you configured in OpenSearch.
	// Bedrock will store the vector embeddings in this field.
	VectorField *string `field:"required" json:"vectorField" yaml:"vectorField"`
}

