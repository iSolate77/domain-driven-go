package service

import (
	"context"

	"github.com/iSolate/quick/domain/user"
)

type Service struct {
}

type Req struct {
	Username string
	Email    string
}

type Res struct {
	ID string
}

func (s *Service) CreateUser(ctx context.Context, req Req) (res Res, err error) {
	userName, err := user.NewUsername(req.Username)
	if err != nil {
		return res, err
	}
	email, err := user.NewEmail(req.Email)
	if err != nil {
		return res, err
	}

    u := user.New(email, userName)

    // add to db

	return res, nil
}
