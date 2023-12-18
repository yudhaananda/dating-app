package premiumfeature

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"regexp"
	"DatingApp/src/filter"
	"DatingApp/src/formatter"
	"DatingApp/src/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO premium_feature () VALUES ()")

	type args struct {
		ctx    context.Context
		models models.Query[models.PremiumFeatureInput]
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				TableName: "premium_feature",
			})
			err = init.Create(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("premium_feature.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE premium_feature SET  WHERE ")

	type args struct {
		ctx    context.Context
		models models.Query[models.PremiumFeatureInput]
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				models: models.Query[models.PremiumFeatureInput]{},
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
				TableName: "premium_feature",
			})
			err = init.Update(tt.args.ctx, tt.args.models, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("premium_feature.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tempModels := models.Query[models.PremiumFeature]{}
	member := tempModels.BuildTableMember()
	query := regexp.QuoteMeta("SELECT " + member + " FROM premium_feature WHERE 1=1")
	queryCount := regexp.QuoteMeta("SELECT COUNT(*) FROM premium_feature")
	mockTime := time.Date(2022, 5, 11, 0, 0, 0, 0, time.UTC)

	type args struct {
		ctx    context.Context
		models filter.Paging[filter.PremiumFeatureFilter]
	}
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		wantPremiumFeature    []models.PremiumFeature
		wantCount   int
		wantErr     bool
	}{
		{
			name: "sql count query failed",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.PremiumFeatureFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				sqlMock.ExpectQuery(queryCount).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantPremiumFeature: []models.PremiumFeature{},
			wantErr:  true,
		},
		{
			name: "sql query failed",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.PremiumFeatureFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				sqlMock.ExpectQuery(query).WillReturnError(errors.New(""))
				return sqlServer, err
			},
			wantErr:   true,
			wantPremiumFeature:  []models.PremiumFeature{},
			wantCount: 1,
		},
		{
			name: "sql success",
			args: args{
				ctx:    context.Background(),
				models: filter.Paging[filter.PremiumFeatureFilter]{},
			},
			prepSqlMock: func() (*sql.DB, error) {
				sqlServer, sqlMock, err := sqlmock.New()
				rowCount := sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				sqlMock.ExpectQuery(queryCount).WillReturnRows(rowCount)
				row := sqlMock.NewRows([]string{"id", "name", "flag",  "status", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"})
				row.AddRow(1, "test", "test",  1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1, formatter.NullableDataType[time.Time]{Valid: true, Data: mockTime}, 1)
				sqlMock.ExpectQuery(query).WillReturnRows(row)
				return sqlServer, err
			},
			wantPremiumFeature: []models.PremiumFeature{
				{
					Id:       1,
					Name: "test", 
Flag: "test", 

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
				TableName: "premium_feature",
			})
			premiumFeatures, count, err := init.Get(tt.args.ctx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("premium_feature.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantPremiumFeature, premiumFeatures)
			assert.Equal(t, tt.wantCount, count)
		})
	}
}
