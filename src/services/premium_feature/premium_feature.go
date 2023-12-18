package premiumfeature

import (
	"context"
	"DatingApp/src/filter"
	"DatingApp/src/models"
	premiumfeature "DatingApp/src/repositories/premium_feature"
	"time"
)

type Interface interface {
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, input models.Query[models.PremiumFeatureInput], id int) error
	Create(ctx context.Context, input models.Query[models.PremiumFeatureInput]) error
	Get(ctx context.Context, paging filter.Paging[filter.PremiumFeatureFilter]) ([]models.PremiumFeature, int, error)
}

type premiumFeatureService struct {
	premiumFeatureRepository premiumfeature.Interface
}

type Param struct {
	PremiumFeatureRepository premiumfeature.Interface
}

func Init(param Param) Interface {
	return &premiumFeatureService{
		premiumFeatureRepository: param.PremiumFeatureRepository,
	}
}

var Now = time.Now

func (s *premiumFeatureService) Delete(ctx context.Context, id int) error {
	input := models.Query[models.PremiumFeatureInput]{
		Model: models.PremiumFeatureInput{
			Status:    -1,
			DeletedAt: Now(),
			DeletedBy: ctx.Value(models.UserKey).(models.User).Id,
		},
	}

	return s.premiumFeatureRepository.Update(ctx, input, id)
}

func (s *premiumFeatureService) Update(ctx context.Context, input models.Query[models.PremiumFeatureInput], id int) error {
	input.Model.UpdatedAt = Now()
	input.Model.UpdatedBy = ctx.Value(string(models.UserKey)).(models.User).Id

	return s.premiumFeatureRepository.Update(ctx, input, id)
}

func (s *premiumFeatureService) Create(ctx context.Context, input models.Query[models.PremiumFeatureInput]) error {
	input.Model.CreatedAt = Now()
	input.Model.CreatedBy = ctx.Value(models.UserKey).(models.User).Id

	return s.premiumFeatureRepository.Create(ctx, input)
}

func (s *premiumFeatureService) Get(ctx context.Context, paging filter.Paging[filter.PremiumFeatureFilter]) ([]models.PremiumFeature, int, error) {
	paging.IsActive = true
	return s.premiumFeatureRepository.Get(ctx, paging)
}
