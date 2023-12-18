package repositories

import (
	"database/sql"
	"DatingApp/src/repositories/auth"
    user "DatingApp/src/repositories/user"
    useractivity "DatingApp/src/repositories/user_activity"
    premiumfeature "DatingApp/src/repositories/premium_feature"
    
)

type Repositories struct {
	Auth auth.Interface
	User user.Interface
    UserActivity useractivity.Interface
    PremiumFeature premiumfeature.Interface
    
}

type Param struct {
	Db *sql.DB
}

func Init(param Param) *Repositories {
	return &Repositories{
		Auth: auth.Init(),
		User: user.Init(user.Param{Db: param.Db, TableName: "users"}),
        UserActivity: useractivity.Init(useractivity.Param{Db: param.Db, TableName: "user_activities"}),
        PremiumFeature: premiumfeature.Init(premiumfeature.Param{Db: param.Db, TableName: "premium_features"}),
        
	}
}
