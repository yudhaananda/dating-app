package services

import (
	"DatingApp/src/repositories"
	"DatingApp/src/services/auth"
	premiumfeature "DatingApp/src/services/premium_feature"
	user "DatingApp/src/services/user"
	useractivity "DatingApp/src/services/user_activity"
)

type Services struct {
	Auth           auth.Interface
	User           user.Interface
	UserActivity   useractivity.Interface
	PremiumFeature premiumfeature.Interface
}

type Param struct {
	Repositories *repositories.Repositories
}

func Init(param Param) *Services {
	return &Services{
		Auth: auth.Init(auth.Param{UserRepository: param.Repositories.User, AuthRepository: param.Repositories.Auth}),
		User: user.Init(user.Param{
			UserRepository: param.Repositories.User,
		},
		),
		UserActivity: useractivity.Init(useractivity.Param{
			UserActivityRepository:   param.Repositories.UserActivity,
			UserRepository:           param.Repositories.User,
			PremiumFeatureRepository: param.Repositories.PremiumFeature,
		},
		),
		PremiumFeature: premiumfeature.Init(premiumfeature.Param{
			PremiumFeatureRepository: param.Repositories.PremiumFeature,
		},
		),
	}
}
