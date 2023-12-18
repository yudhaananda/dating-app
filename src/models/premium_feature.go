package models

import (
	"DatingApp/src/formatter"
	"time"
)

type PremiumFeature struct {
	Id        int64                                 `db:"id" json:"id"`
	Name      string                                `db:"name" json:"name"`
	Flag      string                                `db:"flag" json:"flag"`
	Status    int64                                 `db:"status" json:"status"`
	CreatedAt formatter.NullableDataType[time.Time] `db:"created_at" json:"createdAt"`
	CreatedBy formatter.NullableDataType[int64]     `db:"created_by" json:"createdBy"`
	UpdatedAt formatter.NullableDataType[time.Time] `db:"updated_at" json:"updatedAt"`
	UpdatedBy formatter.NullableDataType[int64]     `db:"updated_by" json:"updatedBy"`
	DeletedAt formatter.NullableDataType[time.Time] `db:"deleted_at" json:"deletedAt"`
	DeletedBy formatter.NullableDataType[int64]     `db:"deleted_by" json:"deletedBy"`
}

type PremiumFeatureInput struct {
	Name      string    `db:"name" json:"name"`
	Flag      string    `db:"flag" json:"flag"`
	Status    int64     `db:"status" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	CreatedBy int64     `db:"created_by" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
	UpdatedBy int64     `db:"updated_by" json:"-"`
	DeletedAt time.Time `db:"deleted_at" json:"-"`
	DeletedBy int64     `db:"deleted_by" json:"-"`
}
