package bedrockagentscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/internal"
)

type BedrockKnowledgeBase interface {
	constructs.Construct
	// `dataSourceId` is the unique identifier for the created data source.
	DataSourceId() *string
	// `knowledgeBaseArn` is the ARN for the created knowledge base.
	KnowledgeBaseArn() *string
	// `knowledgeBaseId` is the unique identifier for the created knowledge base.
	KnowledgeBaseId() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BedrockKnowledgeBase
type jsiiProxy_BedrockKnowledgeBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BedrockKnowledgeBase) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) KnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewBedrockKnowledgeBase(scope constructs.Construct, name *string, props *BedrockKnowledgeBaseProps) BedrockKnowledgeBase {
	_init_.Initialize()

	if err := validateNewBedrockKnowledgeBaseParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockKnowledgeBase{}

	_jsii_.Create(
		"bedrock-agents-cdk.BedrockKnowledgeBase",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewBedrockKnowledgeBase_Override(b BedrockKnowledgeBase, scope constructs.Construct, name *string, props *BedrockKnowledgeBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"bedrock-agents-cdk.BedrockKnowledgeBase",
		[]interface{}{scope, name, props},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func BedrockKnowledgeBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockKnowledgeBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"bedrock-agents-cdk.BedrockKnowledgeBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

