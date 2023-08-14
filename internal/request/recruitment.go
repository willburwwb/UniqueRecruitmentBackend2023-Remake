package request

import (
	"UniqueRecruitmentBackend/internal/constants"
	"time"
)

type CreateRecruitment struct {
	Name      string    `json:"name" binding:"required"`
	Beginning time.Time `json:"beginning" binding:"required"`
	Deadline  time.Time `json:"deadline" binding:"required"`
	End       time.Time `json:"end" binding:"required"`
}

type UpdateRecruitment struct {
	Beginning time.Time `json:"beginning"`
	Deadline  time.Time `json:"deadline"`
	End       time.Time `json:"end"`
}

type InterviewInfo struct {
	Id         string           `json:"id"`
	Date       time.Time        `json:"date"`
	Period     constants.Period `json:"period"`
	SlotNumber int              `json:"slotNumber"`
}

type SetRecruitmentInterviewTime struct {
	Interviews []InterviewInfo
}
