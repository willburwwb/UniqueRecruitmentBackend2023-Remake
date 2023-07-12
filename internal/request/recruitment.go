package request

import (
	"UniqueRecruitmentBackend/internal/constants"
	"time"
)

type CreateRecruitmentRequest struct {
	Name      string    `json:"name" binding:"required"`
	Beginning time.Time `json:"beginning" binding:"required"`
	Deadline  time.Time `json:"deadline" binding:"required"`
	End       time.Time `json:"end" binding:"required"`
}

type UpdateRecruitmentRequest struct {
	Beginning time.Time `json:"beginning" binding:"required"`
	Deadline  time.Time `json:"deadline" binding:"required"`
	End       time.Time `json:"end" binding:"required"`
}

type InterviewInfo struct {
	Id         string           `json:"id"`
	Date       time.Time        `json:"date"`
	Period     constants.Period `json:"period"`
	SlotNumber int              `json:"slotNumber"`
}

type SetRecruitmentInterviewTimeRequest struct {
	Interviews []InterviewInfo
}
