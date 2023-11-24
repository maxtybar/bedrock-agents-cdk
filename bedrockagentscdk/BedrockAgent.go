package bedrockagentscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/internal"
)

type BedrockAgent interface {
	constructs.Construct
	ActionGroup() *ActionGroup
	AgentName() *string
	AgentResourceRoleArn() *string
	BedrockAgentCustomResourceRole() awsiam.Role
	Description() *string
	FoundationModel() *string
	IdleSessionTTLInSeconds() *float64
	Instruction() *string
	KnowledgeBase() *KnowledgeBase
	// The tree node.
	Node() constructs.Node
	Region() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BedrockAgent
type jsiiProxy_BedrockAgent struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BedrockAgent) ActionGroup() *ActionGroup {
	var returns *ActionGroup
	_jsii_.Get(
		j,
		"actionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) AgentResourceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) BedrockAgentCustomResourceRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"bedrockAgentCustomResourceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) FoundationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) IdleSessionTTLInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTTLInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgent) KnowledgeBase() *KnowledgeBase {
	var returns *KnowledgeBase
	_jsii_.Get(
		j,
		"knowledgeBase",
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

func (j *jsiiProxy_BedrockAgent) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
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

