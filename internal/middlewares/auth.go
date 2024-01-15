package middlewares

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/tracer"
	"UniqueRecruitmentBackend/pkg/rerror"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

func AuthMiddleware(c *gin.Context) {
	apmCtx, span := tracer.Tracer.Start(c.Request.Context(), "Authentication")
	defer span.End()

	cookie, err := c.Cookie("SSO_SESSION") // only for check
	if configs.Config.Server.RunMode == "debug" {
		if cookie == "unique_web_admin" {
			c.Request = c.Request.WithContext(common.CtxWithUID(apmCtx, "ffb6e834-3615-4ebb-9d9d-825af333a3ca"))
			span.SetAttributes(attribute.String("UID", "ffb6e834-3615-4ebb-9d9d-825af333a3ca"))
			c.Set("X-UID", "ffb6e834-3615-4ebb-9d9d-825af333a3ca")
			c.Next()
			return
		}
		if cookie == "unique_web_candidate" {
			c.Request = c.Request.WithContext(common.CtxWithUID(apmCtx, "afb6e834-3615-4ebb-9d9d-825af333a3ca"))
			span.SetAttributes(attribute.String("UID", "afb6e834-3615-4ebb-9d9d-825af333a3ca"))
			c.Set("X-UID", "afb6e834-3615-4ebb-9d9d-825af333a3ca")
			c.Next()
			return
		}
	}

	if errors.Is(err, http.ErrNoCookie) {
		c.Abort()
		common.Error(c, rerror.UnauthorizedError)
		return
	}
	s := sessions.Default(c)
	u := s.Get(constants.SessionNameUID)
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
	c.Request = c.Request.WithContext(common.CtxWithUID(apmCtx, uid))
	c.Set("X-UID", uid)
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
	c.Request = c.Request.WithContext(common.CtxWithUID(apmCtx, uid))
	c.Set("X-UID", uid)
	// log.Println("local auth uid", uid, "uid", c.Value("X-UID"))
	span.SetAttributes(attribute.String("UID", uid))
	c.Next()
}
