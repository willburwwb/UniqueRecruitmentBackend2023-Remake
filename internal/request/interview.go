package request

import "time"

type UpdateInterviewRequest struct {
	Date       time.Time `json:"date" form:"date" binding:"required"`
	Period     string    `json:"period" form:"period" binding:"required"`
	Name       string    `json:"name" form:"name" binding:"required"`
	SlotNumber int       `json:"slotNumber" form:"slotNumber" binding:"required"`
}
