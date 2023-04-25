package models

import (
	"time"

	"github.com/jackc/pgx/pgtype"
)

type RecruitmentEntity struct {
	Common
	Name         string              `gorm:"not null;unique"`
	Beginning    time.Time           `gorm:"not null"`
	Deadline     time.Time           `gorm:"not null"`
	End          time.Time           `gorm:"not null"`
	Statistics   pgtype.JSONB        `gorm:"type:jsonb"`
	Interviews   []InterviewEntity   `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE;"` //一个hr->面试
	Applications []ApplicationEntity `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE;"` //一个hr->简历
}

func (c RecruitmentEntity) TableName() string {
	return "recruitments"
}
