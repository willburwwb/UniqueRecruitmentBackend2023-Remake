package middlewares

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	error2 "UniqueRecruitmentBackend/internal/error"
	"UniqueRecruitmentBackend/internal/tracer"
	"context"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
)

func ctxWithUID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, "X-UID", uid)
}

func ctxWithRole(ctx context.Context, role constants.Role) context.Context {
	return context.WithValue(ctx, "role", role)
}

func AuthMiddleware(c *gin.Context) {
	apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Authentication")
	defer span.End()

	cookie, err := c.Cookie("uid")

	if errors.Is(err, http.ErrNoCookie) {
		c.Abort()
		common.Error(c, error2.UnauthorizedError)
		return
	}
	s := sessions.Default(c)
	u := s.Get(cookie)
	if u == nil {
		c.Abort()
		common.Error(c, error2.UnauthorizedError)
		return
	}
	uid, ok := u.(string)
	if !ok {
		c.Abort()
		common.Error(c, error2.UnauthorizedError)
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
		common.Error(c, error2.UnauthorizedError)
		return
	}

	uid := cookie
	c.Request = c.Request.WithContext(ctxWithUID(apmCtx, uid))
	c.Set("X-UID", uid)
	//log.Println("uid", uid, "uid", c.GetString("X-UID"))
	span.SetAttributes(attribute.String("UID", uid))
	c.Next()
}

func RoleMiddleware(roles ...constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Role")
		defer span.End()

		uid := common.GetUID(c)
		client := global.GetSSOClient()

		for _, role := range roles {
			ok, err := client.CheckPermissionByRole(apmCtx, uid, string(role))
			if err == nil && ok {
				c.Request = c.Request.WithContext(ctxWithRole(apmCtx, role))
				c.Next()
				return
			}
		}
		c.Abort()
		common.Error(c, error2.CheckPermissionError)
		return
	}
}

var AdminRoleMiddleWare gin.HandlerFunc = RoleMiddleware(constants.Admin)

// admin is also member

var MemberRoleOrAdminMiddleWare gin.HandlerFunc = RoleMiddleware(constants.MemberRole, constants.Admin)
