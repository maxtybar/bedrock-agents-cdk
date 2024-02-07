// bedrock-agents-cdk
package bedrockagentscdk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.ActionGroup",
		reflect.TypeOf((*ActionGroup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.BaseFieldMapping",
		reflect.TypeOf((*BaseFieldMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"bedrock-agents-cdk.BedrockAgent",
		reflect.TypeOf((*BedrockAgent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BedrockAgent{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.BedrockAgentProps",
		reflect.TypeOf((*BedrockAgentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"bedrock-agents-cdk.BedrockKnowledgeBase",
		reflect.TypeOf((*BedrockKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BedrockKnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.BedrockKnowledgeBaseProps",
		reflect.TypeOf((*BedrockKnowledgeBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.DataSource",
		reflect.TypeOf((*DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.DataSourceConfiguration",
		reflect.TypeOf((*DataSourceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.KnowledgeBaseAssociation",
		reflect.TypeOf((*KnowledgeBaseAssociation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.KnowledgeBaseConfiguration",
		reflect.TypeOf((*KnowledgeBaseConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.OpenSearchFieldMapping",
		reflect.TypeOf((*OpenSearchFieldMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.OpenSearchServerlessConfiguration",
		reflect.TypeOf((*OpenSearchServerlessConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.OpenSearchServerlessStorageConfiguration",
		reflect.TypeOf((*OpenSearchServerlessStorageConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.PineconeConfiguration",
		reflect.TypeOf((*PineconeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.PineconeFieldMapping",
		reflect.TypeOf((*PineconeFieldMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.PineconeStorageConfiguration",
		reflect.TypeOf((*PineconeStorageConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.RdsConfiguration",
		reflect.TypeOf((*RdsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.RdsFieldMapping",
		reflect.TypeOf((*RdsFieldMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.RdsStorageConfiguration",
		reflect.TypeOf((*RdsStorageConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.RedisEnterpriseCloudConfiguration",
		reflect.TypeOf((*RedisEnterpriseCloudConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.RedisEnterpriseCloudStorageConfiguration",
		reflect.TypeOf((*RedisEnterpriseCloudStorageConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.RedisFieldMapping",
		reflect.TypeOf((*RedisFieldMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.S3Configuration",
		reflect.TypeOf((*S3Configuration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"bedrock-agents-cdk.VectorKnowledgeBaseConfiguration",
		reflect.TypeOf((*VectorKnowledgeBaseConfiguration)(nil)).Elem(),
	)
}
