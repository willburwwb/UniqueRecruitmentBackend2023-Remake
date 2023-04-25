package response

import "github.com/gin-gonic/gin"

type SuccessResponser interface {
	SuccessResponse(c *gin.Context)
}
