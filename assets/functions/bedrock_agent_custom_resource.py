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
action_group_parameters_list = json.loads(os.environ['ACTION_GROUPS'])
knowledge_base_associations_list = json.loads(os.environ['KNOWLEDGE_BASE_ASSOCIATIONS'])
region = os.environ['AWS_REGION']

def on_event(event, context):
    account_id = context.invoked_function_arn.split(":")[4]
    physical_id = f"BedrockAgent-{agent_name}-cdk"

    print(json.dumps(event))

    request_type = event['RequestType']
    if request_type == 'Create':
        return on_create(event,
                         agent_name=agent_name,
                         instruction=instruction,
                         foundation_model=foundation_model,
                         agent_resource_role_arn=agent_resource_role_arn,
                         kb_associations_list=knowledge_base_associations_list,
                         agent_description=agent_description,
                         agent_session_timeout=agent_session_timeout,
                         action_group_parameters_list=action_group_parameters_list,
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
                         action_group_parameters_list=action_group_parameters_list)
    raise Exception("Invalid request type: %s" % request_type)


def on_create(event,
              agent_name,
              instruction,
              foundation_model,
              agent_resource_role_arn,
              kb_associations_list,
              agent_description,
              agent_session_timeout,
              action_group_parameters_list,
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
    for action_group_parameters in action_group_parameters_list:
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

    # Associate Knowledge Base(s) if provided
    kb_list = agent_client.list_knowledge_bases()['knowledgeBaseSummaries']
    for kb_association in kb_associations_list:
        if kb_association['knowledgeBaseName'] != 'Undefined':
            for kb in kb_list:
                if kb['name'] == kb_association['knowledgeBaseName']:
                    associate_knowledge_base(agent_id=agent_id,
                                             knowledge_base_id=kb['knowledgeBaseId'],
                                             description=kb_association['instruction'])

    return {'PhysicalResourceId': physical_id}


def on_update(event, physical_id):
    # physical_id = event["PhysicalResourceId"]
    props = event["ResourceProperties"]
    print("Update resource %s with props %s" % (physical_id, props))

    return {'PhysicalResourceId': physical_id}


def on_delete(event, 
              agent_name, 
              physical_id, 
              action_group_parameters_list):
    # physical_id = event["PhysicalResourceId"]
    print("Delete resource %s" % physical_id)
    delete_agent(agent_name=agent_name)
    # Remove lambda:allowInvoke permission
    for action_group_parameters in action_group_parameters_list:
        if action_group_parameters['s3BucketName'] != 'Undefined':
            lambda_arn = action_group_parameters['actionGroupExecutor']
            # Check if Lambda still exists
            if lambda_client.get_function(FunctionName=lambda_arn):
                lambda_remove_agent_permission(agent_name=agent_name,
                                               function_name=lambda_arn)

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

    try:
        lambda_client.add_permission(
          FunctionName=function_name,
          StatementId=f'allowInvoke-{agent_name}',
          Action='lambda:InvokeFunction',
          Principal='bedrock.amazonaws.com',
          SourceArn=f"arn:aws:bedrock:{region}:{account_id}:agent/{agent_id}",
        )
    except lambda_client.exceptions.ResourceConflictException:
        pass


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


def lambda_remove_agent_permission(agent_name, function_name):
    try:
        lambda_client.remove_permission(
          FunctionName=function_name,
          StatementId=f'allowInvoke-{agent_name}',
        )
    except lambda_client.exceptions.ResourceNotFoundException:
        pass
