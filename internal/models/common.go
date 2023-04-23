package models

import (
	"time"
)

type Common struct {
	Uid       string `gorm:"column:Uid;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
