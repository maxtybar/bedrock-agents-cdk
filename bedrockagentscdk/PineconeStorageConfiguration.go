package bedrockagentscdk


type PineconeStorageConfiguration struct {
	// Required.
	//
	// Pinecone Configuration.
	//
	// Example:
	//   pineconeConfiguration: {
	//       credentialsSecretArn: 'arn:aws:secretsmanager:your-region:123456789098:secret:apiKey';
	//       connectionString: 'https://your-connection-string.pinecone.io';
	//       fieldMapping: {
	//           metadataField: 'metadata-field',
	//           textField: 'text-field'
	//       },
	//   },
	//
	PineconeConfiguration *PineconeConfiguration `field:"required" json:"pineconeConfiguration" yaml:"pineconeConfiguration"`
	// Required.
	//
	// Has to be ``"PINECONE"`` for Pinecone Configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
}

