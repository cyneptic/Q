package ports

import "github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"

type MessageServiceContracts interface {
	SaveMessageDB(message entities.Message) error
	GiveAllMessages() ([]entities.Message, error)
}