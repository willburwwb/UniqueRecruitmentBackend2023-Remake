package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"github.com/gin-gonic/gin"
)

func GetUserDetail(c *gin.Context) {
	var (
		user *pkg.UserDetail
		err  error
	)
	defer func() { common.Resp(c, user, err) }()

	uid := common.GetUID(c)
	user, err = grpc.GetUserInfoByUID(uid)
	if err != nil {
		return
	}

	return
}
