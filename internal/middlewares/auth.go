package middlewares

import (
	"UniqueRecruitmentBackend/internal/common"
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
	c.Request = c.Request.WithContext(common.CtxWithUID(apmCtx, uid))

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
