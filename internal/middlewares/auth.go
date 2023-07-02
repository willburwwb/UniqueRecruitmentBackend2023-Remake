package middlewares

import (
	"github.com/gin-gonic/gin"
)

func MemberMiddleware(c *gin.Context) {

}

var AuthMiddleware gin.HandlerFunc = MemberMiddleware
