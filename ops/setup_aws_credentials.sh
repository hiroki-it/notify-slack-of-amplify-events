#!/bin/bash

set -xeuo pipefail

case "$APP_ENV" in
    "stg")
        AWS_ACCOUNT_ID="$STG_AWS_ACCOUNT_ID"
        AWS_ACCESS_KEY_ID=$STG_AWS_ACCESS_KEY_ID
        AWS_SECRET_ACCESS_KEY=$STG_AWS_SECRET_ACCESS_KEY
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

# defaultプロファイルにクレデンシャル情報を設定する．
aws configure << EOF
$(echo $AWS_ACCESS_KEY_ID)
$(echo $AWS_SECRET_ACCESS_KEY)
$(echo $AWS_DEFAULT_REGION)
json
EOF

# 正しく設定されたかを確認する．
aws configure list
