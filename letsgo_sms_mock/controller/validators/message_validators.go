package validators

import (
	"errors"
	"strconv"

	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"
	"github.com/google/uuid"
)

func SaveMessageValidator(message entities.Message) error {
	if message.ID == uuid.Nil {
		return errors.New("ID should not be empty")
	}
	if message.UserID == uuid.Nil {
		return errors.New("UserID should not be empty")
	}
	if message.Content == "" {
		return errors.New("message content should not be empty")
	}

	// validate sender
	_, err := strconv.Atoi(message.Sender)
	if err != nil {
		return err
	}
	// validate recevier
	for _, v := range message.Receiver {
		_, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
	}

	return nil
}
