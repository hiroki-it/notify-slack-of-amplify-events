package amplify

import (
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/mock/amplify"
)

/**
 * 前処理の結果と，後処理の関数を返却します．
 */
func setup() (*m_amplify.MockedAmplifyAPI, func()) {

	// AmplifyAPIのスタブを作成します．
	mockedAPI := &m_amplify.MockedAmplifyAPI{}

	return mockedAPI, func() {
	}
}
