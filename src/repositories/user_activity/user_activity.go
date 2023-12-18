package useractivity

import (
	"DatingApp/src/filter"
	"DatingApp/src/models"
	"DatingApp/src/repositories/base"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Interface interface {
	base.BaseInterface[models.UserActivityInput, models.UserActivity, filter.UserActivityFilter]
	GetTotalTodayActivity(ctx context.Context, userId int) (int, error)
}

type userActivityRepository struct {
	base.BaseRepository[models.UserActivityInput, models.UserActivity, filter.UserActivityFilter]
}
type Param struct {
	Db        *sql.DB
	TableName string
}

func Init(param Param) Interface {
	return &userActivityRepository{
		BaseRepository: base.BaseRepository[models.UserActivityInput, models.UserActivity, filter.UserActivityFilter]{
			Db:        param.Db,
			TableName: param.TableName,
		},
	}
}

func (r *userActivityRepository) GetTotalTodayActivity(ctx context.Context, userId int) (int, error) {
	var (
		timeNowMin = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
		timeNowMax = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 1e9, time.UTC)
		where      = fmt.Sprintf(" WHERE user_id = %d AND created_at > '%s' AND created_at < '%s'", userId, timeNowMin, timeNowMax)
		count      int
	)

	rowCount, err := r.Db.QueryContext(ctx, GetTotalTodayActivity+where)
	if err != nil {
		return count, err
	}
	for rowCount.Next() {
		err = rowCount.Scan(&count)
	}
	return count, err
}
