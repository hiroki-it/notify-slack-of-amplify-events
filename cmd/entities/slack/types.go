package slack

/**
 * Slack通知を構成します．
 */
type SlackNotification struct {
	slackClient  *SlackClient
	slackMessage *SlackMessage
}

/**
 * Slackメッセージを構成します．
 */
type SlackMessage struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

/**/
type Attachment struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

/**/
type Block struct {
	Type     string    `json:"type"`
	Text     *Text     `json:"text,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

/**/
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**/
type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**
 * SlackClientインターフェースを構成します．
 */
type SlackClientInterface interface {
}

/**
 * SlackClientインターフェースの実装を構成します．
 */
type SlackClient struct {
	SlackClientInterface
}
