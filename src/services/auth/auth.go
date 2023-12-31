package auth

import (
	"DatingApp/src/filter"
	"DatingApp/src/models"
	"DatingApp/src/repositories/auth"
	"DatingApp/src/repositories/user"
	"context"
	"errors"
)

type Interface interface {
	Register(ctx context.Context, input models.Query[models.UserInput]) error
	Login(ctx context.Context, input models.Login) ([]models.User, string, error)
}

type authService struct {
	authRepository auth.Interface
	userRepository user.Interface
}

type Param struct {
	AuthRepository auth.Interface
	UserRepository user.Interface
}

func Init(param Param) *authService {
	return &authService{userRepository: param.UserRepository, authRepository: param.AuthRepository}
}

func (s *authService) Register(ctx context.Context, input models.Query[models.UserInput]) error {
	_, count, err := s.userRepository.Get(ctx, filter.Paging[filter.UserFilter]{
		Page: 1,
		Take: 1,
		Filter: filter.UserFilter{
			UserName: input.Model.UserName,
		},
	})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already taken by another user")
	}

	password, err := s.authRepository.HashPassword([]byte(input.Model.Password))
	if err != nil {
		return err
	}
	input.Model.Password = password

	err = s.userRepository.Create(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) Login(ctx context.Context, input models.Login) ([]models.User, string, error) {

	users, _, err := s.userRepository.Get(ctx, filter.Paging[filter.UserFilter]{
		Page: 1,
		Take: 1,
		Filter: filter.UserFilter{
			UserName: input.UserName,
		},
	})
	if err != nil {
		return []models.User{}, "", err
	}
	if len(users) == 0 {
		return []models.User{}, "", errors.New("login failed")
	}

	err = s.authRepository.ComparePassword([]byte(users[0].Password), []byte(input.Password))
	if err != nil {
		return []models.User{}, "", errors.New("login failed")
	}

	token, err := s.authRepository.GenerateToken(int(users[0].Id), users[0].UserName)
	if err != nil {
		return []models.User{}, "", err
	}

	return users, token, nil
}
