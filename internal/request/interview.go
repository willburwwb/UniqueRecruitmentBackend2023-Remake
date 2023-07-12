package request

import "time"

type UpdateInterviewRequest struct {
	Uid        string    `json:"uid" form:"uid" `
	Date       time.Time `json:"date" form:"date" binding:"required"`
	Period     string    `json:"period" form:"period" binding:"required"`
	SlotNumber int       `json:"slotNumber" form:"slotNumber" binding:"required"`
}
