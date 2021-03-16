package amplify

import(
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * AppOption構造体
 */
type AppOption struct {
    AppId string `json:"appId"`
}

func GetAmplifyApp() amplifyApp {

    session, err := session.Must(
        session.NewSession(
            aws.NewConfig().WithRegion("ap-northeast-1")
        )
    )
    
    if err != nil {
        log.Fatalf("Failed: %#v\n", err)
    }
    
    amplify := amplify.New(session)
    
    appOption
    
    response = amplify.GetApp()
}
