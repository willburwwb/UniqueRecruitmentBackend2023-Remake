package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xylonx/zapx"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/sms"
)

// SendSMS send sms to user.
// @Id send_sms
// @Summary Send sms
// @Description Send sms to user, include Accept, Reject, detailed information reference https://uniquestudio.feishu.cn/docx/Yh96d2DoyoCe6zxlR0ecSU5snDd?from=from_copylink
// @Tags Sms
// @Accept  json
// @Produce json
// @Param pkg.SendSMSOpts body pkg.SendSMSOpts true "sms body params"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /sms [Post]
func SendSMS(c *gin.Context) {
	var (
		app  *pkg.Application
		r    *pkg.Recruitment
		user *pkg.UserDetail
		err  error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.SendSMSOpts{}
	if err = c.ShouldBind(&opts); err != nil {
		return
	}

	if err = opts.Validate(); err != nil {
		return
	}

	app, err = models.GetApplicationByIdForCandidate(opts.Aids[0])
	if err != nil {
		return
	}

	// judge whether the recruitment has expired
	r, err = models.GetFullRecruitmentById(app.RecruitmentID)
	if err != nil {
		return
	}
	if r.End.Before(time.Now()) {
		err = fmt.Errorf("recruitment %s has already ended", r.Name)
		return
	}

	var errors []string
	for _, aid := range opts.Aids {
		app, err = models.GetApplicationByIdForCandidate(aid)
		if err != nil {
			errors = append(errors, fmt.Sprintf("get application %s failed, error: %s", aid, err.Error()))
			continue
		}

		user, err = grpc.GetUserInfoByUID(app.CandidateID)
		if err != nil {
			return
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
			if opts.Next == pkg.GroupInterview && len(r.GetInterviews(app.Group)) == 0 {
				errors = append(errors, fmt.Sprintf("no interviews are scheduled for %s", app.Group))
				continue
			}
			if opts.Next == pkg.TeamInterview && len(r.GetInterviews("unique")) == 0 {
				errors = append(errors, fmt.Sprintf("no interviews are scheduled for unique"))
				continue
			}
		} else if opts.Type == pkg.Reject {
			app.Rejected = true
			// save application
			if err = models.UpdateApplicationInfo(app); err != nil {
				errors = append(errors, fmt.Sprintf("update application %s failed, error: %s", aid, err.Error()))
				continue
			}
		}

		var smsBody *sms.SMSBody
		smsBody, err = ApplySMSTemplate(opts, user, app, r)
		if err != nil {
			errors = append(errors, fmt.Sprintf("set smsbody for user %s failed, error: %s", user.Name, err.Error()))
			continue
		}

		// send sms to candidate
		smsBody.Phone = user.Phone
		zapx.Infof("smsbody : %v", *smsBody)
		if _, err = sms.SendSMS(*smsBody); err != nil {
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
			switch smsRequest.Next {
			//組面
			case pkg.GroupInterview:
				fallthrough
			//群面
			case pkg.TeamInterview:
				var allocationTime time.Time
				if smsRequest.Next == pkg.GroupInterview {
					allocationTime = application.InterviewAllocationsGroup.Start
				} else if smsRequest.Next == pkg.TeamInterview {
					allocationTime = application.InterviewAllocationsTeam.Start
				}

				if smsRequest.Place == "" {
					return nil, errors.New("Place is not provided for " + userInfo.Name)
				}
				if allocationTime.IsZero() {
					return nil, errors.New("Interview time is not allocated for " + userInfo.Name)
				}

				// set interview time format
				// interview time get from application instead of smsRequest
				// 2006年1月2日 星期一 15时04分05秒
				formatTime := utils.ConverToLocationTime(allocationTime)
				// {1}你好，请于{2}在启明学院亮胜楼{3}参加{4}，请准时到场。
				smsBody = sms.SMSBody{
					TemplateID: pkg.SMSTemplateMap[pkg.Interviews],
					Params:     []string{userInfo.Name, formatTime, smsRequest.Place, pkg.EnToZhStepMap[smsRequest.Next]},
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
				if smsRequest.Next == pkg.OnlineGroupInterview {
					allocationTime = application.InterviewAllocationsGroup.Start
					smsTemplate = pkg.OnLineGroupInterviewSMS
				} else if smsRequest.Next == pkg.OnlineTeamInterview {
					allocationTime = application.InterviewAllocationsTeam.Start
					smsTemplate = pkg.OnLineTeamInterviewSMS
				}
				if allocationTime.IsZero() {
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
					Params:     []string{userInfo.Name, recruitmentName, string(application.Group), formatTime, smsRequest.MeetingId},
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
				Params:     []string{userInfo.Name, recruitmentName, string(application.Group), pkg.EnToZhStepMap[smsRequest.Current], smsResMessage},
			}
			return &smsBody, nil
		}
	case pkg.Reject:
		defaultRest := "不要灰心，继续学习。期待与更强大的你的相遇！"
		var smsResMessage string
		if smsRequest.Rest == "" {
			smsResMessage = defaultRest + suffix
		} else {
			smsResMessage = smsRequest.Rest + suffix
		}
		// {1}你好，你没有通过{2}{3}组{4}审核，请你{5}
		smsBody = sms.SMSBody{
			TemplateID: pkg.SMSTemplateMap[pkg.Delay],
			Params:     []string{userInfo.Name, recruitmentName, string(application.Group), pkg.EnToZhStepMap[smsRequest.Current], smsResMessage},
		}
		return &smsBody, nil
	}
	return nil, errors.New("sms step is invalid")
}

// SendCode send code to admin
// @Id send_code
// @Summary Send code
// @Description Send code to admin
// @Tags Sms
// @Accept  json
// @Produce json
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /sms/code [post]
func SendCode(c *gin.Context) {
	var (
		user *pkg.UserDetail
		err  error
	)
	defer func() { common.Resp(c, nil, err) }()

	uid := common.GetUID(c)
	if user, err = grpc.GetUserInfoByUID(uid); err != nil {
		return
	}

	var smsCode string
	smsCode, err = utils.GenerateTmpCode(c, user.Phone, 5*time.Minute)
	if err != nil {
		return
	}

	smsBody := sms.SMSBody{
		TemplateID: pkg.SMSTemplateMap[pkg.VerificationCode],
		Phone:      user.Phone,
		Params:     []string{smsCode},
	}

	_, err = sms.SendSMS(smsBody)
	return
}
