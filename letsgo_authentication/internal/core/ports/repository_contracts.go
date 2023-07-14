package ports

import (
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/google/uuid"
)

// اینترفیس دیتابیس هستش

type UserRepositoryContracts interface {
	DisableUser(target uuid.UUID, toggle bool) error
	AddUser(user entities.User) error
	Login(email string) (*entities.User, error)
	IsAdminAccount(id uuid.UUID, role string) (bool, error)
	Verify(number string, id uuid.UUID) (bool, error)
}

type InMemoryRespositoryContracts interface {
	AddToken(token string) error
	RevokeToken(token string) error
	TokenReceiver(token string) (string, error)
}
