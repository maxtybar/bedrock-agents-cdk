package bedrockagentscdk


type RedisEnterpriseCloudConfiguration struct {
	// Required.
	//
	// ARN of the secret defining the username, password, serverCertificate,
	// clientCertificate and clientPrivateKey to use when connecting to the Redis Enterprise Cloud database.
	// Learn more in the link below.
	// See: https://docs.redis.com/latest/rc/cloud-integrations/aws-marketplace/aws-bedrock/set-up-redis/
	//
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Required.
	//
	// The endpoint URL for your Redis Enterprise Cloud database.
	EndpointUrl *string `field:"required" json:"endpointUrl" yaml:"endpointUrl"`
	// Required.
	//
	// Field mapping consisting of ``vectorField``,
	// ``textField`` and ``metadataField``.
	FieldMapping *RedisFieldMapping `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// Required.
	//
	// Vector index name of your Redis Enterprise Cloud.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
}

