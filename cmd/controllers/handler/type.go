package handler

/**
 * EventBrdigeのイベントを構成します．
 */
type Request struct {
	Records []struct {
		EventBridge struct {
			Event string `json:"event"`
		}
	}
}
