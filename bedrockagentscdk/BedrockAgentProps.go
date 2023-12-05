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
	ActionGroups *[]*ActionGroup `field:"optional" json:"actionGroups" yaml:"actionGroups"`
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
	// A list of knowledge base association objects
	// consisting of name and instruction for the associated knowledge base.
	//
	// Example:
	//   knowledgeBaseAssociations: [
	//     {
	//       knowledgeBaseName: "knowledge-base-name",
	//       instruction: "instruction-for-knowledge-base"
	//     }
	//
	KnowledgeBaseAssociations *[]*KnowledgeBaseAssociation `field:"optional" json:"knowledgeBaseAssociations" yaml:"knowledgeBaseAssociations"`
}

