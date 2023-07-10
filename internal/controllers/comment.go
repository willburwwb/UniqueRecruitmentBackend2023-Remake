package controllers

import (
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var req request.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseError(c, msg.RequestBodyError.WithDetail(err.Error()))
		return
	}
	commentId, err := models.CreateComment(&req)
	if err != nil {
		response.ResponseError(c, msg.SaveDatabaseError.WithData("comment"))
		return
	}
	response.ResponseOK(c, "create comment success", gin.H{
		"commentId": commentId,
	})
}

func DeleteComment(c *gin.Context) {
	cid := c.Param("cid")
	err := models.DeleteCommentById(cid)
	if err != nil {
		response.ResponseError(c, msg.SaveDatabaseError.WithData("comment"))
		return
	}
	response.ResponseOK(c, "delete comment success", nil)
}
