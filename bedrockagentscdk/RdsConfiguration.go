package bedrockagentscdk


type RdsConfiguration struct {
	// Required.
	//
	// The Secret ARN of you Amazon Aurora DB cluster.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Required.
	//
	// The name of your Database.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Required.
	//
	// Field mapping consisting of ``vectorField``, ``primaryKeyField``,
	// ``textField`` and ``metadataField``.
	FieldMapping *RdsFieldMapping `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// Required.
	//
	// The ARN of your Amazon Aurora DB cluster.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Required.
	//
	// The Table Name of your Amazon Aurora DB cluster.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

