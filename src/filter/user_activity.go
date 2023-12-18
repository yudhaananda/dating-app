package filter

type UserActivityFilter struct {
	Id           int `db:"id" json:"id" form:"id"`
	UserId       int `db:"user_id" json:"userId" form:"userId"`
	PassedUserId int `db:"passed_user_id" json:"passedUserId" form:"passedUserId"`
	LikedUserId  int `db:"liked_user_id" json:"likedUserId" form:"likedUserId"`
}
