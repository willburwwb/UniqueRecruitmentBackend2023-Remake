package response

import (
	"UniqueRecruitmentBackend/pkg/msg"
	"UniqueRecruitmentBackend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ResponseRecruitmentNotStart(c *gin.Context, msgArgs ...interface{}) {
	c.JSON(msg.R_NOT_STARTED.StatusCode(), gin.H{
		"msg":  utils.FormatMsg(msg.R_NOT_STARTED.Msg, msgArgs),
		"data": msg.R_NOT_STARTED.Data,
	})
}
func ResponseRecruitmentEnded(c *gin.Context, msgArgs ...interface{}) {
	c.JSON(msg.R_ENDED.StatusCode(), gin.H{
		"msg":  utils.FormatMsg(msg.R_ENDED.Msg, msgArgs),
		"data": msg.R_ENDED.Data,
	})
}
func ResponseRecruitmentLongEnded(c *gin.Context, msgArgs ...interface{}) {
	c.JSON(msg.R_ENDED.StatusCode(), gin.H{
		"msg":  utils.FormatMsg(msg.R_ENDED_LONG.Msg, msgArgs),
		"data": msg.R_ENDED_LONG.Data,
	})
}

