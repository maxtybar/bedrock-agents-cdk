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
			_jsii_.MemberProperty{JsiiProperty: "actionGroups", GoGetter: "ActionGroups"},
			_jsii_.MemberProperty{JsiiProperty: "agentName", GoGetter: "AgentName"},
			_jsii_.MemberProperty{JsiiProperty: "agentResourceRoleArn", GoGetter: "AgentResourceRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "bedrockAgentCustomResourceRole", GoGetter: "BedrockAgentCustomResourceRole"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "foundationModel", GoGetter: "FoundationModel"},
			_jsii_.MemberProperty{JsiiProperty: "idleSessionTTLInSeconds", GoGetter: "IdleSessionTTLInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseAssociations", GoGetter: "KnowledgeBaseAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
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
			_jsii_.MemberProperty{JsiiProperty: "bedrockKnowledgeBaseCustomResourceRole", GoGetter: "BedrockKnowledgeBaseCustomResourceRole"},
			_jsii_.MemberProperty{JsiiProperty: "dataSource", GoGetter: "DataSource"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseConfiguration", GoGetter: "KnowledgeBaseConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "removalPolicy", GoGetter: "RemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "storageConfiguration", GoGetter: "StorageConfiguration"},
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
