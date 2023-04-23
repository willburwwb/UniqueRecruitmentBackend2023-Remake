package models

import (
	"time"
)

type RecruitmentEntity struct {
	Common
	Name         string
	Beginning    time.Time
	Deadline     time.Time
	End          time.Time
	Interviews   []InterviewEntity   `gorm:"foreignKey:RecruitmentEntityID;references:Uid;constraint:OnDelete:CASCADE;"` //一个hr->面试
	Applications []ApplicationEntity `gorm:"foreignKey:RecruitmentEntityID;references:Uid;constraint:OnDelete:CASCADE;"` //一个hr->简历
}
