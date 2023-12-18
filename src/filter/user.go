package filter

type UserFilter struct {
	Id               int    `db:"id" json:"id" form:"id"`
	UserName         string `db:"user_name" json:"userName" form:"userName"`
	Password         string `db:"password" json:"password" form:"password"`
	PremiumFeatureId int    `db:"premium_feature_id" json:"premiumFeatureId" form:"premiumFeatureId"`
}
