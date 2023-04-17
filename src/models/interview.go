package models

import (
	"UniqueRecruitmentBackend/constants"
	"time"
)

type InterviewEntity struct {
	Common
	Date         time.Time
	Period       constants.Period
	SlotNumber   int
	Recruitment  RecruitmentEntity
	Applications []ApplicationEntity
}
type InterviewAllocation struct {
	Group time.Time
	Team  time.Time
}
