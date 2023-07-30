package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	error2 "UniqueRecruitmentBackend/internal/error"
	"UniqueRecruitmentBackend/internal/models"
	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	ApplicationID string `json:"applicationId"`
	MemberID      string `json:"memberId"`
	Content       string `json:"content"`
	Evaluation    int    `json:"evaluation"`
}

func CreateComment(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, error2.RequestBodyError.WithDetail(err.Error()))
		return
	}
	commentId, err := models.CreateComment(&req)
	if err != nil {
		common.Error(c, error2.SaveDatabaseError.WithData("comment"))
		return
	}
	common.Success(c, "create comment success", gin.H{
		"commentId": commentId,
	})
}

func DeleteComment(c *gin.Context) {
	cid := c.Param("cid")
	if err := models.DeleteCommentById(cid); err != nil {
		common.Error(c, error2.SaveDatabaseError.WithData("comment"))
		return
	}
	common.Success(c, "delete comment success", nil)
}
