package models

import (
	"DatingApp/src/formatter"
	"time"
)

type UserActivity struct {
	Id           int64                                 `db:"id" json:"id"`
	UserId       int                                   `db:"user_id" json:"userId"`
	PassedUserId formatter.NullableDataType[int]       `db:"passed_user_id" json:"passedUserId"`
	LikedUserId  formatter.NullableDataType[int]       `db:"liked_user_id" json:"likedUserId"`
	Status       int64                                 `db:"status" json:"status"`
	CreatedAt    formatter.NullableDataType[time.Time] `db:"created_at" json:"createdAt"`
	CreatedBy    formatter.NullableDataType[int64]     `db:"created_by" json:"createdBy"`
	UpdatedAt    formatter.NullableDataType[time.Time] `db:"updated_at" json:"updatedAt"`
	UpdatedBy    formatter.NullableDataType[int64]     `db:"updated_by" json:"updatedBy"`
	DeletedAt    formatter.NullableDataType[time.Time] `db:"deleted_at" json:"deletedAt"`
	DeletedBy    formatter.NullableDataType[int64]     `db:"deleted_by" json:"deletedBy"`
}

type UserActivityInput struct {
	UserId       int       `db:"user_id" json:"userId"`
	PassedUserId int       `db:"passed_user_id" json:"passedUserId"`
	LikedUserId  int       `db:"liked_user_id" json:"likedUserId"`
	Status       int64     `db:"status" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"-"`
	CreatedBy    int64     `db:"created_by" json:"-"`
	UpdatedAt    time.Time `db:"updated_at" json:"-"`
	UpdatedBy    int64     `db:"updated_by" json:"-"`
	DeletedAt    time.Time `db:"deleted_at" json:"-"`
	DeletedBy    int64     `db:"deleted_by" json:"-"`
}

type UserActivityInputJson struct {
	TargetUserId int `json:"targetUserId"`
}
