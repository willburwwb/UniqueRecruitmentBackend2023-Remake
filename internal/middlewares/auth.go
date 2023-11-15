package middlewares

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/tracer"
	"UniqueRecruitmentBackend/pkg/rerror"
	"context"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

type contextKey string

const XUID contextKey = "X-UID"
const Role contextKey = "role"

func ctxWithUID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, XUID, uid)
}

func ctxWithRole(ctx context.Context, role constants.Role) context.Context {
	return context.WithValue(ctx, Role, role)
}

func AuthMiddleware(c *gin.Context) {
	apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Authentication")
	defer span.End()

	cookie, err := c.Cookie("uid")

	if errors.Is(err, http.ErrNoCookie) {
		c.Abort()
		common.Error(c, rerror.UnauthorizedError)
		return
	}
	s := sessions.Default(c)
	u := s.Get(cookie)
	if u == nil {
		c.Abort()
		common.Error(c, rerror.UnauthorizedError)
		return
	}
	uid, ok := u.(string)
	if !ok {
		c.Abort()
		common.Error(c, rerror.UnauthorizedError)
		return
	}
	c.Request = c.Request.WithContext(ctxWithUID(apmCtx, uid))

	span.SetAttributes(attribute.String("UID", uid))
	c.Next()
}

/*
Due to session is stored in redis of sso,
I can only think of not fetching data from redis,uid is only fetched from http cookies,
and AuthMiddleware is used when deploying to the server
*/

func LocalAuthMiddleware(c *gin.Context) {
	apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Authentication")
	defer span.End()

	cookie, err := c.Cookie("uid")
	if errors.Is(err, http.ErrNoCookie) {
		c.Abort()
		common.Error(c, rerror.UnauthorizedError)
		return
	}

	uid := cookie
	c.Request = c.Request.WithContext(ctxWithUID(apmCtx, uid))
	c.Set("X-UID", uid)
	// log.Println("local auth uid", uid, "uid", c.Value("X-UID"))
	span.SetAttributes(attribute.String("UID", uid))
	c.Next()
}

func getUserRoleByUID(c *gin.Context) (constants.Role, error) {
	uid := common.GetUID(c)
	s := global.GetSSOClient()
	userDetail, err := s.GetUserInfoByUID(c, uid)
	if err != nil {
		return "", err
	}
	roles := userDetail.Roles
	for _, v := range roles {
		if v == "admin" {
			return constants.Admin, nil
		}
	}
	for _, v := range roles {
		if v == "member" {
			return constants.MemberRole, nil
		}
	}
	return constants.CandidateRole, nil
}
func SetUpUserRole(c *gin.Context) {
	apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Role")
	defer span.End()
	role, err := getUserRoleByUID(c)
	if err != nil {
		c.Abort()
		common.Error(c, rerror.CheckPermissionError)
		return
	}
	c.Request = c.Request.WithContext(ctxWithRole(apmCtx, role))
	c.Set("role", string(role))
	c.Next()
}

var GlobalRoleMiddleWare gin.HandlerFunc = SetUpUserRole

// func RoleMiddleware(roles ...constants.Role) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Role")
// 		defer span.End()

// 		uid := common.GetUID(c)
// 		client := global.GetSSOClient()

//			for _, role := range roles {
//				ok, err := client.CheckPermissionByRole(apmCtx, uid, string(role))
//				if err == nil && ok {
//					c.Request = c.Request.WithContext(ctxWithRole(apmCtx, role))
//					//log.Println(c.GetString("X-UID"), "has role", role)
//					c.Next()
//					return
//				}
//			}
//			c.Abort()
//			common.Error(c, rerror.CheckPermissionError)
//		}
//	}
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
				//log.Println(c.GetString("X-UID"), "has role", role)
				c.Next()
				return
			}
		}
		c.Abort()
		common.Error(c, rerror.CheckPermissionError)
	}
}

var CheckAdminRoleMiddleWare gin.HandlerFunc = CheckRoleMiddleware(constants.Admin)

// admin is also member

var CheckMemberRoleOrAdminMiddleWare gin.HandlerFunc = CheckRoleMiddleware(constants.MemberRole, constants.Admin)
