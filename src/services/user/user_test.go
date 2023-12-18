package user_test

import (
	"DatingApp/src/filter"
	"DatingApp/src/models"
	mock_user "DatingApp/src/repositories/mock/user"
	user "DatingApp/src/services/user"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_userService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
		Id int
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	user.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		user.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "delete user error",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.Query[models.UserInput]{
					Model: models.UserInput{
						DeletedBy: context.Value(models.UserKey).(models.User).Id,
						DeletedAt: mockTime,
						Status:    -1,
					},
				}, 1).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "delete user success",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.Query[models.UserInput]{
					Model: models.UserInput{
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
				t.Errorf("user.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.Background()

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
		Paging filter.Paging[filter.UserFilter]
	}

	restoreAll := func() {
		user.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name      string
		args      args
		mockfunc  func(a args, mock mockfields)
		want      []models.User
		wantCount int
		wantErr   bool
	}{
		{
			name: "get user error",
			args: args{
				filter.Paging[filter.UserFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{IsActive: true}).Return([]models.User{}, 0, assert.AnError)
			},
			want:      []models.User{},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name: "get user success",
			args: args{
				filter.Paging[filter.UserFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{IsActive: true}).Return([]models.User{
					{},
					{},
				}, 2, nil)
			},
			want: []models.User{
				{},
				{},
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			users, count, err := service.Get(context, tt.args.Paging)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, users)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}

func Test_userService_UpdatePremiumFeatureId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
		Input models.Subscribe
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	user.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		user.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "update user error",
			args: args{
				Input: models.Subscribe{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.Query[models.UserInput]{
					Model: models.UserInput{
						UpdatedBy: context.Value(models.UserKey).(models.User).Id,
						UpdatedAt: mockTime,
					},
				}, 1).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "update user success",
			args: args{
				Input: models.Subscribe{
					PremiumFeatureId: 1,
				},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Update(context, models.Query[models.UserInput]{
					Model: models.UserInput{
						PremiumFeatureId: 1,
						UpdatedBy:        context.Value(models.UserKey).(models.User).Id,
						UpdatedAt:        mockTime,
					},
				}, 1).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			err := service.UpdatePremiumFeatureId(context, tt.args.Input)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.UpdatePremiumFeatureId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userService_GetRecomendedUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	userRepo := mock_user.NewMockInterface(ctrl)
	type mockfields struct {
		user *mock_user.MockInterface
	}
	mocks := mockfields{
		user: userRepo,
	}
	params := user.Param{
		UserRepository: userRepo,
	}
	service := user.Init(params)
	type args struct {
	}

	restoreAll := func() {
		user.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		want     models.RecomendationUser
		wantErr  bool
	}{
		{
			name: "get user recomendation error",
			args: args{},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().GetRecomendedUser(context, int(context.Value(models.UserKey).(models.User).Id)).Return(models.RecomendationUser{}, assert.AnError)
			},
			want:    models.RecomendationUser{},
			wantErr: true,
		},
		{
			name: "get user success",
			args: args{},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().GetRecomendedUser(context, int(context.Value(models.UserKey).(models.User).Id)).Return(models.RecomendationUser{}, nil)
			},
			want:    models.RecomendationUser{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			users, err := service.GetRecomendedUser(context)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.GetRecomendedUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, users)
		})
	}
}
