package controllers

import (
	"github.com/gin-gonic/gin"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
)

// GetUserDetail get user detail.
// @Id get_user_detail
// @Summary Get user detail
// @Description Get user detail include applications and interview selections
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} common.JSONResult{data=pkg.UserDetailResp} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /user/me [get]
func GetUserDetail(c *gin.Context) {
	var (
		user *pkg.UserDetail
		apps *[]pkg.Application
		resp pkg.UserDetailResp
		err  error
	)
	defer func() { common.Resp(c, resp, err) }()

	uid := common.GetUID(c)
	user, err = grpc.GetUserInfoByUID(uid)
	if err != nil {
		return
	}

	apps, err = models.GetApplicationsByUserId(uid)
	if err != nil {
		return
	}

	resp.UserDetail = *user
	resp.Applications = *apps
	return
}
