package models

import (
	"time"
)

type ApplicationEntity struct {
	Common
	Grade                     string `gorm:"not null"` //constants.Grade
	Institute                 string `gorm:"not null"`
	Major                     string `gorm:"not null"`
	Rank                      string `gorm:"not null"` //constants.Rank
	Group                     string `gorm:"not null"` //constants.Group
	Intro                     string `gorm:"not null"`
	IsQuick                   bool   `gorm:"column:isQuick;not null"`
	Referrer                  string
	Resume                    string
	Abandoned                 bool               `gorm:"not null; default false" `
	Rejected                  bool               `gorm:"not null; default false"`
	Step                      string             `gorm:"not null"`                                                 //constants.Step
	CandidateID               string             `gorm:"column:candidateId;uniqueIndex:UQ_CandidateID_RecruitmentID"`   //manytoone
	RecruitmentID             string             `gorm:"column:recruitmentId;uniqueIndex:UQ_CandidateID_RecruitmentID"` //manytoone
	InterviewAllocationsGroup time.Time          `gorm:"column:interviewAllocationsGroup;"`
	InterviewAllocationsTeam  time.Time          `gorm:"column:interviewAllocationsTeam;"`
	InterviewSelections       []*InterviewEntity `gorm:"many2many:interview_selections; constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //manytomany
	Comments                  []CommentEntity    `gorm:"foreignKey:ApplicationID; references:Uid; constraint:OnDelete:CASCADE;"`        //onetomany
}

func (a ApplicationEntity) TableName() string {
	return "applications"
}

/*

 */
