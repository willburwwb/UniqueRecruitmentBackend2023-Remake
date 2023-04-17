package models

import (
	"UniqueRecruitmentBackend/constants"
)

type CandidateEntity struct {
	Common
	Applications []ApplicationEntity
	Name         string
	Password     Password
	Phone        string
	Mail         string
	Gender       constants.Gender
}
type Password struct {
}
