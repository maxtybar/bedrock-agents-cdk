package bedrockagentscdk


type KnowledgeBaseConfiguration struct {
	// Required.
	//
	// Type of configuration.
	// Default: - "VECTOR".
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Required.
	//
	// Embeddings model to convert your data into an embedding.
	// Default: - vectorKnowledgeBaseConfiguration: {
	// embeddingModelArn: "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"
	// }.
	//
	VectorKnowledgeBaseConfiguration *VectorKnowledgeBaseConfiguration `field:"required" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

