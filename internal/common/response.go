package common

import (
	"UniqueRecruitmentBackend/pkg/rerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, err *rerror.Error) {
	c.JSON(err.StatusCode(), gin.H{
		"msg":     err.Msg(),
		"details": err.Details(),
	})
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
	})
}
