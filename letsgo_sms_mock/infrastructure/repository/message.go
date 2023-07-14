package repository

import "github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"

func (r *PGRepository) SaveMessageDB(message entities.Message) error {
	err := r.DB.Create(&message).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PGRepository) GiveAllMessages() ([]entities.Message, error) {
	var messages []entities.Message
	err := r.DB.Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
