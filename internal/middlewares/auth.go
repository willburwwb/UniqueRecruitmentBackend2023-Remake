package middlewares

import (
	"UniqueRecruitmentBackend/pkg/grpcsso"
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func memberMiddleware(c *gin.Context) {
	uid := c.GetHeader("X-UID") //may be sso field is X-UID
	if uid == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": msg.UnauthorizedError.WithData("1", "2").Msg(),
		})
		return
	}
	user, err := grpcsso.GetUserByUID(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg":    msg.SSOError.Msg(),
			"detail": err,
		})
		return
	}
	c.Set("uid", user.Uid)
	c.Next()
}

// CandidateMiddleware used to detect whether the current user is a candidate
func CandidateMiddleware(c *gin.Context) {
	c.Next()
}

var MemberMiddleware gin.HandlerFunc = memberMiddleware
