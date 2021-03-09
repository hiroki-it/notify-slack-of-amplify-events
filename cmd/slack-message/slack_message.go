package slack-message

import(
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// 構造体を定義し，JSONにマッピング
type SlackMessage struct {
    Token       string   `json:"token"`
    Channel     string   `json:"channel"`
    Text        string   `json:"text"`
    Username    string   `json:"username"`
//    Attachments []string `json:"attachments"`
}

func buildMessage() {
    // URL
    url := "https://xxxx.slack.com"

    // ボディを定義する．
    slackMessage := SlackMessage {
        Token: "<トークン文字列>",
        Channel: "<チャンネル名，もしくは@ユーザ名>",
        Text: "<メッセージ>",
        Username: "<as_userオプションがfalseの場合にBot名>",
//         Attachments: [{
//           // 任意のオプション     
//           // 参考：
//           // https://api.slack.com/messaging/composing/layouts#attachments
//         }]
    }
}

func postMessage(SlackMessage slackMessage) {
    // マッピングを元に，構造体をJSONに変換する．
    json, err := json.Marshal(slackMessage)

    if err != nil {
        log.Fatalf("ERROR: %#v\n", err)
    }

    // リクエストメッセージを定義する．
    request, err := http.NewRequest(
        "POST",
        url,
        bytes.NewBuffer(json),
    )

    if err != nil {
        log.Fatalf("ERROR: %#v\n", err)
    }

    // ヘッダーを定義する．
    request.Header.Set("Content-Type", "application/json")

    client := &http.Client {}

    // HTTPリクエストを送信する．
    response, err := client.Do(request)
    
    // deferで宣言しておき，HTTP通信を必ず終了できるようにする．
    defer response.Body.Close()
    
    if err != nil {
        log.Fatalf("ERROR: %#v\n", err)
    }
    
    if response.StatusCode != 200 {
        log.Fatalf("ERROR: %#v\n", response)
    }
    
    fmt.Printf("INFO: %#v\n", response)
}
