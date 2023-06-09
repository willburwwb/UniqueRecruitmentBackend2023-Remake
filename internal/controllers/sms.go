package controllers

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/internal/services"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendSMS(c *gin.Context) {
	req := struct {
		Phone string           `json:"phone"`
		Type  services.SMSType `json:"type"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseError(c, msg.RequestBodyError)
		return
	}

	switch req.Type {
	case services.RegisterCode:
		code := utils.GenerateCode()
		sms, err := services.SendSMS(services.SMSBody{
			Phone:      req.Phone,
			TemplateID: configs.Config.SMS.RegisterCodeTemplateId,
			Params:     []string{code},
		})
		if err != nil || sms.StatusCode != http.StatusOK {
			response.ResponseError(c, msg.SendSMSError)
			return
		}
	case services.ResetPasswordCode:
		code := utils.GenerateCode()
		sms, err := services.SendSMS(services.SMSBody{
			Phone:      req.Phone,
			TemplateID: configs.Config.SMS.ResetPasswordCodeTemplateId,
			Params:     []string{code},
		})
		if err != nil || sms.StatusCode != http.StatusOK {
			response.ResponseError(c, msg.SendSMSError)
			return
		}
		// TODO
	}
}
