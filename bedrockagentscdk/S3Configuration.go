package bedrockagentscdk


type S3Configuration struct {
	// Required.
	//
	// S3 bucket with files that you want to create embeddings
	// on for agent to make search on.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Optional.
	//
	// Prefix for a bucket if your files located in a folder.
	// If you have a folder ``files``inside the bucket,
	// and the folder contains files you want to perform
	// search on, then use ``[files/]`` as an ``inclusionPrefixes``.
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

