package models

import (
	"UniqueRecruitmentBackend/internal/constants"
)

type CandidateEntity struct {
	Common
	Applications []ApplicationEntity `gorm:"foreignKey:CandidateID;references:Uid;constraint:OnDelete:CASCADE;"` //onetomany
	Name         string
	Password     string
	Phone        string
	Mail         string
	Gender       constants.Gender
}

// maybe used
type Password struct {
}
