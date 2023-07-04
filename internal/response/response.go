package response

import (
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(c *gin.Context, err *msg.Error) {
	c.JSON(err.StatusCode(), gin.H{
		"msg":     err.Msg(),
		"details": err.Details(),
	})
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Success!",
		"data": data,
	})
}
