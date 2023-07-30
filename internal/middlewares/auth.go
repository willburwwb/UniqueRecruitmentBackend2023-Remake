package middlewares

import (
	error2 "UniqueRecruitmentBackend/internal/error"
	"UniqueRecruitmentBackend/pkg/grpcsso"
	"github.com/gin-gonic/gin"
	"net/http"
)

func memberMiddleware(c *gin.Context) {
	uid := c.GetHeader("X-UID") //may be sso field is X-UID
	if uid == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": error2.UnauthorizedError.WithData("1", "2").Msg(),
		})
		return
	}
	user, err := grpcsso.GetUserInfoByUID(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg":    error2.SSOError.Msg(),
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
