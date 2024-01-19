package middlewares

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/tracer"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"errors"
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
var CheckMemberRoleOrAdminMiddleWare gin.HandlerFunc = CheckRoleMiddleware(pkg.MemberRole, pkg.Admin)
var CheckAdminRoleMiddleWare gin.HandlerFunc = CheckRoleMiddleware(pkg.Admin)

func CheckRoleMiddleware(roles ...pkg.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, role := range roles {
			var ok bool
			switch role {
			case pkg.Admin:
				ok = common.IsAdmin(c)
			case pkg.MemberRole:
				ok = common.IsMember(c)
			case pkg.CandidateRole:
				ok = common.IsCandidate(c)
			}
			if ok {
				c.Next()
				return
			}
		}
		c.Abort()
		common.Resp(c, nil, errors.New("check permission error"))
	}
}

func getUserRoleByUID(c *gin.Context) (pkg.Role, error) {
	uid := common.GetUID(c)
	userRoles, err := grpc.GetRolesByUID(uid)
	if err != nil {
		return "", err
	}
	for _, v := range userRoles {
		if v == "admin" {
			return pkg.Admin, nil
		}
	}
	for _, v := range userRoles {
		if v == "member" {
			return pkg.MemberRole, nil
		}
	}
	return pkg.CandidateRole, nil
}
