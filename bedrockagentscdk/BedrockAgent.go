package bedrockagentscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/internal"
)

type BedrockAgent interface {
	constructs.Construct
	// `agentArn` is the ARN for the created agent.
	AgentArn() *string
	// `agentId` is the unique identifier for the created agent.
	AgentId() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BedrockAgent
type jsiiProxy_BedrockAgent struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BedrockAgent) AgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewBedrockAgent(scope constructs.Construct, name *string, props *BedrockAgentProps) BedrockAgent {
	_init_.Initialize()

	if err := validateNewBedrockAgentParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockAgent{}

	_jsii_.Create(
		"bedrock-agents-cdk.BedrockAgent",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewBedrockAgent_Override(b BedrockAgent, scope constructs.Construct, name *string, props *BedrockAgentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"bedrock-agents-cdk.BedrockAgent",
		[]interface{}{scope, name, props},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func BedrockAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"bedrock-agents-cdk.BedrockAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

