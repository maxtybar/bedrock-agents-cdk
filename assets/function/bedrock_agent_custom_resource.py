import boto3
import os
import time
import json

agent_client = boto3.client("bedrock-agent")
lambda_client = boto3.client("lambda")

agent_name = os.environ['AGENT_NAME']
instruction = os.environ['INSTRUCTION']
foundation_model = os.environ['FOUNDATION_MODEL']
agent_resource_role_arn = os.environ['AGENT_RESOURCE_ROLE_ARN']
agent_description = os.environ['DESCRIPTION']
agent_session_timeout = os.environ['IDLE_SESSION_TTL_IN_SECONDS']
action_group_parameters = json.loads(os.environ['ACTION_GROUP'])
knowledge_base_parameters = json.loads(os.environ['KNOWLEDGE_BASE'])
region = os.environ['AWS_REGION']

def on_event(event, context):
    account_id = context.invoked_function_arn.split(":")[4]
    physical_id = f"BedrockAgent-{agent_name}-cdk"

    print(json.dumps(event))
    kb_storage_configuration = knowledge_base_parameters['storageConfiguration']
    print("Storage configuration: " + json.dumps(kb_storage_configuration))

    request_type = event['RequestType']
    if request_type == 'Create':
        return on_create(event,
                         agent_name=agent_name,
                         instruction=instruction,
                         foundation_model=foundation_model,
                         agent_resource_role_arn=agent_resource_role_arn,
                         knowledge_base_parameters=knowledge_base_parameters,
                         agent_description=agent_description,
                         agent_session_timeout=agent_session_timeout,
                         action_group_parameters=action_group_parameters,
                         physical_id=physical_id,
                         region=region,
                         account_id=account_id)
    if request_type == 'Update':
        return on_update(event,
                         physical_id=physical_id)
    if request_type == 'Delete':
        return on_delete(event,
                         agent_name=agent_name,
                         physical_id=physical_id,
                         action_group_parameters=action_group_parameters,
                         knowledge_base_parameters=knowledge_base_parameters)
    raise Exception("Invalid request type: %s" % request_type)


def on_create(event,
              agent_name,
              instruction,
              foundation_model,
              agent_resource_role_arn,
              knowledge_base_parameters,
              agent_description,
              agent_session_timeout,
              action_group_parameters,
              physical_id,
              region,
              account_id):

    props = event["ResourceProperties"]
    print("create new resource with props %s" % props)

    agent_id = create_agent(agent_name=agent_name,
                            agent_resource_role_arn=agent_resource_role_arn,
                            foundation_model=foundation_model,
                            agent_description=agent_description,
                            agent_session_timeout=agent_session_timeout,
                            instruction=instruction)
    # Pause to make sure agents has been created
    time.sleep(15)
    if action_group_parameters['s3BucketName'] != 'Undefined':

        action_group_name = action_group_parameters['actionGroupName']
        lambda_arn = action_group_parameters['actionGroupExecutor']
        s3_bucket_name = action_group_parameters['s3BucketName']
        s3_bucket_key = action_group_parameters['s3ObjectKey']
        action_group_description = action_group_parameters.get(
            'description', 'Undefined')

        lambda_add_agent_permission(agent_id=agent_id,
                                    agent_name=agent_name,
                                    function_name=lambda_arn,
                                    region=region,
                                    account_id=account_id)
        # Pause to make sure policy was attached
        time.sleep(10)

        create_agent_action_group(action_group_name=action_group_name,
                                  lambda_arn=lambda_arn,
                                  bucket_name=s3_bucket_name,
                                  key=s3_bucket_key,
                                  action_group_description=action_group_description,
                                  agent_id=agent_id)

    if knowledge_base_parameters['roleArn'] != 'Undefined':
        kb_name = knowledge_base_parameters['name']
        kb_role_arn = knowledge_base_parameters['roleArn']
        kb_configuration = knowledge_base_parameters['knowledgeBaseConfiguration']
        kb_storage_configuration = knowledge_base_parameters['storageConfiguration']
        kb_data_source_configuration = knowledge_base_parameters['dataSource'][
            'dataSourceConfiguration']
        kb_data_source_name = knowledge_base_parameters['dataSource']['name']
        kb_instruction = knowledge_base_parameters['instruction']
        kb_description = knowledge_base_parameters.get(
            'description', 'Undefined')

        knowledge_base_id = create_knowledge_base(kb_name=kb_name,
                                                  kb_role_arn=kb_role_arn,
                                                  kb_configuration=kb_configuration,
                                                  kb_description=kb_description,
                                                  kb_storage_configuration=kb_storage_configuration)

        create_data_source(kb_data_source_name=kb_data_source_name,
                           knowledge_base_id=knowledge_base_id,
                           kb_data_source_configuration=kb_data_source_configuration)
        
        associate_knowledge_base(agent_id=agent_id,
                                 knowledge_base_id=knowledge_base_id,
                                 description=kb_instruction)

    return {'PhysicalResourceId': physical_id}


def on_update(event, physical_id):
    # physical_id = event["PhysicalResourceId"]
    props = event["ResourceProperties"]
    print("Update resource %s with props %s" % (physical_id, props))

    return {'PhysicalResourceId': physical_id}


def on_delete(event, 
              agent_name, 
              physical_id, 
              action_group_parameters,
              knowledge_base_parameters):
    # physical_id = event["PhysicalResourceId"]
    print("Delete resource %s" % physical_id)
    delete_agent(agent_name=agent_name)
    # Remove lambda:allowInvoke permission
    if action_group_parameters['s3BucketName'] != 'Undefined':
        lambda_arn = action_group_parameters['actionGroupExecutor']
        # Check if Lambda still exists
        if lambda_client.get_function(FunctionName=lambda_arn):
            lambda_remove_agent_permission(agent_name=agent_name,
                                           function_name=lambda_arn)
            
    if knowledge_base_parameters['roleArn'] != 'Undefined':
        # Delete knowledge base if Policy Removal was not set up
        # for either 'retain' or 'retain-on-update-or-delete'
        if (knowledge_base_parameters['removalPolicy'] != "retain-on-update-or-delete" 
          and knowledge_base_parameters['removalPolicy'] != "retain"):
          for kb in agent_client.list_knowledge_bases()['knowledgeBaseSummaries']:
            if kb['name'] == knowledge_base_parameters['name']:
              delete_knowledge_base(knowledge_base_id=kb['knowledgeBaseId'])

    return {'PhysicalResourceId': physical_id}


def create_agent(agent_name,
                 agent_resource_role_arn,
                 foundation_model,
                 agent_description,
                 agent_session_timeout,
                 instruction):

    args = {
        'agentName': agent_name,
        'agentResourceRoleArn': agent_resource_role_arn,
        'foundationModel': foundation_model,
        'idleSessionTTLInSeconds': int(agent_session_timeout),
        'instruction': instruction,
        'description': agent_description
    }

    if args['description'] == 'Undefined':
        args.pop('description')

    response = agent_client.create_agent(**args)

    return response['agent']['agentId']


def lambda_add_agent_permission(agent_name, function_name,
                                region, account_id, agent_id):

    return lambda_client.add_permission(
        FunctionName=function_name,
        StatementId=f'allowInvoke-{agent_name}',
        Action='lambda:InvokeFunction',
        Principal='bedrock.amazonaws.com',
        SourceArn=f"arn:aws:bedrock:{region}:{account_id}:agent/{agent_id}",
    )


def create_knowledge_base(kb_name,
                          kb_role_arn,
                          kb_configuration,
                          kb_description,
                          kb_storage_configuration):
    
    print("Storage configuration: " + json.dumps(kb_storage_configuration))
    
    # Check if it is Pinecone and add 
    # empty namespace if it was not provided
    if "pineconeConfiguration" in kb_storage_configuration:
      pinecone_config = kb_storage_configuration["pineconeConfiguration"]
      if "namespace" not in pinecone_config:
        pinecone_config["namespace"] = ''
        kb_storage_configuration["pineconeConfiguration"] = pinecone_config

    args = {
        'name': kb_name,
        'roleArn': kb_role_arn,
        'knowledgeBaseConfiguration': kb_configuration,
        'storageConfiguration': kb_storage_configuration,
        'description': kb_description
    }

    if args['description'] == 'Undefined':
        args.pop('description')

    response = agent_client.create_knowledge_base(**args)

    return response['knowledgeBase']['knowledgeBaseId']


def create_data_source(knowledge_base_id,
                       kb_data_source_configuration,
                       kb_data_source_name):

    args = {
        'name': kb_data_source_name,
        'dataSourceConfiguration': kb_data_source_configuration,
        'knowledgeBaseId': knowledge_base_id
    }

    response = agent_client.create_data_source(**args)

    return response


def associate_knowledge_base(agent_id, 
                             knowledge_base_id,
                             description):
    
    args = {
        'agentId': agent_id,
        'agentVersion': 'DRAFT',
        'knowledgeBaseId': knowledge_base_id,
        'description': description
    }

    agent_kb_description = agent_client.associate_agent_knowledge_base(**args)

    return agent_kb_description


def create_agent_action_group(agent_id,
                              lambda_arn,
                              bucket_name,
                              key,
                              action_group_description,
                              action_group_name):

    args = {
        'agentId': agent_id,
        'actionGroupExecutor': {
            'lambda': lambda_arn,
        },
        'actionGroupName': action_group_name,
        'agentVersion': 'DRAFT',
        'apiSchema': {
            's3': {
                's3BucketName': bucket_name,
                's3ObjectKey': key
            }
        },
        'description': action_group_description
    }

    if args['description'] == 'Undefined':
        args.pop('description')

    return agent_client.create_agent_action_group(**args)


def delete_agent(agent_name):
    # Get list of all agents
    response = agent_client.list_agents()
    print('This is agent name from delete: ', agent_name)
    # Find agent with the given name
    for agent in response["agentSummaries"]:
        if agent["agentName"] == agent_name:
            agent_id = agent["agentId"]
            return agent_client.delete_agent(agentId=agent_id)
    
    return None


def delete_knowledge_base(knowledge_base_id):
    return agent_client.delete_knowledge_base(knowledgeBaseId=knowledge_base_id)


def lambda_remove_agent_permission(agent_name, function_name):
    return lambda_client.remove_permission(
        FunctionName=function_name,
        StatementId=f'allowInvoke-{agent_name}',
    )
