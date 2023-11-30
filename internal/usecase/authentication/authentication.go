package authenticationUsecase

import (
	"encoding/json"
	"net/http"

	"gitlab.com/sunrise-dev/utest/utest-auth-app/internal/repositories"
)

// AuthenticationUseCase -.
type AuthenticationUseCase struct {
	repo repositories.UserRepo
}

// New -.
func New(r repositories.UserRepo) *AuthenticationUseCase {
	return &AuthenticationUseCase{
		repo: r,
	}
}

// GetAll - getting translate history from store.
func (uc *AuthenticationUseCase) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, _ := uc.repo.GetAll()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(users)
	}
}
