package useractivity

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
	query := regexp.QuoteMeta("INSERT INTO user_activity () VALUES ()")

	type args struct {
		ctx    context.Context
		models models.Query[models.UserActivityInput]
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				TableName: "user_activity",
			})
			err = init.Create(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("user_activity.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE user_activity SET  WHERE ")

	type args struct {
		ctx    context.Context
		models models.Query[models.UserActivityInput]
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				models: models.Query[models.UserActivityInput]{},
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
				TableName: "user_activity",
			})
			err = init.Update(tt.args.ctx, tt.args.models, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user_activity.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tempModels := models.Query[models.UserActivity]{}
	member := tempModels.BuildTableMember()
	query := regexp.QuoteMeta("SELECT " + member + " FROM user_activity WHERE 1=1")
	queryCount := regexp.QuoteMeta("SELECT COUNT(*) FROM user_activity")
	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.UTC)

	type args struct {
		ctx    context.Context
		models filter.Paging[filter.UserActivityFilter]
	}
	tests := []struct {
		name             string
		args             args
		prepSqlMock      func() (*sql.DB, error)
		wantUserActivity []models.UserActivity
		wantCount        int
		wantErr          bool
	}{
		{
			name: "sql count query failed",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.UserActivityFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectQuery(queryCount).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantUserActivity: []models.UserActivity{},
			wantErr:          true,
		},
		{
			name: "sql query failed",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.UserActivityFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				sqlMock.ExpectQuery(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr:          true,
			wantUserActivity: []models.UserActivity{},
			wantCount:        1,
		},
		{
			name: "sql success",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.UserActivityFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				row := sqlMock.NewRows([]string{"id", "user_id", "passed_user_id", "liked_user_id", "status", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"})
				row.AddRow(1, 1, formatter.NullableDataType[int]{Valid: false, Data: 0}, formatter.NullableDataType[int]{Valid: false, Data: 0}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1)
				sqlMock.ExpectQuery(query).WillReturnRows(row)
				return sqlServer, err
			},
			wantUserActivity: []models.UserActivity{
				{
					Id:           1,
					UserId:       1,
					PassedUserId: formatter.NullableDataType[int]{Valid: false, Data: 0},
					LikedUserId:  formatter.NullableDataType[int]{Valid: false, Data: 0},

					Status: 1,
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
				TableName: "user_activity",
			})
			userActivitys, count, err := init.Get(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("user_activity.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantUserActivity, userActivitys)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}

func TestGetTotalTodayActivity(t *testing.T) {
	timeNowMin := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	timeNowMax := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 1e9, time.UTC)
	where := fmt.Sprintf(" WHERE user_id = %d AND created_at > '%s' AND created_at < '%s'", 1, timeNowMin, timeNowMax)
	query := regexp.QuoteMeta(GetTotalTodayActivity + where)

	type args struct {
		ctx    context.Context
		userId int
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		want        int
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
			wantErr: true,
			want:    0,
		},
		{
			name: "sql success",
			args: args{
				ctx:    context.Background(),
				userId: 1,
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectQuery(query).WillReturnRows(rowCount)
				return sqlServer, err
			},
			want: 1,
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
				TableName: "user_activity",
			})
			count, err := init.GetTotalTodayActivity(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("user_activity.GetTotalTodayActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.want, count)
		})
	}
}
