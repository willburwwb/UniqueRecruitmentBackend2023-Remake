package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Resp(c *gin.Context, data interface{}, err error) {
	if err != nil {
		errResp(c, err)
	} else {
		if data == nil {
			data = map[string]string{}
		}
		successResp(c, data)
	}
}

func successResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, JSONResult{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func errResp(c *gin.Context, err error) {
	errRespWithCode(c, -1, err.Error())
}

func errRespWithCode(c *gin.Context, code int, errMsg string) {
	c.JSON(http.StatusBadRequest, JSONResult{
		Code: code,
		Msg:  errMsg,
		Data: map[string]string{},
	})
}
