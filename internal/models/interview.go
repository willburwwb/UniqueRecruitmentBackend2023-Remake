package models

import (
	"UniqueRecruitmentBackend/internal/constants"
	"time"
)

type InterviewEntity struct {
	Common
	Date          time.Time
	Period        constants.Period
	Name          constants.GroupOrTeam
	SlotNumber    int
	RecruitmentID uint                 //manytoone
	Applications  []*ApplicationEntity `gorm:"many2many:application_interview_selections"` //manytomany
}
type InterviewAllocation struct {
	Group time.Time
	Team  time.Time
}
