package bedrockagentscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maxtybar/bedrock-agents-cdk/bedrockagentscdk/internal"
)

type BedrockKnowledgeBase interface {
	constructs.Construct
	BedrockKnowledgeBaseCustomResourceRole() awsiam.Role
	DataSource() *DataSource
	Description() *string
	KnowledgeBaseConfiguration() *KnowledgeBaseConfiguration
	Name() *string
	// The tree node.
	Node() constructs.Node
	Region() *string
	RemovalPolicy() awscdk.RemovalPolicy
	RoleArn() *string
	StorageConfiguration() interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BedrockKnowledgeBase
type jsiiProxy_BedrockKnowledgeBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BedrockKnowledgeBase) BedrockKnowledgeBaseCustomResourceRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"bedrockKnowledgeBaseCustomResourceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) DataSource() *DataSource {
	var returns *DataSource
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) KnowledgeBaseConfiguration() *KnowledgeBaseConfiguration {
	var returns *KnowledgeBaseConfiguration
	_jsii_.Get(
		j,
		"knowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
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

func (j *jsiiProxy_BedrockKnowledgeBase) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) RemovalPolicy() awscdk.RemovalPolicy {
	var returns awscdk.RemovalPolicy
	_jsii_.Get(
		j,
		"removalPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBase) StorageConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConfiguration",
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

