package repositories

import "github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/entities"

type FileRepository interface {
	GetFile(*entities.File) []byte
}
