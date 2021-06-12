package amplify

import (
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/mock/amplify"
)

/**
 * 事前処理を実行し，事後処理を返却します．
 */
func setup() (*m_amplify.MockedAmplifyAPI, func()) {

	// AmplifyAPIのスタブを作成する．
	mockedAPI := &m_amplify.MockedAmplifyAPI{}

	return mockedAPI, func() {
	}
}
