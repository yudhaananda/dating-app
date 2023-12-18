package premiumfeature

import (
	"database/sql"
	"DatingApp/src/filter"
	"DatingApp/src/models"
	"DatingApp/src/repositories/base"
)

type Interface interface {
	base.BaseInterface[models.PremiumFeatureInput, models.PremiumFeature, filter.PremiumFeatureFilter]
}

type premiumFeatureRepository struct {
	base.BaseRepository[models.PremiumFeatureInput, models.PremiumFeature, filter.PremiumFeatureFilter]
}
type Param struct {
	Db        *sql.DB
	TableName string
}

func Init(param Param) Interface {
	return &premiumFeatureRepository{
		BaseRepository: base.BaseRepository[models.PremiumFeatureInput, models.PremiumFeature, filter.PremiumFeatureFilter]{
			Db:        param.Db,
			TableName: param.TableName,
		},
	}
}
