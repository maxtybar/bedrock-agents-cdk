import {
  App, Stack, assertions,
} from 'aws-cdk-lib';
import { BedrockAgent } from '../src';


test('Match snapshot', () => {
  const app = new App();
  const stack = new Stack(app, 'BedrockAgentStack');

  new BedrockAgent(stack, 'BedrockAgentStack', {
    agentName: 'MyAgentTestName',
    instruction: 'This is a template instruction for my agent. You were created by AWS TypeScript CDK.',
    foundationModel: 'anthropic.claude-instant-v1',
    description: 'My custom description for an agent.',
  });

  const template = assertions.Template.fromStack(stack);
  expect(template).toMatchSnapshot();
});

