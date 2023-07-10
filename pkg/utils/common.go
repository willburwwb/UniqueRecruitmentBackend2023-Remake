package utils

import "github.com/gin-gonic/gin"

func GetUserId(c *gin.Context) string {
	return c.GetString("userID")
}
