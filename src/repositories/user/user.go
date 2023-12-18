package user

import (
	"DatingApp/src/filter"
	"DatingApp/src/models"
	"DatingApp/src/repositories/base"
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type Interface interface {
	base.BaseInterface[models.UserInput, models.User, filter.UserFilter]
	GetRecomendedUser(ctx context.Context, userId int) (models.RecomendationUser, error)
}

type userRepository struct {
	base.BaseRepository[models.UserInput, models.User, filter.UserFilter]
}
type Param struct {
	Db        *sql.DB
	TableName string
}

func Init(param Param) Interface {
	return &userRepository{
		BaseRepository: base.BaseRepository[models.UserInput, models.User, filter.UserFilter]{
			Db:        param.Db,
			TableName: param.TableName,
		},
	}
}

func (r *userRepository) GetRecomendedUser(ctx context.Context, userId int) (models.RecomendationUser, error) {
	var (
		tempUser   = models.Query[models.RecomendationUser]{}
		member     = tempUser.BuildTableMember()
		query      = fmt.Sprintf(GetRecomendUser, member)
		result     = models.RecomendationUser{}
		timeNowMin = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
		timeNowMax = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 1e9, time.UTC)
	)

	rows, err := r.Db.QueryContext(ctx, query, userId, timeNowMin, timeNowMax, userId, timeNowMin, timeNowMax, userId)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		s := reflect.ValueOf(&result).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}
