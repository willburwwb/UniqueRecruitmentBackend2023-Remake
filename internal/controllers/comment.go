package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/pkg"
)

// CreateComment create comment
// @Id create_comment.
// @Summary create comment for application
// @Description create comment for applications, only can be created by member.
// @Tags comment
// @Accept  json
// @Produce  json
// @Param 	pkg.CreateCommentOpts body pkg.CreateCommentOpts true "create comment opts"
// @Success 200 {object} common.JSONResult{data=pkg.Comment} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /comments [POST]
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

// DeleteComment delete comment
// @Id delete_comment.
// @Summary delete comment of application
// @Description delete comment of application, only can be deleted by comment's owner.
// @Tags comment
// @Accept  json
// @Produce  json
// @Param 	cid path string true "comment uid"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /comments/{cid} [DELETE]
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
