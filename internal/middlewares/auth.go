package middlewares

import (
	"UniqueRecruitmentBackend/pkg/grpcsso"
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MemberMiddleware(c *gin.Context) {
	uid := c.GetHeader("X-UID") //may be sso field is X-UID
	if uid == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": msg.UnauthorizedError.WithData("1", "2").Msg(),
		})
	}
	user, err := grpcsso.GetUserByUID(uid)
	if err != nil {
		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "无法获取用户信息"})
		return
	}
	c.Set("uid", user.Uid)
	c.Next()
}

var AuthMiddleware gin.HandlerFunc = MemberMiddleware
