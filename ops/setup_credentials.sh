#!/bin/bash

set -xeuo pipefail

aws configure --profile $APP_ENV << EOF
$(echo "$AWS_ACCESS_KEY_ID")
$(echo "$AWS_SECRET_ACCESS_KEY")
$(echo "$AWS_DEFAULT_REGION")
json
EOF
