package bedrockagentscdk


type S3Configuration struct {
	// Required.
	//
	// S3 bucket with files that you want to create embeddings
	// on for agent to make search on.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
}

