#!/bin/bash
STACK_NAME=sdn-streaming-data-analysis
REGION=us-east-1
CLI_PROFILE=sdn

EC2_INSTANCE_TYPE=t2.micro
KEY_NAME=id_rsa_sergio_2020-03-25

# Deploy the CloudFormation template
echo -e "\n\n=========== Deploying main.yml ==========="
aws cloudformation deploy --region $REGION \
    --profile $CLI_PROFILE \
    --stack-name $STACK_NAME \
    --template-file main.yml \
    --no-fail-on-empty-changeset \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameter-overrides \
    EC2InstanceType=$EC2_INSTANCE_TYPE \
    KeyName=$KEY_NAME

# If the deploy succeeded, show the DNS name of the created instance
if [ $? -eq 0 ]; then
  aws cloudformation list-exports \
    --profile $CLI_PROFILE \
    --region $REGION | cat  > exports
fi
