package integration

import (
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
)

/**
 * 事前処理を実行し，事後処理を返却します．
 */
func setup(t *testing.T) ([]byte, func()) {

	// 前処理
	detail, err := file.ReadDataFile("./testdata/event.json")

	if err != nil {
		t.Fatal(err.Error())
	}

	return detail, func() {
		// 後処理
	}
}
