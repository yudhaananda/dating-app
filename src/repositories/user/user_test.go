package user

import (
	"DatingApp/src/filter"
	"DatingApp/src/formatter"
	"DatingApp/src/models"
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO user () VALUES ()")

	type args struct {
		ctx    context.Context
		models models.Query[models.UserInput]
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantErr     bool
	}{
		{
			name: "sql begin failed",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin().WillReturnError(err)
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql exec failed",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql no row affected",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(0))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql commit failed",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(1))
				sqlMock.ExpectCommit().WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql commit success",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(1))
				sqlMock.ExpectCommit()
				return sqlServer, err
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()
			init := Init(Param{
				Db:        sqlServer,
				TableName: "user",
			})
			err = init.Create(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE user SET  WHERE ")

	type args struct {
		ctx    context.Context
		models models.Query[models.UserInput]
		id     int
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantErr     bool
	}{
		{
			name: "sql begin failed",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin().WillReturnError(err)
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql exec failed",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql no row affected",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(0))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql commit failed",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(1))
				sqlMock.ExpectCommit().WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr: true,
		},
		{
			name: "sql commit success",
			args: args{
				ctx:    context.Background(),
				models: models.Query[models.UserInput]{},
				id:     1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectBegin()
				sqlMock.ExpectExec(query).WillReturnResult(driver.RowsAffected(1))
				sqlMock.ExpectCommit()
				return sqlServer, err
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()
			init := Init(Param{
				Db:        sqlServer,
				TableName: "user",
			})
			err = init.Update(tt.args.ctx, tt.args.models, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tempModels := models.Query[models.User]{}
	member := tempModels.BuildTableMember()
	query := regexp.QuoteMeta("SELECT " + member + " FROM user WHERE 1=1")
	queryCount := regexp.QuoteMeta("SELECT COUNT(*) FROM user")
	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.UTC)

	type args struct {
		ctx    context.Context
		models filter.Paging[filter.UserFilter]
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantUser    []models.User
		wantCount   int
		wantErr     bool
	}{
		{
			name: "sql count query failed",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.UserFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectQuery(queryCount).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantUser: []models.User{},
			wantErr:  true,
		},
		{
			name: "sql query failed",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.UserFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				sqlMock.ExpectQuery(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr:   true,
			wantUser:  []models.User{},
			wantCount: 1,
		},
		{
			name: "sql success",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.UserFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				row := sqlMock.NewRows([]string{"id", "user_name", "password", "image", "premium_feature_id", "status", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"})
				row.AddRow(1, "test", "test", formatter.NullableDataType[string]{Valid: true, Data: "test"}, formatter.NullableDataType[int]{Valid: false, Data: 0}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1)
				sqlMock.ExpectQuery(query).WillReturnRows(row)
				return sqlServer, err
			},
			wantUser: []models.User{
				{
					Id:               1,
					UserName:         "test",
					Password:         "test",
					Image:            formatter.NullableDataType[string]{Valid: true, Data: "test"},
					PremiumFeatureId: formatter.NullableDataType[int]{Valid: false, Data: 0},
					Status:           1,
					CreatedAt: formatter.NullableDataType[time.Time]{
						Data:  mockTime,
						Valid: true,
					},
					UpdatedAt: formatter.NullableDataType[time.Time]{
						Data:  mockTime,
						Valid: true,
					},
					DeletedAt: formatter.NullableDataType[time.Time]{
						Data:  mockTime,
						Valid: true,
					},
					CreatedBy: formatter.NullableDataType[int64]{
						Data:  1,
						Valid: true,
					},
					UpdatedBy: formatter.NullableDataType[int64]{
						Data:  1,
						Valid: true,
					},
					DeletedBy: formatter.NullableDataType[int64]{
						Data:  1,
						Valid: true,
					},
				},
			},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()
			init := Init(Param{
				Db:        sqlServer,
				TableName: "user",
			})
			users, count, err := init.Get(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantUser, users)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}

func TestGetRecomendedUser(t *testing.T) {
	tempModels := models.Query[models.RecomendationUser]{}
	member := tempModels.BuildTableMember()
	query := regexp.QuoteMeta(fmt.Sprintf(GetRecomendUser, member))
	var mockImage *string
	var mockId *int
	timeNowMin := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	timeNowMax := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 1e9, time.UTC)

	type args struct {
		ctx    context.Context
		userId int
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantUser    models.RecomendationUser
		wantErr     bool
	}{
		{
			name: "sql query failed",
			args: args{
				ctx:    context.Background(),
				userId: 1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectQuery(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr:  true,
			wantUser: models.RecomendationUser{},
		},
		{
			name: "sql success",
			args: args{
				ctx:    context.Background(),
				userId: 1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				row := sqlMock.NewRows([]string{"id", "user_name", "image", "premium_feature_id"})
				row.AddRow(2, "test", mockImage, mockId)
				sqlMock.ExpectQuery(query).WithArgs(1, timeNowMin, timeNowMax, 1, timeNowMin, timeNowMax, 1).WillReturnRows(row)
				return sqlServer, err
			},
			wantUser: models.RecomendationUser{
				Id:               2,
				UserName:         "test",
				PremiumFeatureId: mockId,
				Image:            mockImage,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()
			init := Init(Param{
				Db:        sqlServer,
				TableName: "user",
			})
			user, err := init.GetRecomendedUser(tt.args.ctx, tt.args.userId)
			fmt.Println("hehe", user)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.GetRecomendedUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantUser, user)
		})
	}
}
