// Package usecase implements application business logic. Each logic group in own file.
package repositories

import (
	"gitlab.com/sunrise-dev/utest/utest-auth-app/internal/entity"
)

type (
	// UserRepo -.
	UserRepo interface {
		GetAll() ([]entity.User, error)
	}
)
