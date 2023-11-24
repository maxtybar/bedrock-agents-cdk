package bedrockagentscdk


type RedisEnterpriseCloudStorageConfiguration struct {
	// Required.
	//
	// Redis Enterprise Cloud Configuration.
	//
	// Example:
	//   redisEnterpriseCloudConfiguration: {
	//       credentialsSecretArn: 'arn:aws:secretsmanager:your-region:123456789098:secret:apiKey';
	//       endpointUrl: 'your-endpoint-url';
	//       fieldMapping: {
	//           metadataField: 'metadata-field',
	//           textField: 'text-field',
	//           vectorField: "vector"
	//       },
	//       vectorIndexName: 'your-vector-index-name',
	//   },
	//
	RedisEnterpriseCloudConfiguration *RedisEnterpriseCloudConfiguration `field:"required" json:"redisEnterpriseCloudConfiguration" yaml:"redisEnterpriseCloudConfiguration"`
	// Required.
	//
	// Has to be ``"REDIS_ENTERPRISE_CLOUD"`` for Redis Enterprise Cloud Configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
}

