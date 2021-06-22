package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

// NewAmplifyAPI コンストラクタ
func NewAmplifyAPI(region string) (amplifyiface.AmplifyAPI, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		return nil, err
	}

	return aws_amplify.New(sess), nil
}
