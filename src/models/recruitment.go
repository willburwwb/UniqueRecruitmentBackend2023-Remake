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
	Interviews   []InterviewEntity
	Applications []ApplicationEntity
	//Statistics [] ??是啥
}
