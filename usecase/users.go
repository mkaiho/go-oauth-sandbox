package usecase

import "github.com/mkaiho/go-oauth-sandbox/entity/users"

type UserUsecase interface {
	Find(id string) users.User
}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

type userUsecase struct{}

func (u *userUsecase) Find(id string) users.User {
	return users.NewUser(id)
}
