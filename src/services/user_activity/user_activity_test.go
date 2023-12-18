package useractivity_test

import (
	"DatingApp/src/filter"
	"DatingApp/src/formatter"
	"DatingApp/src/models"
	mock_premium_feature "DatingApp/src/repositories/mock/premium_feature"
	mock_user "DatingApp/src/repositories/mock/user"
	mock_user_activity "DatingApp/src/repositories/mock/user_activity"
	useractivity "DatingApp/src/services/user_activity"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_userActivityService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	userActivityRepo := mock_user_activity.NewMockInterface(ctrl)
	user := mock_user.NewMockInterface(ctrl)
	premiumFeature := mock_premium_feature.NewMockInterface(ctrl)
	type mockfields struct {
		userActivity   *mock_user_activity.MockInterface
		user           *mock_user.MockInterface
		premiumFeature *mock_premium_feature.MockInterface
	}
	mocks := mockfields{
		userActivity:   userActivityRepo,
		user:           user,
		premiumFeature: premiumFeature,
	}
	params := useractivity.Param{
		UserActivityRepository:   userActivityRepo,
		UserRepository:           mocks.user,
		PremiumFeatureRepository: premiumFeature,
	}
	service := useractivity.Init(params)
	type args struct {
		Input models.Query[models.UserActivityInput]
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	useractivity.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		useractivity.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "get user error",
			args: args{
				Input: models.Query[models.UserActivityInput]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{
					Filter: filter.UserFilter{
						Id: int(context.Value(models.UserKey).(models.User).Id),
					},
				}).Return([]models.User{}, 0, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "get no user",
			args: args{
				Input: models.Query[models.UserActivityInput]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{
					Filter: filter.UserFilter{
						Id: int(context.Value(models.UserKey).(models.User).Id),
					},
				}).Return([]models.User{}, 0, nil)
			},
			wantErr: true,
		},
		{
			name: "get premium feature error",
			args: args{
				Input: models.Query[models.UserActivityInput]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{
					Filter: filter.UserFilter{
						Id: int(context.Value(models.UserKey).(models.User).Id),
					},
				}).Return([]models.User{{PremiumFeatureId: formatter.NullableDataType[int]{Data: 1, Valid: true}}}, 1, nil)
				mock.premiumFeature.EXPECT().Get(context, filter.Paging[filter.PremiumFeatureFilter]{
					Filter: filter.PremiumFeatureFilter{
						Flag: "no-swipe-quota-limit",
					},
				}).Return([]models.PremiumFeature{}, 0, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "get no premium feature",
			args: args{
				Input: models.Query[models.UserActivityInput]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{
					Filter: filter.UserFilter{
						Id: int(context.Value(models.UserKey).(models.User).Id),
					},
				}).Return([]models.User{{PremiumFeatureId: formatter.NullableDataType[int]{Data: 1, Valid: true}}}, 1, nil)
				mock.premiumFeature.EXPECT().Get(context, filter.Paging[filter.PremiumFeatureFilter]{
					Filter: filter.PremiumFeatureFilter{
						Flag: "no-swipe-quota-limit",
					},
				}).Return([]models.PremiumFeature{}, 0, nil)
			},
			wantErr: true,
		},
		{
			name: "create userActivity error",
			args: args{
				Input: models.Query[models.UserActivityInput]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{
					Filter: filter.UserFilter{
						Id: int(context.Value(models.UserKey).(models.User).Id),
					},
				}).Return([]models.User{{PremiumFeatureId: formatter.NullableDataType[int]{Data: 1, Valid: true}}}, 1, nil)
				mock.premiumFeature.EXPECT().Get(context, filter.Paging[filter.PremiumFeatureFilter]{
					Filter: filter.PremiumFeatureFilter{
						Flag: "no-swipe-quota-limit",
					},
				}).Return([]models.PremiumFeature{{Id: 1}}, 1, nil)
				mock.userActivity.EXPECT().Create(context, models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{
						CreatedBy: context.Value(models.UserKey).(models.User).Id,
						CreatedAt: mockTime,
					},
				}).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "create userActivity success",
			args: args{
				models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{},
				},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.user.EXPECT().Get(context, filter.Paging[filter.UserFilter]{
					Filter: filter.UserFilter{
						Id: int(context.Value(models.UserKey).(models.User).Id),
					},
				}).Return([]models.User{{PremiumFeatureId: formatter.NullableDataType[int]{Data: 1, Valid: true}}}, 1, nil)
				mock.premiumFeature.EXPECT().Get(context, filter.Paging[filter.PremiumFeatureFilter]{
					Filter: filter.PremiumFeatureFilter{
						Flag: "no-swipe-quota-limit",
					},
				}).Return([]models.PremiumFeature{{Id: 1}}, 1, nil)
				mock.userActivity.EXPECT().Create(context, models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{
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
				t.Errorf("userActivity.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userActivityService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	userActivityRepo := mock_user_activity.NewMockInterface(ctrl)
	type mockfields struct {
		userActivity *mock_user_activity.MockInterface
	}
	mocks := mockfields{
		userActivity: userActivityRepo,
	}
	params := useractivity.Param{
		UserActivityRepository: userActivityRepo,
	}
	service := useractivity.Init(params)
	type args struct {
		Input models.Query[models.UserActivityInput]
		Id    int
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	useractivity.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		useractivity.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "update userActivity error",
			args: args{
				Input: models.Query[models.UserActivityInput]{},
				Id:    1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.userActivity.EXPECT().Update(context, models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{
						UpdatedBy: context.Value(models.UserKey).(models.User).Id,
						UpdatedAt: mockTime,
					},
				}, 1).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "update userActivity success",
			args: args{
				Input: models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{},
				},
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.userActivity.EXPECT().Update(context, models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{
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
				t.Errorf("userActivity.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userActivityService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.WithValue(context.Background(), models.UserKey, models.User{Id: 1})

	userActivityRepo := mock_user_activity.NewMockInterface(ctrl)
	type mockfields struct {
		userActivity *mock_user_activity.MockInterface
	}
	mocks := mockfields{
		userActivity: userActivityRepo,
	}
	params := useractivity.Param{
		UserActivityRepository: userActivityRepo,
	}
	service := useractivity.Init(params)
	type args struct {
		Id int
	}

	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.Local)
	useractivity.Now = func() time.Time {
		return mockTime
	}

	restoreAll := func() {
		useractivity.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name     string
		args     args
		mockfunc func(a args, mock mockfields)
		wantErr  bool
	}{
		{
			name: "delete userActivity error",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.userActivity.EXPECT().Update(context, models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{
						DeletedBy: context.Value(models.UserKey).(models.User).Id,
						DeletedAt: mockTime,
						Status:    -1,
					},
				}, 1).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "delete userActivity success",
			args: args{
				Id: 1,
			},
			mockfunc: func(a args, mock mockfields) {
				mock.userActivity.EXPECT().Update(context, models.Query[models.UserActivityInput]{
					Model: models.UserActivityInput{
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
				t.Errorf("userActivity.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_userActivityService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := context.Background()

	userActivityRepo := mock_user_activity.NewMockInterface(ctrl)
	type mockfields struct {
		userActivity *mock_user_activity.MockInterface
	}
	mocks := mockfields{
		userActivity: userActivityRepo,
	}
	params := useractivity.Param{
		UserActivityRepository: userActivityRepo,
	}
	service := useractivity.Init(params)
	type args struct {
		Paging filter.Paging[filter.UserActivityFilter]
	}

	restoreAll := func() {
		useractivity.Now = time.Now
	}
	defer restoreAll()

	tests := []struct {
		name      string
		args      args
		mockfunc  func(a args, mock mockfields)
		want      []models.UserActivity
		wantCount int
		wantErr   bool
	}{
		{
			name: "get userActivity error",
			args: args{
				filter.Paging[filter.UserActivityFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.userActivity.EXPECT().Get(context, filter.Paging[filter.UserActivityFilter]{IsActive: true}).Return([]models.UserActivity{}, 0, assert.AnError)
			},
			want:      []models.UserActivity{},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name: "get userActivity success",
			args: args{
				filter.Paging[filter.UserActivityFilter]{},
			},
			mockfunc: func(a args, mock mockfields) {
				mock.userActivity.EXPECT().Get(context, filter.Paging[filter.UserActivityFilter]{IsActive: true}).Return([]models.UserActivity{
					{},
					{},
				}, 2, nil)
			},
			want: []models.UserActivity{
				{},
				{},
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc(tt.args, mocks)

			userActivitys, count, err := service.Get(context, tt.args.Paging)
			if (err != nil) != tt.wantErr {
				t.Errorf("userActivity.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, userActivitys)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}
