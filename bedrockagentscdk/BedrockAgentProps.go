package bedrockagentscdk


type BedrockAgentProps struct {
	// Required.
	//
	// Name of the agent.
	AgentName *string `field:"required" json:"agentName" yaml:"agentName"`
	// Required.
	//
	// Instruction for the agent.
	// Characters length:.
	Instruction *string `field:"required" json:"instruction" yaml:"instruction"`
	// Optional.
	//
	// Action group for the agent. If specified must contain ``s3BucketName`` and ``s3ObjectKey`` with JSON
	// or yaml Open API schema, as well as Lambda ARN for the action group executor,
	// and action group name.
	ActionGroup *ActionGroup `field:"optional" json:"actionGroup" yaml:"actionGroup"`
	// Optional.
	//
	// Resource role ARN for an agent.
	// Role name must start with ``AmazonBedrockExecutionRoleForAgents_`` prefix and assumed by ``bedrock.amazonaws.com``.
	// If role is not specified the default one will be created with no attached policies to it.
	// If actionGroup is specified and the role is not, then the default created role will have an attached S3 read access
	// to the bucket provided in the actionGroup.
	AgentResourceRoleArn *string `field:"optional" json:"agentResourceRoleArn" yaml:"agentResourceRoleArn"`
	// Optional.
	//
	// Description for the agent.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// Foundation model.
	//
	// Example:
	//   - "anthropic.claude-instant-v1" or "anthropic.claude-v2"
	//
	// Default: - "anthropic.claude-v2"
	//
	FoundationModel *string `field:"optional" json:"foundationModel" yaml:"foundationModel"`
	// Optional.
	//
	// Max Session Time in seconds.
	// Default: - 3600.
	//
	IdleSessionTTLInSeconds *float64 `field:"optional" json:"idleSessionTTLInSeconds" yaml:"idleSessionTTLInSeconds"`
	// Optional.
	//
	// Knowledge base. If provided must be either of type ``openSearchServerlessConfiguration``
	// or ``pineconeConfiguration`` or ``redisEnterpriseCloudConfiguration``.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-create.html
	//
	KnowledgeBase *KnowledgeBase `field:"optional" json:"knowledgeBase" yaml:"knowledgeBase"`
}

