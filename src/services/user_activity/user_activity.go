package useractivity

import (
	"DatingApp/src/filter"
	"DatingApp/src/models"
	premiumfeature "DatingApp/src/repositories/premium_feature"
	"DatingApp/src/repositories/user"
	useractivity "DatingApp/src/repositories/user_activity"
	"context"
	"errors"
	"time"
)

const (
	noLimitSwipeQuotaFlag = "no-swipe-quota-limit"
)

type Interface interface {
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, input models.Query[models.UserActivityInput], id int) error
	Create(ctx context.Context, input models.Query[models.UserActivityInput]) error
	Get(ctx context.Context, paging filter.Paging[filter.UserActivityFilter]) ([]models.UserActivity, int, error)
}

type userActivityService struct {
	userActivityRepository   useractivity.Interface
	userRepository           user.Interface
	premiumFeatureRepository premiumfeature.Interface
}

type Param struct {
	UserActivityRepository   useractivity.Interface
	UserRepository           user.Interface
	PremiumFeatureRepository premiumfeature.Interface
}

func Init(param Param) Interface {
	return &userActivityService{
		userActivityRepository:   param.UserActivityRepository,
		userRepository:           param.UserRepository,
		premiumFeatureRepository: param.PremiumFeatureRepository,
	}
}

var Now = time.Now

func (s *userActivityService) Delete(ctx context.Context, id int) error {
	input := models.Query[models.UserActivityInput]{
		Model: models.UserActivityInput{
			Status:    -1,
			DeletedAt: Now(),
			DeletedBy: ctx.Value(models.UserKey).(models.User).Id,
		},
	}

	return s.userActivityRepository.Update(ctx, input, id)
}

func (s *userActivityService) Update(ctx context.Context, input models.Query[models.UserActivityInput], id int) error {
	input.Model.UpdatedAt = Now()
	input.Model.UpdatedBy = ctx.Value(string(models.UserKey)).(models.User).Id

	return s.userActivityRepository.Update(ctx, input, id)
}

func (s *userActivityService) Create(ctx context.Context, input models.Query[models.UserActivityInput]) error {
	userId := ctx.Value(models.UserKey).(models.User).Id
	users, _, err := s.userRepository.Get(ctx, filter.Paging[filter.UserFilter]{
		Filter: filter.UserFilter{
			Id: int(userId),
		},
	})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("user doesnt exists")
	}
	user := users[0]

	features, _, err := s.premiumFeatureRepository.Get(ctx, filter.Paging[filter.PremiumFeatureFilter]{
		Filter: filter.PremiumFeatureFilter{
			Flag: noLimitSwipeQuotaFlag,
		},
	})
	if err != nil {
		return err
	}
	if len(features) == 0 {
		return errors.New("feature doesnt exists")
	}
	feeature := features[0]

	if !user.PremiumFeatureId.Valid || user.PremiumFeatureId.Data != int(feeature.Id) {
		totalActivity, err := s.userActivityRepository.GetTotalTodayActivity(ctx, int(userId))
		if err != nil {
			return err
		}
		if totalActivity >= 10 {
			return errors.New("reached total of max activity today")
		}
	}

	input.Model.CreatedAt = Now()
	input.Model.CreatedBy = userId

	return s.userActivityRepository.Create(ctx, input)
}

func (s *userActivityService) Get(ctx context.Context, paging filter.Paging[filter.UserActivityFilter]) ([]models.UserActivity, int, error) {
	paging.IsActive = true
	return s.userActivityRepository.Get(ctx, paging)
}
