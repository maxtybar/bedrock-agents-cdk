import boto3
import os
import json

agent_client = boto3.client("bedrock-agent")

knowledgebase_name = os.environ['KNOWLEDGEBASE_NAME']
role_arn = os.environ['ROLE_ARN']
knowledgebase_configuration = json.loads(os.environ['KNOWLEDGEBASE_CONFIGURATION'])
storage_configuration = json.loads(os.environ['STORAGE_CONFIGURATION'])
data_source = json.loads(os.environ['DATA_SOURCE'])
description = os.environ.get('DESCRIPTION') or 'Undefined'
removal_policy = os.environ['REMOVAL_POLICY']
region = os.environ['AWS_REGION']

def on_event(event, context):

    print(json.dumps(event))

    request_type = event['RequestType']
    if request_type == 'Create':
        return on_create(event,
                         kb_name=knowledgebase_name,
                         kb_role_arn=role_arn,
                         kb_configuration=knowledgebase_configuration,
                         kb_storage_configuration=storage_configuration,
                         kb_data_source=data_source,
                         kb_description=description)
    if request_type == 'Update':
        return on_update(event)
    if request_type == 'Delete':
        return on_delete(kb_name=knowledgebase_name,
                         kb_removal_policy=removal_policy)
    raise Exception("Invalid request type: %s" % request_type)


def on_create(event,
              kb_name,
              kb_role_arn,
              kb_configuration,
              kb_storage_configuration,
              kb_data_source,
              kb_description):

    props = event["ResourceProperties"]
    print("create new resource with props %s" % props)

    kb_data_source_configuration = kb_data_source['dataSourceConfiguration']
    kb_data_source_name = kb_data_source['name']

    knowledge_base = create_knowledge_base(kb_name=kb_name,
                                           kb_role_arn=kb_role_arn,
                                           kb_configuration=kb_configuration,
                                           kb_description=kb_description,
                                           kb_storage_configuration=kb_storage_configuration)

    data_source = create_data_source(kb_data_source_name=kb_data_source_name,
                                     knowledge_base_id=knowledge_base['knowledgeBaseId'],
                                     kb_data_source_configuration=kb_data_source_configuration)
        

    return {
        'Data': {
            'knowledgeBaseId': knowledge_base['knowledgeBaseId'],
            'knowledgeBaseArn': knowledge_base['knowledgeBaseArn'],
            'dataSourceId': data_source['dataSourceId']
        }
    }


def on_update(event):
    props = event["ResourceProperties"]
    print("Update resource with props %s" % (props))

    return {'ResourceProperties': props}


def on_delete(kb_name, 
              kb_removal_policy):
    print("Delete resource(s)")    
    # Delete knowledge base if Policy Removal was not set up
    # for either 'retain' or 'retain-on-update-or-delete'
    if (kb_removal_policy != "retain-on-update-or-delete" 
      and kb_removal_policy != "retain"):
      for kb in agent_client.list_knowledge_bases()['knowledgeBaseSummaries']:
        if kb['name'] == kb_name:
          delete_knowledge_base(knowledge_base_id=kb['knowledgeBaseId'])


def create_knowledge_base(kb_name,
                          kb_role_arn,
                          kb_configuration,
                          kb_description,
                          kb_storage_configuration):
    
    print("Storage configuration: " + json.dumps(kb_storage_configuration))
    
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

    return response['knowledgeBase']


def create_data_source(knowledge_base_id,
                       kb_data_source_configuration,
                       kb_data_source_name):

    args = {
        'name': kb_data_source_name,
        'dataSourceConfiguration': kb_data_source_configuration,
        'knowledgeBaseId': knowledge_base_id
    }

    response = agent_client.create_data_source(**args)

    return response['dataSource']


def delete_knowledge_base(knowledge_base_id):
    return agent_client.delete_knowledge_base(knowledgeBaseId=knowledge_base_id)

