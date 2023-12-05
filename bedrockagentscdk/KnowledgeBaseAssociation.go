package bedrockagentscdk


type KnowledgeBaseAssociation struct {
	// Required.
	//
	// Instruction based on the design and type of information of the knowledge base.
	// This will impact how the knowledge base interacts with the agent.
	Instruction *string `field:"required" json:"instruction" yaml:"instruction"`
	// Required.
	//
	// Name of the existing Knowledge Base that
	// you want to associate with the agent.
	KnowledgeBaseName *string `field:"required" json:"knowledgeBaseName" yaml:"knowledgeBaseName"`
}

