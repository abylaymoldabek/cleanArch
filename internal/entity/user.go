// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

// User -.
type User struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Surname    *string   `json:"surname" db:"surname"`
	FatherName *string   `json:"fathername" db:"fathername"`
	Email      string    `json:"email" db:"email"`
	Iin        int       `json:"iin" db:"iin"`
	Password   string    `json:"password" db:"password"`
	Role       int8      `json:"role" db:"role"`
	Token      *string   `json:"token" db:"token"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  *string   `json:"updated_at" db:"updated_at"`
	DeletedAt  *string   `json:"deleted_at" db:"deleted_at"`
}
