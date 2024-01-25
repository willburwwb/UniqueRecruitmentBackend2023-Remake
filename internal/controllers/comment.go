package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/pkg"
)

func CreateComment(c *gin.Context) {
	var (
		comment *pkg.Comment
		err     error
	)

	defer func() { common.Resp(c, comment, err) }()

	uid := common.GetUID(c)
	opts := &pkg.CreateCommentOpts{}
	if err = c.ShouldBindJSON(&opts); err != nil {
		return
	}

	opts.MemberID = uid

	comment, err = models.CreateComment(opts)
	if err != nil {
		return
	}

	return
}

func DeleteComment(c *gin.Context) {
	var (
		comment *pkg.Comment
		err     error
	)

	defer func() { common.Resp(c, nil, err) }()

	cid := c.Param("cid")
	if cid == "" {
		err = fmt.Errorf("request param error, comment id is nil")
		return
	}
	comment, err = models.GetCommentById(cid)
	if err != nil {
		return
	}

	if comment.MemberID != common.GetUID(c) {
		err = fmt.Errorf("you can't delete other's comment")
		return
	}

	err = models.DeleteCommentById(cid)
	return
}
