package service

import (
	"github.com/cyneptic/letsgo_smspanel_mockapi/infrastructure/repository"
	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"
	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/ports"
)

type MessageService struct {
	db ports.MessageRepositoryContracts
}

func NewMessageService() *MessageService {
	db := repository.NewGormDatabase()
	return &MessageService{
		db: db,
	}
}

func (r *MessageService) SaveMessageDB(message entities.Message) error  {
	err := r.db.SaveMessageDB(message)
	return err
}

func (r *MessageService) GiveAllMessages() ([]entities.Message , error) {
	messages , err := r.db.GiveAllMessages() 
	return messages , err
}