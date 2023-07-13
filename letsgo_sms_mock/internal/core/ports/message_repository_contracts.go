package ports

import "github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"

type MessageRepositoryContracts interface {
	SaveMessageDB(message entities.Message) error
	GiveAllMessages() ([]entities.Message , error)
}