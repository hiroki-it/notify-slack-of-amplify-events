package boundaries

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/requests"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/responses"
)

type DetailInputBoundary interface {
	NotifyEventDetail(*requests.DetailRequest) (*responses.GetDetailResponse, error)
}
