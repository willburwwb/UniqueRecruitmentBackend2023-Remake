package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"UniqueRecruitmentBackend/pkg/sms"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SendSMS(c *gin.Context) {
	var req request.SendSMS
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	req.Next = constants.ZhToEnStepMap[req.Next]
	req.Current = constants.ZhToEnStepMap[req.Current]
	if req.Next == "" || req.Current == "" {
		common.Error(c, rerror.RequestBodyError.WithDetail("next or current is invalid"))
		return
	}

	var errors []string
	for _, aid := range req.Aids {
		application, err := models.GetApplicationById(aid)
		if err != nil {
			errors = append(errors, rerror.GetDatabaseError.WithData("application").Msg()+err.Error())
			continue
		}

		// check recuritment time
		recruitment, err := models.GetRecruitmentById(application.RecruitmentID, constants.MemberRole)
		if err != nil {
			errors = append(errors, rerror.GetDatabaseError.WithData("recruitment").Msg()+err.Error())
			continue
		}
		if recruitment.End.Before(time.Now()) {
			errors = append(errors, rerror.CheckPermissionError.Msg()+err.Error())
			continue
		}

		uid := common.GetUID(c)
		userInfo, err := grpc.GetUserInfoByUID(uid)
		if err != nil {
			errors = append(errors, rerror.CheckPermissionError.Msg()+err.Error())
			continue
		}
		// check applicaiton group == member group
		if !utils.CheckInGroups(userInfo.Groups, application.Group) {
			errors = append(errors, rerror.CheckPermissionError.Msg()+"member's group != application group")
			continue
		}

		if application.Abandoned {
			errors = append(errors, rerror.Abandoned.WithData(application.Uid).Msg())
			continue
		}

		if application.Rejected {
			errors = append(errors, rerror.Rejected.WithData(application.Uid).Msg())
			continue
		}

		if req.Type == constants.Accept {
			// check the interview time has been allocated
			if req.Next == string(constants.GroupInterview) && len(recruitment.FindInterviews(string(application.Group))) == 0 {
				errors = append(errors, rerror.NoInterviewScheduled.WithData(string(application.Group)).Msg())
				continue
			}
			if req.Next == string(constants.TeamInterview) && len(recruitment.FindInterviews("unique")) == 0 {
				errors = append(errors, rerror.NoInterviewScheduled.WithData("unique").Msg())
				continue
			}
		} else if req.Type == constants.Reject {
			application.Rejected = true
			// save application
			if err := models.UpdateApplicationInfo(application); err != nil {
				errors = append(errors, rerror.SaveDatabaseError.WithData("application").Msg())
			}
			continue
		} else {
			errors = append(errors, "sms type is invalid")
			continue
		}

		smsBody, err := ApplySMSTemplate(&req, userInfo, application, recruitment)
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}

		// send sms to candidate
		smsBody.Phone = userInfo.Phone
		if _, err := sms.SendSMS(*smsBody); err != nil {
			errors = append(errors, err.Error()+"send sms failed for "+userInfo.Name)
			continue
		}

	}
	if len(errors) != 0 {
		common.Error(c, rerror.SendSMSError.WithDetail(errors...))
		return
	}
	common.Success(c, "Success send sms", nil)
}

func ApplySMSTemplate(smsRequest *request.SendSMS, userInfo *grpc.UserDetail,
	application *models.ApplicationEntity, recruitment *models.RecruitmentEntity) (*sms.SMSBody, error) {

	var smsBody sms.SMSBody

	suffix := " (请勿回复本短信)"
	recruitmentName := utils.ConvertRecruitmentName(recruitment.Name)

	switch smsRequest.Type {
	case constants.Accept:
		{

			var defaultRest = ""
			switch constants.Step(smsRequest.Next) {
			//組面
			case constants.GroupInterview:
				fallthrough
			//群面
			case constants.TeamInterview:
				var allocationTime time.Time
				if smsRequest.Next == string(constants.GroupInterview) {
					allocationTime = application.InterviewAllocationsGroup
				} else if smsRequest.Next == string(constants.TeamInterview) {
					allocationTime = application.InterviewAllocationsTeam
				}
				if smsRequest.Place == "" {
					return nil, errors.New("Place is not provided for " + userInfo.Name)
				}
				if allocationTime == (time.Time{}) {
					return nil, errors.New("Interview time is not allocated for " + userInfo.Name)
				}

				// set interview time format
				// interview time get from application instead of smsRequest
				// 2006年1月2日 星期一 15时04分05秒
				formatTime := utils.ConverToLocationTime(allocationTime)
				log.Println("组面", formatTime, allocationTime)
				// FIXME
				// {1}你好，请于{2}在启明学院亮胜楼{3}参加{4}，请准时到场。
				smsBody = sms.SMSBody{
					TemplateID: constants.SMSTemplateMap[constants.PassSMS],
					Params:     []string{userInfo.Name, formatTime, smsRequest.Place, string(constants.EnToZhStepMap[smsRequest.Next])},
				}
				return &smsBody, nil
			//在线组面
			case constants.OnlineGroupInterview:
				fallthrough
			//在线群面
			case constants.OnlineTeamInterview:

				var allocationTime time.Time
				var smsTemplate constants.SMSTemplateType
				// 为什么golang没有三目运算符orz
				if smsRequest.Next == string(constants.OnlineGroupInterview) {
					allocationTime = application.InterviewAllocationsGroup
					smsTemplate = constants.OnLineGroupInterviewSMS
				} else if smsRequest.Next == string(constants.OnlineTeamInterview) {
					allocationTime = application.InterviewAllocationsTeam
					smsTemplate = constants.OnLineTeamInterviewSMS
				}
				if allocationTime == (time.Time{}) {
					return nil, errors.New("interview time is not allocated for " + userInfo.Name)
				}
				if smsRequest.MeetingId == "" {
					return nil, errors.New("meetingId is not provided for " + userInfo.Name)
				}

				// set interview time format
				// interview time get from application instead of smsRequest
				// 2006年1月2日 星期一 15时04分05秒
				formatTime := utils.ConverToLocationTime(allocationTime)

				// {1}你好，欢迎参加{2}{3}组在线群面，面试将于{4}进行，请在PC端点击腾讯会议参加面试，会议号{5}，并提前调试好摄像头和麦克风，祝你面试顺利。
				smsBody = sms.SMSBody{
					TemplateID: constants.SMSTemplateMap[smsTemplate],
					Params:     []string{userInfo.Name, recruitmentName, application.Group, formatTime, smsRequest.MeetingId},
				}
				return &smsBody, nil

			//笔试
			case constants.WrittenTest:
				fallthrough
			//熬测
			case constants.StressTest:
				if smsRequest.Place == "" {
					return nil, errors.New("place is not provided for " + userInfo.Name)
				}
				if smsRequest.Time == "" {
					return nil, errors.New("time is not provided for " + userInfo.Name)
				}

				defaultRest = fmt.Sprintf("，请于%s在%s参加%s，请务必准时到场",
					smsRequest.Time, smsRequest.Place, constants.EnToZhStepMap[smsRequest.Next])

			//通过
			case constants.Pass:
				defaultRest = fmt.Sprintf("，你已成功加入%s组", application.Group)

			//组面时间选择
			case constants.GroupTimeSelection:
				fallthrough
			//群面时间选择
			case constants.TeamTimeSelection:

				defaultRest = "，请进入选手dashboard系统选择面试时间"

			default:
				return nil, fmt.Errorf("next step %s is invalid", smsRequest.Next)
			}

			if smsRequest.Current == "" {
				return nil, errors.New("current step is not provided")
			}
			// check the customize message
			var smsResMessage string
			if smsRequest.Rest == "" {
				smsResMessage = defaultRest + suffix
			} else {
				smsResMessage = smsRequest.Rest + suffix
			}
			// {1}你好，你通过了{2}{3}组{4}审核{5}
			smsBody = sms.SMSBody{
				TemplateID: constants.SMSTemplateMap[constants.PassSMS],
				Params:     []string{userInfo.Name, recruitmentName, application.Group, constants.EnToZhStepMap[smsRequest.Current], smsResMessage},
			}
			return &smsBody, nil
		}
	case constants.Reject:
		defaultRest := "不要灰心，继续学习。期待与更强大的你的相遇！"
		if smsRequest.Current == "" {
			return nil, errors.New("current step is not provided")
		}
		var smsResMessage string
		if smsRequest.Rest == "" {
			smsResMessage = defaultRest + suffix
		} else {
			smsResMessage = smsRequest.Rest + suffix
		}
		// {1}你好，你没有通过{2}{3}组{4}审核，请你{5}
		smsBody = sms.SMSBody{
			TemplateID: constants.SMSTemplateMap[constants.Delay],
			Params:     []string{userInfo.Name, recruitmentName, application.Group, constants.EnToZhStepMap[smsRequest.Current], smsResMessage},
		}
		return &smsBody, nil
	}
	return nil, errors.New("sms step is invalid")
}

//func SendCode(c *gin.Context) {
//	req := struct {
//		Phone string      `json:"phone"`
//		Type  sms.SMSType `json:"type"`
//	}{}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		common.Error(c, rerror.RequestBodyError)
//		return
//	}
//
//	switch req.Type {
//	case sms.RegisterCode:
//		code := utils.GenerateCode()
//		sms, err := sms.SendSMS(sms.SMSBody{
//			Phone:      req.Phone,
//			TemplateID: configs.Config.SMS.RegisterCodeTemplateId,
//			Params:     []string{code},
//		})
//		if err != nil || sms.StatusCode != http.StatusOK {
//			common.Error(c, rerror.SendSMSError)
//			return
//		}
//	case sms.ResetPasswordCode:
//		code := utils.GenerateCode()
//		sms, err := sms.SendSMS(sms.SMSBody{
//			Phone:      req.Phone,
//			TemplateID: configs.Config.SMS.ResetPasswordCodeTemplateId,
//			Params:     []string{code},
//		})
//		if err != nil || sms.StatusCode != http.StatusOK {
//			common.Error(c, rerror.SendSMSError)
//			return
//		}
//		// TODO
//	}
//}
