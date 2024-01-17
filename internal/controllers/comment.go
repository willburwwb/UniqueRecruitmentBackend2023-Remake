package controllers

import (
	"github.com/gin-gonic/gin"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/rerror"
)

func CreateComment(c *gin.Context) {
	var req pkg.CreateCommentOpts
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	commentId, err := models.CreateComment(&req)
	if err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData("comment"))
		return
	}

	common.Success(c, "create comment success", gin.H{
		"commentId": commentId,
	})
}

func DeleteComment(c *gin.Context) {
	cid := c.Param("cid")
	if err := models.DeleteCommentById(cid); err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData("comment"))
		return
	}
	common.Success(c, "delete comment success", nil)
}
