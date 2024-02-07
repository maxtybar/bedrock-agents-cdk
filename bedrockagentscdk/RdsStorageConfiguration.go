package bedrockagentscdk


type RdsStorageConfiguration struct {
	// Required.
	//
	// RDS Configuration.
	//
	// Example:
	//   rdsConfiguration: {
	//       resourceArn: "arn:aws:rds:us-east-2:12345:cluster:my-aurora-cluster-1",
	//       databaseName: "mydbcluster",
	//       tableName: "mytable",
	//       credentialsSecretArn: "arn:aws:rds:us-east-2:12345:cluster:my-aurora-cluster-1",
	//       fieldMapping: {
	//           vectorField: "vectorField",
	//           textField: "text"
	//           metadataField: "metadata",
	//           primaryKeyField: "id",
	//       },
	//   },
	//
	RdsConfiguration *RdsConfiguration `field:"required" json:"rdsConfiguration" yaml:"rdsConfiguration"`
	// Required.
	//
	// Has to be ``"RDS"`` for RDS (Aurora) Configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
}

