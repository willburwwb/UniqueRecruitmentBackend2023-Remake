package request

import (
	"UniqueRecruitmentBackend/internal/constants"
	"time"
)

// type UpdateInterviewRequest struct {
// 	Uid        string           `json:"uid" form:"uid" `
// 	Date       time.Time        `json:"date" form:"date" binding:"required"`
// 	Period     constants.Period `json:"period" form:"period" binding:"required"`
// 	SlotNumber int              `json:"slotNumber" form:"slotNumber" binding:"required"`
// }

type CreateInterview struct {
	Date       time.Time        `json:"date" form:"date" binding:"required"`
	Period     constants.Period `json:"period" form:"period" binding:"required"`
	SlotNumber int              `json:"slotNumber" form:"slotNumber" binding:"required"`
}
type UpdateInterview struct {
	Uid        string           `json:"uid" form:"uid" `
	Date       time.Time        `json:"date" form:"date"`
	Period     constants.Period `json:"period" form:"period"`
	SlotNumber int              `json:"slotNumber" form:"slotNumber"`
}
type DeleteInterviewUID string
