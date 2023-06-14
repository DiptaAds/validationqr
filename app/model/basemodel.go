package model

import "time"

type BaseModel struct {
	ID        int64      `bson:"id"`
	CreatedAt time.Time  `bson:"CreatedAt"`
	UpdateAt  *time.Time `bson:"UpdateAt"`
	DeleteAt  *time.Time `bson:"DeleteAt"`
}
