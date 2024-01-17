package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"fmt"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateApplication create an application. Remember to submit data with form instead of json!!!
// POST applications/
func CreateApplication(c *gin.Context) {
	var req pkg.CreateAppOpts
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	if constants.GroupMap[req.Group] == "" {
		common.Error(c, rerror.RequestBodyError.WithDetail("group wrong"))
		return
	}

	recruitment, err := models.GetRecruitmentById(req.RecruitmentID)
	if err != nil || recruitment.Uid == "" {
		common.Error(c, rerror.GetDatabaseError.WithData("recruitment").
			WithDetail("when submit the application"))
		return
	}

	// Compare the recruitment time with application time
	if !checkRecruitmentInBtoD(c, recruitment, time.Now()) {
		return
	}

	uid := common.GetUID(c)
	filePath := ""
	if req.Resume != nil {
		// file path example: 2023秋(rname)/web(group)/wwb(uid)/filename
		filePath = fmt.Sprintf("%s/%s/%s/%s", recruitment.Name, req.Group, uid, req.Resume.Filename)

		//resume upload to COS
		err = upLoadAndSaveFileToCos(req.Resume, filePath)
		if err != nil {
			common.Error(c, rerror.UpLoadFileError.WithData(uid).WithDetail(err.Error()))
			return
		}
		zapx.Info("upload resume to tos", zap.String("filepath", filePath))
	}

	//save application to database
	err = models.CreateApplication(&req, uid, filePath)
	if err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData("recruitment").
			WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success save application", nil)
}

// GetApplicationById get candidate's application by applicationId
// GET applications/:aid
// candidate and member will see two different views of application
func GetApplicationById(c *gin.Context) {
	aid := c.Param("aid")
	uid := common.GetUID(c)
	var application *pkg.Application
	var err error

	if common.IsCandidate(c) {
		application, err = models.GetApplicationByIdForCandidate(aid)
	} else {
		application, err = models.GetApplicationById(aid)
	}

	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").
			WithDetail("Get application info fail"))
		return
	}

	userDetail, err := grpc.GetUserInfoByUID(uid)
	if err != nil {
		common.Error(c, rerror.SSOError.WithDetail("when get application"))
	}

	application.UserDetail = userDetail
	common.Success(c, "Get application success", application)
}

// UpdateApplicationById update candidate's application by applicationId
// PUT applications/:aid
// only by application's candidate
func UpdateApplicationById(c *gin.Context) {
	aid := c.Param("aid")
	var req pkg.UpdateAppOpts
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	recruitment, err := models.GetRecruitmentById(req.RecruitmentID)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("recruitment").
			WithDetail("when update the application"))
		return
	}
	// Compare the new recruitment time with application time
	if !checkRecruitmentInBtoD(c, recruitment, time.Now()) {
		return
	}

	uid := common.GetUID(c)

	application, err := models.GetApplicationById(aid)
	if err != nil || application.CandidateID != uid {
		common.Error(c, rerror.UpdateDatabaseError.WithData("application").
			WithDetail("you can't update other's application"))
		return
	}

	filePath := ""
	if req.Resume != nil {
		filePath = fmt.Sprintf("%s/%s/%s/%s", recruitment.Name, req.Group, uid, req.Resume.Filename)
		if err := upLoadAndSaveFileToCos(req.Resume, filePath); err != nil {
			common.Error(c, rerror.UpLoadFileError.WithData(uid).WithDetail(err.Error()))
			return
		}
	}

	if err := models.UpdateApplication(aid, filePath, &req); err != nil {
		common.Error(c, rerror.UpdateDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	common.Success(c, "update application success", nil)
}

// DeleteApplicationById delete candidate's application by applicationId
// DELETE applications/:aid
// only by application's candidate
func DeleteApplicationById(c *gin.Context) {
	aid := c.Param("aid")
	if aid == "" {
		common.Error(c, rerror.RequestBodyError.WithDetail("lost aid param"))
		return
	}

	uid := common.GetUID(c)
	application, err := models.GetApplicationById(aid)
	if err != nil || application.CandidateID != uid {
		common.Error(c, rerror.UpdateDatabaseError.WithData("application").WithDetail("you can't delete other's application"))
		return
	}

	if err := models.DeleteApplication(aid); err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData("application"))
		return
	}
	common.Success(c, "delete application success", nil)
}

// AbandonApplicationById abandon candidate's application by applicationId
// DELETE applications/:aid/abandoned
// only by the member of application's group
func AbandonApplicationById(c *gin.Context) {
	aid := c.Param("aid")
	uid := common.GetUID(c)

	// check member's role to abandon application
	if !checkMemberGroup(c, aid, uid) {
		return
	}
	if err := models.AbandonApplication(aid); err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData("application"))
		return
	}
	common.Success(c, "abandon application success", nil)
}

// GetResumeById Download resume by application's
// GET applications/:aid/resume
// TODO(wwb)
// check this api
func GetResumeById(c *gin.Context) {
	aid := c.Param("aid")
	application, err := models.GetApplicationById(aid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail("Get application info fail"))
		return
	}
	resp, err := utils.GetCOSObjectResp(application.Resume)
	if err != nil {
		common.Error(c, rerror.DownloadFileError.WithData("application").WithDetail("download resume fail"))
		return
	}

	reader := resp.Body
	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)

}

// GetApplicationByRecruitmentId get all applications by recruitmentId
// GET applications/recruitment/:rid
// member role
func GetApplicationByRecruitmentId(c *gin.Context) {
	rid := c.Param("rid")
	applications, err := models.GetApplicationByRecruitmentId(rid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail("Get application info fail"))
		return
	}
	common.Success(c, "get applications success", applications)
}

// PUT applications/:aid/step
// only by the member of application's group
func SetApplicationStepById(c *gin.Context) {
	aid := c.Param("aid")
	var req pkg.SetAppStepOpts
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}
	uid := common.GetUID(c)
	// check member's role to set application step
	if !checkMemberGroup(c, aid, uid) {
		return
	}
	if err := models.SetApplicationStepById(aid, &req); err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData(err.Error()))
		return
	}
	common.Success(c, "set application step success", nil)

}

// SetApplicationInterviewTimeById allocate application's group/team interview time
// PUT /:aid/interview/:type
// by the member of application's group
func SetApplicationInterviewTimeById(c *gin.Context) {
	aid := c.Param("aid")
	interviewType := c.Param("type")
	if interviewType != "group" && interviewType != "team" {
		common.Error(c, rerror.RequestParamError.WithDetail("type wrong"))
		return
	}
	var req pkg.SetAppInterviewTimeOpts
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	// check application's status such as abandoned
	application, err := models.GetApplicationById(aid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	if !checkApplyStatus(c, application) {
		return
	}

	// check member's role to set application interview time
	uid := common.GetUID(c)
	// if common.IsCandidate(c) && !checkIsApplicationOwner(c, uid, application) {
	// 	return
	// }
	if common.IsMember(c) && !checkMemberGroup(c, aid, uid) {
		return
	}

	// check update application time is between the start and the end
	recruitment, err := models.GetRecruitmentById(application.RecruitmentID)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	if !checkRecruitmentTimeInBtoE(c, recruitment) {
		return
	}

	if err := models.SetApplicationInterviewTime(aid, interviewType, req.Time); err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData(err.Error()))
		return
	}
	common.Success(c, "set interview time success", nil)
}

// GetInterviewsSlots get the interviews times candidates can select
// Follow the old HR code, this api will get all the interviews assigned by this group's member
// I think this api should get the interviews times candidate selected
// And the interviews selected by candidate can be get by GetApplicationById
// GET /:aid/slots/:type
// candidate / member role
func GetInterviewsSlots(c *gin.Context) {
	aid := c.Param("aid")
	interviewType := c.Param("type")
	application, err := models.GetApplicationByIdForCandidate(aid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}

	// check application's step such as group/team when get interview time
	// 	if !checkStep(c, application.Step, constants.GroupTimeSelection) {
	// 		return
	// 	}

	recruitment, err := models.GetFullRecruitmentById(application.RecruitmentID)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}

	// TODO(tmy) type
	var name string
	if interviewType == "group" {
		name = application.Group
	} else {
		name = "unique"
	}

	var res []pkg.Interview
	for _, interview := range recruitment.Interviews {
		if string(interview.Name) == name {
			res = append(res, interview)
		}
	}
	common.Success(c, "Success get interview time", res)

}

// SelectInterviewSlots select group/team interview time for candidate
// to save time, this api will not check Whether slotnum exceeds the limit
// PUT /:aid/slots/:type
// candidate role
func SelectInterviewSlots(c *gin.Context) {
	aid := c.Param("aid")
	interviewType := c.Param("type")

	var req struct {
		Iids []string `json:"iids"`
	}
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithData("application").WithDetail(err.Error()))
		return
	}

	application, err := models.GetApplicationById(aid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	recruitmentById, err := models.GetRecruitmentById(application.RecruitmentID)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}

	// check if user is the application's owner
	if !checkIsApplicationOwner(c, common.GetUID(c), application) {
		return
	}

	if !checkApplyStatus(c, application) {
		return
	}

	if !checkRecruitmentTimeInBtoE(c, recruitmentById) {
		return
	}

	if !checkStep(c, interviewType, application) {
		return
	}

	var name constants.Group
	if interviewType == string(constants.InGroup) {
		name = constants.GroupMap[application.Group]
	} else {
		name = "unique"
	}

	// 这啥意思？？？？?
	// for _, interview := range application.InterviewSelections {
	// 	if interview.Name != name {
	// 		common.Error(c, rerror.ReselectInterviewError.WithData("application"))
	// 		return
	// 	}
	// }

	var errors []string

	var interviews []*pkg.Interview
	for _, iid := range req.Iids {
		// check the select interview is in the recruitment
		interview, err := models.GetInterviewById(iid)
		if err != nil {
			errors = append(errors, rerror.GetDatabaseError.WithData("interview").Msg()+err.Error())
			continue
		}
		// check the select interview name == param name
		if interview.Name != name {
			errors = append(errors, rerror.CheckPermissionError.Msg()+"the select interview name != group/team name")
			continue
		}
		interviews = append(interviews, interview)
	}

	if err = models.UpdateInterviewSelection(application, interviews); err != nil {
		errors = append(errors, rerror.SaveDatabaseError.WithData("application").Msg()+err.Error())
	}
	if len(errors) != 0 {
		common.Error(c, rerror.SaveDatabaseError.WithData("application").WithDetail(errors...))
		return
	}
	common.Success(c, "Success select interview time", nil)
}

// MoveApplication move the step of application by member
// PUT applications/:aid/step
// member role
func MoveApplication(c *gin.Context) {
	req := struct {
		From string `form:"from" json:"from,omitempty"`
		To   string `form:"to" json:"to,omitempty"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}
	applicationId := c.Param("aid")
	if applicationId == "" {
		common.Error(c, rerror.RequestBodyError.WithDetail("lost aid param"))
		return
	}
	application, err := models.GetApplicationById(applicationId)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithDetail("failed to get application for member"))
		return
	}
	recruitment, err := models.GetRecruitmentById(application.RecruitmentID)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("recruitment").WithDetail("when you move application"))
		return
	}
	//check application's status
	if !checkApplyStatus(c, application) {
		return
	}
	if !checkRecruitmentTimeInBtoE(c, recruitment) {
		return
	}
	// TODO(wwb)
	// Add check member's group
	//if b := checkMemberGroup(c,application);b
	if application.Step != req.From {
		common.Error(c, rerror.RequestBodyError.WithDetail("application's step != request's from"))
		return
	}
	if err := models.UpdateApplicationStep(applicationId, req.To); err != nil {
		common.Error(c, rerror.UpdateDatabaseError.WithData("application").WithDetail("when you update application's step"))
		return
	}
	common.Success(c, "Update application step success", nil)
}

// SetApplicationInterviewTime set interview time
// PUT /interview/:type
// only by the member of application's group
func SetApplicationInterviewTime(c *gin.Context) {
	aid := c.Param("aid")
	interviewType := c.Param("type")
	var req struct {
		Time time.Time `json:"time"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithData("application").WithDetail(err.Error()))
		return
	}
	application, err := models.GetApplicationById(aid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	if !checkApplyStatus(c, application) {
		return
	}

	recruitment, err := models.GetRecruitmentById(application.RecruitmentID)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	if !checkRecruitmentTimeInBtoE(c, recruitment) {
		return
	}

	// if !checkStep(c, interviewType) {
	// 	return
	// }

	switch interviewType {
	case "team":
		application.InterviewAllocationsTeam = req.Time
	case "group":
		application.InterviewAllocationsGroup = req.Time
	}

	if err := models.UpdateApplicationInfo(application); err != nil {
		common.Error(c, rerror.SaveDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}

	common.Success(c, "Success set interview time", nil)

}

// checkRecruitmentInBtoD check whether the recruitment is between the start and the deadline
// such as summit the application/update the application
func checkRecruitmentInBtoD(c *gin.Context, recruitment *pkg.Recruitment, now time.Time) bool {
	if recruitment.Beginning.After(now) {
		// submit too early
		common.Error(c, rerror.RecruitmentNotReady.WithData(recruitment.Name))
		return false
	} else if recruitment.Deadline.Before(now) {
		common.Error(c, rerror.RecruitmentStopped.WithData(recruitment.Name))
		return false
	} else if recruitment.End.Before(now) {
		common.Error(c, rerror.RecruitmentEnd.WithData(recruitment.Name))
		return false
	}
	return true
}

// checkRecruitmentInBtoE check whether the recruitment is between the start and the end
// such as move the application's step
func checkRecruitmentTimeInBtoE(c *gin.Context, recruitment *pkg.Recruitment) bool {
	now := time.Now()
	if recruitment.Beginning.After(now) {
		common.Error(c, rerror.RecruitmentNotReady.WithData(recruitment.Name))
		return false
	} else if recruitment.End.Before(now) {
		common.Error(c, rerror.RecruitmentEnd.WithData(recruitment.Name))
		return false
	}
	return true
}

// check application's status
// If the resume has already been rejected or abandoned return false
func checkApplyStatus(c *gin.Context, application *pkg.Application) bool {
	if application.Rejected {
		common.Error(c, rerror.Rejected.WithData(application.CandidateID))
		return false
	}
	if application.Abandoned {
		common.Error(c, rerror.Abandoned.WithData(application.CandidateID))
		return false
	}
	return true
}

// check if application step is in interview select status
func checkStep(c *gin.Context, interviewType string, application *pkg.Application) bool {
	if interviewType == "group" && application.Step != string(constants.GroupTimeSelection) {
		common.Error(c, rerror.CheckPermissionError.WithDetail("you can't set group interview time"))
		return false
	}
	if interviewType == "team" && application.Step != string(constants.TeamTimeSelection) {
		common.Error(c, rerror.CheckPermissionError.WithDetail("you can't set team interview time"))
		return false
	}
	return true
}

// check if the user is a member of group the application applied
func checkMemberGroup(c *gin.Context, aid string, uid string) bool {
	application, err := models.GetApplicationByIdForCandidate(aid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return false
	}

	userInfo, err := grpc.GetUserInfoByUID(uid)
	if err != nil {
		common.Error(c, rerror.CheckPermissionError.WithDetail(err.Error()))
		return false
	}
	if utils.CheckInGroups(userInfo.Groups, application.Group) {
		return true
	}
	common.Error(c, rerror.GroupNotMatch)
	return false
}

func checkIsApplicationOwner(c *gin.Context, uid string, application *pkg.Application) bool {
	if application.CandidateID == uid {
		return true
	}
	common.Error(c, rerror.CheckPermissionError.WithDetail("you are not the owner of this application"))
	return false
}
