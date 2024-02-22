package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/tracer"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
)

// GetUserDetail get user detail.
// @Id get_user_detail
// @Summary Get user detail
// @Description Get user detail include applications and interview selections (without comments)
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

	apmCtx, span := tracer.Tracer.Start(c, "GetUserDetail")
	defer span.End()

	//	spanContext := span.SpanContext()
	//	zapx.Infof("Span TraceID: %s, SpanID: %s", spanContext.TraceID().String(), spanContext.SpanID().String())

	uid := common.GetUID(c)
	user, err = grpc.GetUserInfoByUID(uid)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get user info failed", zap.String("UID", uid))
		return
	}

	apps, err = models.GetApplicationsByUserId(uid)
	if err != nil {
		span.RecordError(err)
		zapx.WithContext(apmCtx).Error("get application failed", zap.String("UID", uid))
		return
	}

	resp.UserDetail = *user
	resp.Applications = *apps
	return
}

// GetMembersDetail get members detail.
// @Id get_members_detail
// @Summary Get members detail
// @Description Get members detail
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} common.JSONResult{data=pkg.MembersDetail} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /user/me [get]
func GetMembersDetail(c *gin.Context) {

}
