package request

import "mime/multipart"

type CreateApplicationRequest struct {
	Grade         string `form:"grade" json:"grade" binding:"required"`
	Institute     string `form:"institute" json:"institute" binding:"required"`
	Major         string `form:"major" json:"major" binding:"required"`
	Rank          string `form:"rank" json:"rank" binding:"required"`
	Group         string `form:"group" json:"group" binding:"required"`
	Intro         string `form:"intro" json:"intro" binding:"required"` //自我介绍
	RecruitmentID string `form:"recruitmentID" json:"recruitmentID" binding:"required"`
	Referrer      string `form:"referrer" json:"referrer"` //推荐人
	IsQuick       bool   `form:"isQuick" json:"isQuick"`   //速通

	Resume *multipart.FileHeader `form:"resume" json:"resume"` //简历
}

type UpdateApplicationRequest struct {
	Grade         string                `form:"grade" json:"grade,omitempty"`
	Institute     string                `form:"institute" json:"institute,omitempty"`
	Major         string                `form:"major" json:"major,omitempty"`
	Rank          string                `form:"rank" json:"rank,omitempty"`
	Group         string                `form:"group" json:"group,omitempty"`
	Intro         string                `form:"intro" json:"intro,omitempty"`       //自我介绍
	Referrer      string                `form:"referrer" json:"referrer,omitempty"` //推荐人
	RecruitmentID string                `form:"recruitmentID" json:"recruitmentID,omitempty"`
	Resume        *multipart.FileHeader `form:"resume" json:"resume,omitempty"` //简历
}
