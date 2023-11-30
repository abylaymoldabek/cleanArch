package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/sunrise-dev/utest/utest-auth-app/internal/entity"
)

// TranslationRepo -.
type UserRepoSQL struct {
	*sqlx.DB
}

// New -.
func NewUserRepoSQL(db *sqlx.DB) *UserRepoSQL {
	return &UserRepoSQL{db}
}

// GetAll -.
func (r *UserRepoSQL) GetAll() ([]entity.User, error) {

	users := []entity.User{}
	err := r.DB.Select(&users, "select * from users")
	if err != nil {
		log.Println("DB Error: get all users")
	}
	log.Println(users)
	return users, nil
}
