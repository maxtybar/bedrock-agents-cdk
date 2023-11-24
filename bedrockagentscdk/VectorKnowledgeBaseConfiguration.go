package bedrockagentscdk


type VectorKnowledgeBaseConfiguration struct {
	// Required.
	//
	// Embeddings model to convert your data into an embedding.
	// Default: - "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"
	//
	EmbeddingModelArn *string `field:"required" json:"embeddingModelArn" yaml:"embeddingModelArn"`
}

