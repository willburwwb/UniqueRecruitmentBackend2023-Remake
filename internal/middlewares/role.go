package middlewares

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/tracer"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"github.com/gin-gonic/gin"
)

var GlobalRoleMiddleWare gin.HandlerFunc = SetUpUserRole

func SetUpUserRole(c *gin.Context) {
	apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Role")
	defer span.End()
	role, err := getUserRoleByUID(c)
	if err != nil {
		c.Abort()
		common.Error(c, rerror.CheckPermissionError)
		return
	}
	c.Request = c.Request.WithContext(common.CtxWithRole(apmCtx, role))
	c.Set("role", string(role))
	c.Next()
}

// admin is also member
var CheckMemberRoleOrAdminMiddleWare gin.HandlerFunc = CheckRoleMiddleware(constants.MemberRole, constants.Admin)
var CheckAdminRoleMiddleWare gin.HandlerFunc = CheckRoleMiddleware(constants.Admin)

func CheckRoleMiddleware(roles ...constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, role := range roles {
			var ok bool
			switch role {
			case constants.Admin:
				ok = common.IsAdmin(c)
			case constants.MemberRole:
				ok = common.IsMember(c)
			case constants.CandidateRole:
				ok = common.IsCandidate(c)
			}
			if ok {
				c.Next()
				return
			}
		}
		c.Abort()
		common.Error(c, rerror.CheckPermissionError)
	}
}

func getUserRoleByUID(c *gin.Context) (constants.Role, error) {
	uid := common.GetUID(c)
	userRoles, err := grpc.GetRolesByUID(uid)
	if err != nil {
		return "", err
	}
	for _, v := range userRoles {
		if v == "admin" {
			return constants.Admin, nil
		}
	}
	for _, v := range userRoles {
		if v == "member" {
			return constants.MemberRole, nil
		}
	}
	return constants.CandidateRole, nil
}
