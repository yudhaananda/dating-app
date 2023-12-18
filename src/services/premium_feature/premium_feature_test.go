package premiumfeature_test

import (
	"DatingApp/src/filter"
	"DatingApp/src/models"
	mock_premium_feature "DatingApp/src/repositories/mock/premium_feature"
	premiumfeature "DatingApp/src/services/premium_feature"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_premiumFeatureService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	premiumFeatureRepo := mock_premium_feature.NewMockInterface(ctrl)
	type mockfields struct {
		premiumFeature *mock_premium_feature.MockInterface
	}
	mocks := mockfields{
		premiumFeature: premiumFeatureRepo,
	}
	params := premiumfeature.Param{
		PremiumFeatureRepository: premiumFeatureRepo,
	}
	service := premiumfeature.Init(params)
	type args struct {
		Input models.Query[models.PremiumFeatureInput]
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	premiumfeature.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		premiumfeature.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "create premiumFeature error",
			args: args{
				Input: models.Query[models.PremiumFeatureInput]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Create(context, models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{
						CreatedBy: context.Value(models.UserKey).(models.User).Id,
						CreatedAt: mockTime,
					},
				}).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "create premiumFeature success",
			args: args{
				models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{},
				},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Create(context, models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{
						CreatedBy: context.Value(models.UserKey).(models.User).Id,
						CreatedAt: mockTime,
					},
				}).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.Create(context, tt.args.Input)
			if (err != nil) != tt.wantErr {
				t.Errorf("premiumFeature.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_premiumFeatureService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	premiumFeatureRepo := mock_premium_feature.NewMockInterface(ctrl)
	type mockfields struct {
		premiumFeature *mock_premium_feature.MockInterface
	}
	mocks := mockfields{
		premiumFeature: premiumFeatureRepo,
	}
	params := premiumfeature.Param{
		PremiumFeatureRepository: premiumFeatureRepo,
	}
	service := premiumfeature.Init(params)
	type args struct {
		Input models.Query[models.PremiumFeatureInput]
		Id    int
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	premiumfeature.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		premiumfeature.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "update premiumFeature error",
			args: args{
				Input: models.Query[models.PremiumFeatureInput]{},
				Id:    1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Update(context, models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{
						UpdatedBy: context.Value(models.UserKey).(models.User).Id,
						UpdatedAt: mockTime,
					},
				}, 1).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "update premiumFeature success",
			args: args{
				Input: models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{},
				},
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Update(context, models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{
						UpdatedBy: context.Value(models.UserKey).(models.User).Id,
						UpdatedAt: mockTime,
					},
				}, 1).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.Update(context, tt.args.Input, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("premiumFeature.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_premiumFeatureService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	premiumFeatureRepo := mock_premium_feature.NewMockInterface(ctrl)
	type mockfields struct {
		premiumFeature *mock_premium_feature.MockInterface
	}
	mocks := mockfields{
		premiumFeature: premiumFeatureRepo,
	}
	params := premiumfeature.Param{
		PremiumFeatureRepository: premiumFeatureRepo,
	}
	service := premiumfeature.Init(params)
	type args struct {
		Id int
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	premiumfeature.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		premiumfeature.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "delete premiumFeature error",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Update(context, models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{
						DeletedBy: context.Value(models.UserKey).(models.User).Id,
						DeletedAt: mockTime,
						Status:    -1,
					},
				}, 1).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "delete premiumFeature success",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Update(context, models.Query[models.PremiumFeatureInput]{
					Model: models.PremiumFeatureInput{
						DeletedBy: context.Value(models.UserKey).(models.User).Id,
						DeletedAt: mockTime,
						Status:    -1,
					},
				}, 1).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.Delete(context, tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("premiumFeature.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_premiumFeatureService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.Background()

	premiumFeatureRepo := mock_premium_feature.NewMockInterface(ctrl)
	type mockfields struct {
		premiumFeature *mock_premium_feature.MockInterface
	}
	mocks := mockfields{
		premiumFeature: premiumFeatureRepo,
	}
	params := premiumfeature.Param{
		PremiumFeatureRepository: premiumFeatureRepo,
	}
	service := premiumfeature.Init(params)
	type args struct {
		Paging filter.Paging[filter.PremiumFeatureFilter]
	}

	restoreAll := func() {
		premiumfeature.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name      string
		args      args
		mockfunc  func(a args, mock mockfields)
		want      []models.PremiumFeature
		wantCount int
		wantErr   bool
	}{
		{
			name: "get premiumFeature error",
			args: args{
				filter.Paging[filter.PremiumFeatureFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Get(context, filter.Paging[filter.PremiumFeatureFilter]{IsActive: true}).Return([]models.PremiumFeature{}, 0, assert.AnError)
			},
			want:      []models.PremiumFeature{},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name: "get premiumFeature success",
			args: args{
				filter.Paging[filter.PremiumFeatureFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.premiumFeature.EXPECT().Get(context, filter.Paging[filter.PremiumFeatureFilter]{IsActive: true}).Return([]models.PremiumFeature{
					{},
					{},
				}, 2, nil)
			},
			want: []models.PremiumFeature{
				{},
				{},
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			premiumFeatures, count, err := service.Get(context, tt.args.Paging)
			if (err != nil) != tt.wantErr {
				t.Errorf("premiumFeature.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, premiumFeatures)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}
