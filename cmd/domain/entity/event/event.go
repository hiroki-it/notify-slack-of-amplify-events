package event

/**
 * Eventを構成します．
 */
type Event struct {
	eventDetail *EventDetail
}

/**
 * コンストラクタ
 * Eventを作成します．
 */
func NewEvent(eventDetail *EventDetail) *Event {

	return &Event{
		eventDetail: eventDetail,
	}
}

/**
 * EventDetailを返却します．
 */
func (event *Event) GetEventDetail() *EventDetail {
	return event.eventDetail
}
