package slack

import(
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

/**
 * Amplifyイベント構造体
 */
type Request struct {
    Records []struct { 
        Amplify struct {
            Event string `json:event`
        }
    }
}

/**
 * Event構造体
 */
type Event struct {
    Version    string   `json:"version"`
    Id         string   `json:"id"`
    DetailType string   `json:"detail-type"`
    Source     string   `json:"source"`
    Account    string   `json:"account"`
    Time       string   `json:"time"`
    Region     string   `json:"region"`
    Resources  []string `json:"resources"`
    Detail     struct {
        AppId      string `json:"appId"`
        BranchName string `json:"branchName"`
        JobId      string `json:"jobId"`
        JobStatus  string `json:"jobStatus"`
    } `json:"detail"`
}

/**
 * メッセージ構造体
 */
type Message struct {
    Channel string  `json:"channel"`
    Text    string  `json:"text"`
    Attachments []Attachment `json:"attachments"`
}

/**
 * Attachmentスライスの要素
 */
type Attachment struct {
    Color  string `json:"color"`
    Blocks  []Block `json:"blocks"`
}

/**
 * Blockスライスの要素
 */
type Block struct {
    Type string `json:"type"`
    Text struct {
        Type string `json:"type"`
        Text string `json:"text"`
    } `json:"text,omitempty"`
    Elements []Element `json:"elements,omitempty"`
}

/**
 * Elementスライスの要素
 */
type Element struct {
    Type string `json:"type"`
    Text string `json:"text"`
}

/**
 * ハンドラー関数
 */
func handler(request Request) {

    var event Event
    
    // EventBridgeから転送されたイベントをマッピングします．
    err := json.Unmarshal([]byte(request.Records[0].Amplify.Event), &event)
    
    if err != nil {
        log.Fatalf("Failed: %#v\n", err)
    }

    message := BuildMessage(event)
  
    return PostMessage(message)
}

/**
 * Slackに送信するメッセージを構成します．
 */
func BuildMessage(event Event) Message {
    
    // 通知の色を判定します．
    var color string
    if (event.jobStatus == "SUCCEED") {
        color = "good"
    } else {
        color = "danger"
    }
    
    // メッセージを構成します．
    return Message{
        Channel: os.Getenv("SLACK_CHANNEL_ID"),
        Text: "検証用dev環境",
        Attachments: []Attachment{
            Color: color,
            Blocks: []Block{
                Block{
                    Type: "section",
                    Element: {
                        Type: "mrkdwn",
                        Text: "*検証用dev環境*",
                    },
                },
                Block{
                    Type: "context",
                    Element: {
                        Type: "mrkdwn",
                        Text: Sprintf(
                                "*結果*: %s",
                                event.Detail.JobStatus,
                            ),
                    },
                },
                Block{
                    Type: "context",
                    Element: {
                        Type: "mrkdwn",
                        Text: Sprintf(
                                "*ブランチ名*: %s",
                                event.Detail.BranchName,
                            ),
                    },
                },
                Block{
                    Type: "context",
                    Element: {
                        Type: "mrkdwn",
                        Text: Sprintf(
                                ":github:*プルリクURL*: %s/compare/%s",
                                data.repository,
                                event.Detail.BranchName,
                            ),
                    },
                },                             
                Block{
                    Type: "context",
                    Element: {
                        Type: "mrkdwn",
                        Text: Sprintf(
                                "*検証URL*: https://%s.%s.amplifyapp.com", 
                                event.Detail.BranchName,
                                event.Detail.AppId,
                            ),
                    },
                },
                Block{
                    Type: "context",
                    Element: {
                        Type: "mrkdwn",
                        Text: Sprintf(
                                ":amplify: <https://%s.console.aws.amazon.com/amplify/home?region=%s#/%s/%s/%s|*Amplifyコンソール画面はこちら*>",
                                event.Region,
                                event.Region,
                                event.Detail.AppId,
                                event.Detail.BranchName,
                                event.Detail.JobId,
                            ),
                    },
                },
                Block{
                    Type: "divider",
                },
            },
        },
    }
}

/**
 * Slackにメッセージを送信します．
 */
func PostMessage(Message message) bool {

    // マッピングを元に，構造体をJSONに変換する．
    json, err := json.Marshal(message)

    if err != nil {
        log.Fatalf("Failed: %#v\n", err)
    }

    // リクエストメッセージを定義する．
    request, err := http.NewRequest(
        "POST",
        os.Getenv("SLACK_API_URL"),
        bytes.NewBuffer(json),
    )

    if err != nil {
        log.Fatalf("Failed: %#v\n", err)
    }

    // ヘッダーを定義する．
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Authorization", Sprintf("Bearer %s", os.Getenv("SLACK_API_TOKEN")))

    client := &http.Client {}

    // HTTPリクエストを送信する．
    response, err := client.Do(request)
    
    // deferで宣言しておき，HTTP通信を必ず終了できるようにする．
    defer response.Body.Close()
    
    if err != nil {
        log.Fatalf("Failed: %#v\n", err)
    }
    
    if response.StatusCode != 200 {
        log.Fatalf("Failed: %#v\n", err)
    }
    
    fmt.Printf("Success %#v\n", response)
}
