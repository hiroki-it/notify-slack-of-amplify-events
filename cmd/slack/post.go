package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/**
 * メッセージを送信します．
 */
func postMessage(message Message) (bool, error) {

	// マッピングを元に，構造体をJSONに変換する．
	json, err := json.Marshal(message)

	if err != nil {
		return false, err
	}

	// リクエストメッセージを定義する．
	request, err := http.NewRequest(
		"POST",
		os.Getenv("SLACK_API_URL"),
		bytes.NewBuffer(json),
	)

	if err != nil {
		return false, err
	}

	// ヘッダーを定義する．
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SLACK_API_TOKEN")))

	client := &http.Client{}

	// HTTPリクエストを送信する．
	response, err := client.Do(request)

	// deferで宣言しておき，HTTP通信を必ず終了できるようにする．
	defer response.Body.Close()

	if err != nil || response.StatusCode != 200 {
		return false, err
	}

	fmt.Printf("Success: %#v\n", response)

	return true, err
}
