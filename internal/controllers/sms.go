package controllers

import (
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/internal/services"
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SendSMS(c *gin.Context) {
	req := struct {
		Phone string           `json:"phone"`
		Type  services.SMSType `json:"type"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		zap.Error(err)
		response.ResponseError(c, msg.SendSMSError)
		return
	}

	switch req.Type {
	case services.VerificationCode:
		//code := utils.GenerateCode()
		// TODO(yuuki) finish sms
	}
}
