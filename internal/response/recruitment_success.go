package response

import (
	"UniqueRecruitmentBackend/pkg/msg"

	"github.com/gin-gonic/gin"
)

type ResponseGetRecruitmentDetailBody struct {
}

func ResponseRecruitmentSuccess(c *gin.Context, data interface{}) {
	c.JSON(msg.R_SUCCESS.StatusCode(), gin.H{
		"msg":  msg.R_SUCCESS.Msg,
		"data": data,
	})
}

// func (resp *ResponseGetRecruitmentDetailBody) SuccessResponse(c *gin.Context) {
// 	c.JSON(msg.R_SUCCESS.StatusCode(), gin.H{
// 		"msg":  msg.R_SUCCESS.Msg,
// 		"data": *resp,
// 	})
// }
