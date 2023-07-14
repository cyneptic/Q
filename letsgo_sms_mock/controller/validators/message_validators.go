package validators

import (
	"errors"
	"strconv"

	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"
)

func SaveMessageValidator(message entities.DTOMessage) error {
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
