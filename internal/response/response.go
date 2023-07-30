package response

import (
	error2 "UniqueRecruitmentBackend/internal/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

//	func ResponseError(c *gin.Context, err *msg.Error) {
//		c.JSON(err.StatusCode(), gin.H{
//			"msg":     err.Msg(),
//			"details": err.Details(),
//		})
//	}
func ResponseError(c *gin.Context, err *error2.Error) {
	c.JSON(err.StatusCode(), gin.H{
		"msg":     err.Msg(),
		"details": err.Details(),
	})
}
func ResponseOK(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
	})
}
