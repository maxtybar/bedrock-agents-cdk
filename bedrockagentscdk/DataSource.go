package bedrockagentscdk


type DataSource struct {
	// Required.
	//
	// Data Source Configuration.
	//
	// Example:
	//   dataSourceConfiguration = {
	//   s3Configuration: {
	//     bucketArn: "yourS3BucketArn",
	//   },
	//     "type": "S3"
	//   }
	//
	DataSourceConfiguration *DataSourceConfiguration `field:"required" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// Optional.
	//
	// Name for your data source.
	// Default: - `MyDataSource-${agentName}`.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

