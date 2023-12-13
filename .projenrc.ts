import { awscdk } from 'projen';
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Max Tybar',
  authorAddress: 'maxtybar@gmail.com',
  cdkVersion: '2.109.0',
  defaultReleaseBranch: 'main',
  jsiiVersion: '~5.0.0',
  name: 'bedrock-agents-cdk',
  projenrcTs: true,
  repositoryUrl: 'https://github.com/maxtybar/bedrock-agents-cdk.git',
  githubOptions: {
    mergify: false,
  },

  deps: [
    'aws-cdk-lib@2.109.0',
  ], /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  devDeps: [
    'aws-cdk',
    'ts-node',
    'aws-cdk-lib@2.109.0',
  ], /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
  publishToPypi: {
    distName: 'bedrock-agent',
    module: 'bedrock_agent',
  },
  publishToGo: {
    moduleName: 'github.com/maxtybar/bedrock-agents-cdk',
  },
});

const common_exclude = ['cdk.out', 'cdk.context.json', '*.DS_Store', 'src/integ.default.ts', 'src/integ.default.agent.ts'];
project.npmignore!.exclude(...common_exclude);
project.gitignore.exclude(...common_exclude);

project.synth();