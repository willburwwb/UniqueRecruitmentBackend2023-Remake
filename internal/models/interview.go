package models

import (
	"time"
)

type InterviewEntity struct {
	Common
	Date          time.Time            `gorm:"not null;uniqueIndex:interviews_all"`
	Period        string               `gorm:"not null;uniqueIndex:interviews_all"` //constants.Period
	Name          string               `gorm:"not null;uniqueIndex:interviews_all"` //constants.GroupOrTeam
	SlotNumber    int                  `gorm:"column:slotNumber;not null"`
	RecruitmentID string               `gorm:"column:recruitmentId;uniqueIndex:interviews_all"`           //manytoone
	Applications  []*ApplicationEntity `gorm:"many2many:interview_selections"` //manytomany
}

func (c InterviewEntity) TableName() string {
	return "interviews"
}
