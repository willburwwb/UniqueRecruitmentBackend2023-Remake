package middlewares

import (
	"UniqueRecruitmentBackend/internal/constants"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		sess.Get(constants.SessionNameUID)
	}

}
