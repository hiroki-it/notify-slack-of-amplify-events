#!/bin/bash

set -xeuo pipefail

case $APP_ENV in
    "stg")
        AWS_ACCOUNT_ID="$STG_AWS_ACCOUNT_ID"
        AWS_ACCESS_KEY_ID="$STG_AWS_ACCESS_KEY_ID"
        AWS_SECRET_ACCESS_KEY="$STG_AWS_SECRET_ACCESS_KEY"
    ;;
    "prd")
        AWS_ACCOUNT_ID="$PRD_AWS_ACCOUNT_ID"
        AWS_ACCESS_KEY_ID="$PRD_AWS_ACCESS_KEY_ID"
        AWS_SECRET_ACCESS_KEY="$PRD_AWS_SECRET_ACCESS_KEY"
    ;;
    *)
        echo "The parameter ${ENV} is invalid."
        exit 1
    ;;
esac

aws configure --profile "${APP_ENV}-notify-slack-of-amplify-events" << EOF
echo AWS_ACCOUNT_ID
echo AWS_ACCESS_KEY_ID
echo AWS_SECRET_ACCESS_KEY
json
EOF

cat << EOT > "envs.sh"
export APP_ENV="$APP_ENV"
export AWS_ECR_ACCOUNT_URL="${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
EOT
