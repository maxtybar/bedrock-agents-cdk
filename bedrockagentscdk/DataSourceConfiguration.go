package bedrockagentscdk


type DataSourceConfiguration struct {
	// Required.
	//
	// S3 Configuration.
	//
	// Example:
	//     s3Configuration: {
	//       bucketArn: "yourS3BucketArn"
	//     }
	//
	S3Configuration *S3Configuration `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// Optional.
	//
	// Type of configuration.
	// Default: - "S3".
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

