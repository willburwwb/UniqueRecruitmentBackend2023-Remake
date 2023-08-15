package models

import "time"

type Common struct {
	//缺少uuid-ossp 无法使用uuid_generate_v4	->   gen_random_uuid
	Uid       string    `gorm:"column:uid;type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time `gorm:"column:createdAt;not null"` //为了和原数据库匹配，gorm默认会将其在数据表中默认为蛇形q
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;index"`
}
