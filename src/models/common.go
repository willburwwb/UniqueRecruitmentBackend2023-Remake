package models

import "time"

type Common struct {
	ID       string //UUid
	CreateAt time.Time
	UpdateAt time.Time
}
