package models

import (
	"UniqueRecruitmentBackend/internal/constants"
)

type ApplicationEntity struct {
	Common
	Grade               constants.Grade
	Institute           string
	Major               string
	Rank                string
	Group               constants.Group
	Intro               string
	IsQuick             bool
	Referrer            string
	Resume              string
	Abandoned           bool
	Rejected            bool
	Step                string
	InterviewAllocation InterviewAllocation

	InterviewSelections []*InterviewEntity `gorm:"many2many:application_interview_selections"` //manytomany
	CandidateID         uint               //manytoone
	RecruitmentID       uint               //manytoone
	Comments            []CommentEntity    `gorm:"foreignKey:ApplicationID;references:Uid;constraint:OnDelete:CASCADE;"` //onetomany
}
