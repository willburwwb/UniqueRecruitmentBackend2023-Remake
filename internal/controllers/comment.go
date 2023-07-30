package controllers

import (
	error2 "UniqueRecruitmentBackend/internal/error"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/response"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var req request.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseError(c, error2.RequestBodyError.WithDetail(err.Error()))
		return
	}
	commentId, err := models.CreateComment(&req)
	if err != nil {
		response.ResponseError(c, error2.SaveDatabaseError.WithData("comment"))
		return
	}
	response.ResponseOK(c, "create comment success", gin.H{
		"commentId": commentId,
	})
}

func DeleteComment(c *gin.Context) {
	cid := c.Param("cid")
	if err := models.DeleteCommentById(cid); err != nil {
		response.ResponseError(c, error2.SaveDatabaseError.WithData("comment"))
		return
	}
	response.ResponseOK(c, "delete comment success", nil)
}
