//go:build no_runtime_type_checking

package bedrockagentscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateBedrockKnowledgeBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBedrockKnowledgeBaseParameters(scope constructs.Construct, name *string, props *BedrockKnowledgeBaseProps) error {
	return nil
}

