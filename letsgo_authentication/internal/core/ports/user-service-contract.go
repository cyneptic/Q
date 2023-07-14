package ports

import (
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type UserServiceContract interface {
	DisableUser(target uuid.UUID, toggle bool) error
	AddUser(newUser entities.User) error
	LoginService(user entities.User) (string, string, error)
	Logout(token string) error
	IsAdminAccount(id uuid.UUID, role string) (bool, error)
	Verify(number string, id uuid.UUID) (bool, error)
}

type InMemoryServiceContracts interface {
	AddToken(token string)
	RevokeToken(token string) *redis.StatusCmd
	TokenReceiver() (string, error)
}
