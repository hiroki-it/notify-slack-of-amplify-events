#!/bin/bash

set -xeuo pipefail

case "$APP_ENV" in
    "stg")
        AWS_ACCOUNT_ID="$STG_AWS_ACCOUNT_ID"
    ;;
    "prd")
        AWS_ACCOUNT_ID="$PRD_AWS_ACCOUNT_ID"
    ;;
    *)
        echo "The parameter ${APP_ENV} is invalid."
        exit 1
    ;;
esac

# 環境変数を定義します．
echo 'export AWS_ECR_ACCOUNT_URL="${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com"' >> $BASH_ENV

# 環境変数を出力します．
source $BASH_ENV
