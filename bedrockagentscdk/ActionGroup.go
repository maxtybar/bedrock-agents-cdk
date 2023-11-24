package bedrockagentscdk


type ActionGroup struct {
	// Required.
	//
	// Lambda function ARN that will be used in this action group.
	// The provided Lambda function will be assigned a resource-based policy
	// to allow access from the newly created agent.
	ActionGroupExecutor *string `field:"required" json:"actionGroupExecutor" yaml:"actionGroupExecutor"`
	// Required.
	//
	// Action group name.
	ActionGroupName *string `field:"required" json:"actionGroupName" yaml:"actionGroupName"`
	// Required.
	//
	// S3 bucket name where Open API schema is stored.
	// Bucket must be in the same region where agent is created.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Required.
	//
	// S3 bucket key for the actual schema.
	// Must be either JSON or yaml file.
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// Optional.
	//
	// Description for the action group.
	Desription *string `field:"optional" json:"desription" yaml:"desription"`
}

