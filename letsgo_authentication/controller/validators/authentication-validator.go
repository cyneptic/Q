package validators

import (
	"errors"

	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/google/uuid"
)

func LoginValidation(u entities.User) error {
	if u.Email == "" {
		err := errors.New("please enter a valid email")
		return err
	}
	if u.Password == "" {
		err := errors.New("please enter a valid password")
		return err
	}
	return nil
}

func LogoutValidation(h string) error {
	if h == "" {
		err := errors.New("please enter a valid token")
		return err
	}
	return nil
}

func RegisterValidation(u entities.User) error {
	if u.Name == "" {
		err := errors.New("please enter a valid name")
		return err
	}
	if u.DateOfBirth.IsZero() {
		err := errors.New("please enter a valid name")
		return err
	}
	if u.PhoneNumber == "" {
		err := errors.New("please enter a valid phone number")
		return err
	}
	if u.Email == "" {
		err := errors.New("please enter a valid email")
		return err
	}
	if u.Password == "" {
		err := errors.New("please enter a valid password")
		return err
	}
	return nil
}
func IsAdmin(id string, role string) error {
	if id == "" {
		err := errors.New("please enter a valid id")
		return err
	}
	if role == "" || (role != "super_admin" && role != "ticket_admin" && role != "sms_admin") {
		err := errors.New("please enter a valid role")
		return err
	}
	return nil
}
func VerifyValidation(number, id string) error {
	if number == "" || id == "" {
		err := errors.New("please enter a valid id and phone number")
		return err
	}
	return nil
}

func DisableUser(target string, toggle bool) error {
	_, err := uuid.Parse(target)
	if err != nil {
		return err
	}

	return nil
}
