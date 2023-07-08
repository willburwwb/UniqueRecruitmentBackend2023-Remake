package request

import "time"

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
