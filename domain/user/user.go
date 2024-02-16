package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID       ID
	Email    Email
	Username Username
}

func New(email Email, un Username) User {
	return User{
		ID:       ID(uuid.New().String()),
		Email:    email,
		Username: un,
	}
}

type ID string

func NewID(i string) (id ID, err error) {
	if _, err := uuid.Parse(string(i)); err != nil {
		return "", err
	}
	return ID(i), nil
}

func (id ID) String() string {
	return string(id)
}

type Email string

func NewEmail(i string) (email Email, err error) {
	if email == "" {
		return "", errors.New("email is required")
	}

	return Email(i), nil
}

func (email Email) String() string {
	return string(email)
}

type Username string

func NewUsername(i string) (un Username, err error) {
	if un == "" {
		return "", errors.New("Username is required")
	}

	return Username(i), nil
}

func (un Username) String() string {
	return string(un)
}
