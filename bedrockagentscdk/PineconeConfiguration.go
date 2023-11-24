package bedrockagentscdk


type PineconeConfiguration struct {
	// Required.
	//
	// Connection string for your Pinecone index management page.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Required.
	//
	// ARN of the secret containing the API Key to use when connecting to the Pinecone database.
	// Learn more in the link below.
	// See: https://www.pinecone.io/blog/amazon-bedrock-integration/
	//
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Required.
	//
	// Field mapping consisting of``textField`` and ``metadataField``.
	FieldMapping *PineconeFieldMapping `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// Optional.
	//
	// Name space that will be used for writing new data to your Pinecone database.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

