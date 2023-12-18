package user

import (
	"DatingApp/src/filter"
	"DatingApp/src/models"
	user "DatingApp/src/repositories/user"
	"context"
	"time"
)

type Interface interface {
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, paging filter.Paging[filter.UserFilter]) ([]models.User, int, error)
	UpdatePremiumFeatureId(ctx context.Context, input models.Subscribe) error
	GetRecomendedUser(ctx context.Context) (models.RecomendationUser, error)
}

type userService struct {
	userRepository user.Interface
}

type Param struct {
	UserRepository user.Interface
}

func Init(param Param) Interface {
	return &userService{
		userRepository: param.UserRepository,
	}
}

var Now = time.Now

func (s *userService) Delete(ctx context.Context, id int) error {
	input := models.Query[models.UserInput]{
		Model: models.UserInput{
			Status:    -1,
			DeletedAt: Now(),
			DeletedBy: ctx.Value(models.UserKey).(models.User).Id,
		},
	}

	return s.userRepository.Update(ctx, input, id)
}

func (s *userService) Get(ctx context.Context, paging filter.Paging[filter.UserFilter]) ([]models.User, int, error) {
	paging.IsActive = true
	return s.userRepository.Get(ctx, paging)
}

func (s *userService) UpdatePremiumFeatureId(ctx context.Context, input models.Subscribe) error {
	userId := ctx.Value(string(models.UserKey)).(models.User).Id
	model := models.Query[models.UserInput]{
		Model: models.UserInput{
			PremiumFeatureId: input.PremiumFeatureId,
			UpdatedAt:        Now(),
			UpdatedBy:        userId,
		},
	}
	return s.userRepository.Update(ctx, model, int(userId))
}

func (s *userService) GetRecomendedUser(ctx context.Context) (models.RecomendationUser, error) {
	return s.userRepository.GetRecomendedUser(ctx, int(ctx.Value(string(models.UserKey)).(models.User).Id))
}
