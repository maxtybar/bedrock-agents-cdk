//go:build no_runtime_type_checking

package bedrockagentscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateBedrockAgent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBedrockAgentParameters(scope constructs.Construct, name *string, props *BedrockAgentProps) error {
	return nil
}

