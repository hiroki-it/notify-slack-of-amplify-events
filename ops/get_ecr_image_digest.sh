#!/bin/bash

set -xeuo pipefail

ECR_IMAGE_DIGEST=$(aws ecr describe-images \
    --repository-name ${APP_ENV}-notify-slack-of-amplify-events-repository \
    --image-ids imageTag=${CIRCLE_SHA1} \
    --output text \
    --query 'imageDetails[0].imageDigest')
