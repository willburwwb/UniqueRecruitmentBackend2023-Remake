package controllers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/sms"
)

func SendSMS(c *gin.Context) {
	var (
		app  *pkg.Application
		r    *pkg.Recruitment
		user *pkg.UserDetail
		err  error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.SendSMSOpts{}
	if err := c.ShouldBind(&opts); err != nil {
		return
	}

	if err = opts.Validate(); err != nil {
		return
	}

	uid := common.GetUID(c)
	user, err = grpc.GetUserInfoByUID(uid)
	if err != nil {
		return
	}

	app, err = models.GetApplicationByIdForCandidate(opts.Aids[0])
	if err != nil {
		return
	}

	// judge whether the recruitment has expired
	r, err = models.GetRecruitmentById(app.RecruitmentID)
	if err != nil {
		return
	}
	if r.End.Before(time.Now()) {
		return
	}

	var errors []string
	for _, aid := range opts.Aids {
		app, err = models.GetApplicationByIdForCandidate(aid)
		if err != nil {
			errors = append(errors, fmt.Sprintf("get application %s failed, error: %s", aid, err.Error()))
			continue
		}

		// check applicaiton group == member group
		if !utils.CheckInGroups(user.Groups, app.Group) {
			errors = append(errors, fmt.Sprintf("send application %s sms failed, error: you are not in the same group", aid))
			continue
		}

		if app.Abandoned {
			errors = append(errors, fmt.Sprintf("application of %s has already been abandoned", aid))
			continue
		}

		if app.Rejected {
			errors = append(errors, fmt.Sprintf("application of %s has already been rejected", aid))
			continue
		}

		if opts.Type == pkg.Accept {
			// check the interview time has been allocated
			if opts.Next == string(pkg.GroupInterview) && len(r.GetInterviews(string(app.Group))) == 0 {
				errors = append(errors, fmt.Sprintf("no interviews are scheduled for %s", app.Group))
				continue
			}
			if opts.Next == string(pkg.TeamInterview) && len(r.GetInterviews("unique")) == 0 {
				errors = append(errors, fmt.Sprintf("no interviews are scheduled for unique"))
				continue
			}
		} else if opts.Type == pkg.Reject {
			app.Rejected = true
			// save application
			if err := models.UpdateApplicationInfo(app); err != nil {
				errors = append(errors, fmt.Sprintf("update application %s failed, error: %s", aid, err.Error()))
			}
			continue
		} else {
			errors = append(errors, "sms type is invalid")
			continue
		}

		var smsBody *sms.SMSBody
		smsBody, err = ApplySMSTemplate(opts, user, app, r)
		if err != nil {
			errors = append(errors, fmt.Sprintf("set smsbody for user %s failed, error: %s", user.Name, err.Error()))
			continue
		}

		// send sms to candidate
		smsBody.Phone = user.Phone
		if _, err := sms.SendSMS(*smsBody); err != nil {
			errors = append(errors, fmt.Sprintf("send sms for user %s failed, error: %s", user.Name, err.Error()))
			continue
		}

	}
	if len(errors) != 0 {
		err = fmt.Errorf("%v", errors)
		return
	}
	return
}

func ApplySMSTemplate(smsRequest *pkg.SendSMSOpts, userInfo *pkg.UserDetail,
	application *pkg.Application, recruitment *pkg.Recruitment) (*sms.SMSBody, error) {

	var smsBody sms.SMSBody

	suffix := " (请勿回复本短信)"
	recruitmentName := utils.ConvertRecruitmentName(recruitment.Name)

	switch smsRequest.Type {
	case pkg.Accept:
		{

			var defaultRest = ""
			switch pkg.Step(smsRequest.Next) {
			//組面
			case pkg.GroupInterview:
				fallthrough
			//群面
			case pkg.TeamInterview:
				var allocationTime time.Time
				if smsRequest.Next == string(pkg.GroupInterview) {
					allocationTime = application.InterviewAllocationsGroup
				} else if smsRequest.Next == string(pkg.TeamInterview) {
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
					TemplateID: pkg.SMSTemplateMap[pkg.PassSMS],
					Params:     []string{userInfo.Name, formatTime, smsRequest.Place, string(pkg.EnToZhStepMap[smsRequest.Next])},
				}
				return &smsBody, nil
			//在线组面
			case pkg.OnlineGroupInterview:
				fallthrough
			//在线群面
			case pkg.OnlineTeamInterview:

				var allocationTime time.Time
				var smsTemplate pkg.SMSTemplateType
				// 为什么golang没有三目运算符orz
				if smsRequest.Next == string(pkg.OnlineGroupInterview) {
					allocationTime = application.InterviewAllocationsGroup
					smsTemplate = pkg.OnLineGroupInterviewSMS
				} else if smsRequest.Next == string(pkg.OnlineTeamInterview) {
					allocationTime = application.InterviewAllocationsTeam
					smsTemplate = pkg.OnLineTeamInterviewSMS
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
					TemplateID: pkg.SMSTemplateMap[smsTemplate],
					Params:     []string{userInfo.Name, recruitmentName, application.Group, formatTime, smsRequest.MeetingId},
				}
				return &smsBody, nil

			//笔试
			case pkg.WrittenTest:
				fallthrough
			//熬测
			case pkg.StressTest:
				if smsRequest.Place == "" {
					return nil, errors.New("place is not provided for " + userInfo.Name)
				}
				if smsRequest.Time == "" {
					return nil, errors.New("time is not provided for " + userInfo.Name)
				}

				defaultRest = fmt.Sprintf("，请于%s在%s参加%s，请务必准时到场",
					smsRequest.Time, smsRequest.Place, pkg.EnToZhStepMap[smsRequest.Next])

			//通过
			case pkg.Pass:
				defaultRest = fmt.Sprintf("，你已成功加入%s组", application.Group)

			//组面时间选择
			case pkg.GroupTimeSelection:
				fallthrough
			//群面时间选择
			case pkg.TeamTimeSelection:

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
				TemplateID: pkg.SMSTemplateMap[pkg.PassSMS],
				Params:     []string{userInfo.Name, recruitmentName, application.Group, pkg.EnToZhStepMap[smsRequest.Current], smsResMessage},
			}
			return &smsBody, nil
		}
	case pkg.Reject:
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
			TemplateID: pkg.SMSTemplateMap[pkg.Delay],
			Params:     []string{userInfo.Name, recruitmentName, application.Group, pkg.EnToZhStepMap[smsRequest.Current], smsResMessage},
		}
		return &smsBody, nil
	}
	return nil, errors.New("sms step is invalid")
}
