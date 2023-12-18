package models

import (
	"DatingApp/src/formatter"
	"time"
)

type User struct {
	Id               int64                                 `db:"id" json:"id"`
	UserName         string                                `db:"user_name" json:"userName"`
	Password         string                                `db:"password" json:"password"`
	Image            formatter.NullableDataType[string]    `db:"image" json:"image"`
	PremiumFeatureId formatter.NullableDataType[int]       `db:"premium_feature_id" json:"premiumFeatureId"`
	Status           int64                                 `db:"status" json:"status"`
	CreatedAt        formatter.NullableDataType[time.Time] `db:"created_at" json:"createdAt"`
	CreatedBy        formatter.NullableDataType[int64]     `db:"created_by" json:"createdBy"`
	UpdatedAt        formatter.NullableDataType[time.Time] `db:"updated_at" json:"updatedAt"`
	UpdatedBy        formatter.NullableDataType[int64]     `db:"updated_by" json:"updatedBy"`
	DeletedAt        formatter.NullableDataType[time.Time] `db:"deleted_at" json:"deletedAt"`
	DeletedBy        formatter.NullableDataType[int64]     `db:"deleted_by" json:"deletedBy"`
}

type UserInput struct {
	UserName         string    `db:"user_name" json:"userName"`
	Password         string    `db:"password" json:"password"`
	Image            string    `db:"image" json:"image"`
	PremiumFeatureId int       `db:"premium_feature_id" json:"-"`
	Status           int64     `db:"status" json:"-"`
	CreatedAt        time.Time `db:"created_at" json:"-"`
	CreatedBy        int64     `db:"created_by" json:"-"`
	UpdatedAt        time.Time `db:"updated_at" json:"-"`
	UpdatedBy        int64     `db:"updated_by" json:"-"`
	DeletedAt        time.Time `db:"deleted_at" json:"-"`
	DeletedBy        int64     `db:"deleted_by" json:"-"`
}

type Subscribe struct {
	PremiumFeatureId int `json:"premiumFeatureId"`
}

type RecomendationUser struct {
	Id               int64   `db:"id" json:"id"`
	UserName         string  `db:"user_name" json:"userName"`
	PremiumFeatureId *int    `db:"premium_feature_id" json:"premiumFeatureId"`
	Image            *string `db:"image" json:"image"`
}
