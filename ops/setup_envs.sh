#!/bin/bash

set -xeuo pipefail

case $APP_ENV in
    "stg")
        AWS_ACCOUNT_ID=$STG_AWS_ACCOUNT_ID
        AWS_ACCESS_KEY_ID=$STG_AWS_ACCESS_KEY_ID
        AWS_SECRET_ACCESS_KEY=$STG_AWS_SECRET_ACCESS_KEY
    ;;
    "prd")
        AWS_ACCOUNT_ID=$PRD_AWS_ACCOUNT_ID
        AWS_ACCESS_KEY_ID=$PRD_AWS_ACCESS_KEY_ID
        AWS_SECRET_ACCESS_KEY=$PRD_AWS_SECRET_ACCESS_KEY
    ;;
    *)
        echo "The parameter ${ENV} is invalid."
        exit 1
    ;;
esac

# CircleCIの環境変数として，値を出力する．
echo "export APP_ENV=$APP_ENV" >> $BASH_ENV
echo "export AWS_ACCOUNT_ID" >> $BASH_ENV
echo "export AWS_ACCESS_KEY_ID" >> $BASH_ENV
echo "export AWS_SECRET_ACCESS_KEY" >> $BASH_ENV
echo "export AWS_ECR_ACCOUNT_URL=${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com" >> $BASH_ENV

# 値が正しく出力されたかを確認する．
printenv | sort -f
