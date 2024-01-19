package controllers

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateApplication create an application. Remember to submit data with form instead of json!!!
// POST applications/
func CreateApplication(c *gin.Context) {
	var (
		app *pkg.Application
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, app, err) }()

	opts := &pkg.CreateAppOpts{}
	if err = c.ShouldBind(&opts); err != nil {
		return
	}

	if err = opts.Validate(); err != nil {
		return
	}

	r, err = models.GetRecruitmentById(opts.RecruitmentID)
	if err != nil {
		return
	}

	// Compare the recruitment time with application time
	if err = checkRecruitmentInBtoD(r, time.Now()); err != nil {
		return
	}

	uid := common.GetUID(c)
	filePath := ""
	if opts.Resume != nil {
		// file path example: 2023秋(rname)/web(group)/wwb(uid)/filename
		filePath = fmt.Sprintf("%s/%s/%s/%s", r.Name, opts.Group, uid, opts.Resume.Filename)
	}

	//save application to database
	app, err = models.CreateApplication(opts, uid, filePath)
	return
}

// GetApplication get candidate's application by applicationId
// GET applications/:aid
// candidate and member will see two different views of application
func GetApplication(c *gin.Context) {
	var (
		app  *pkg.Application
		user *pkg.UserDetail
		err  error
	)
	defer func() { common.Resp(c, app, err) }()

	aid := c.Param("aid")
	uid := common.GetUID(c)
	if aid == "" {
		err = errors.New("request param error, application id is nil")
		return
	}

	if common.IsCandidate(c) {
		app, err = models.GetApplicationByIdForCandidate(aid)
		if app.CandidateID != uid {
			err = errors.New("for candidate,you can't see other's application")
			return
		}
	} else {
		app, err = models.GetApplicationById(aid)
	}

	if err != nil {
		return
	}

	user, err = grpc.GetUserInfoByUID(uid)
	if err != nil {
		return
	}

	app.UserDetail = user
	return
}

// UpdateApplication update candidate's application by applicationId
// PUT applications/:aid
// only by application's candidate
func UpdateApplication(c *gin.Context) {
	var (
		app *pkg.Application
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, app, err) }()

	opts := &pkg.UpdateAppOpts{}
	opts.Aid = c.Param("aid")
	uid := common.GetUID(c)

	if err = c.ShouldBind(opts); err != nil {
		return
	}
	if err = opts.Validate(); err != nil {
		return
	}

	app, err = models.GetApplicationByIdForCandidate(opts.Aid)
	if err != nil {
		return
	}
	if app.Abandoned || app.Rejected {
		err = fmt.Errorf("you have been abandoned / rejected")
		return
	}

	r, err = models.GetRecruitmentById(app.RecruitmentID)
	if err != nil {
		return
	}

	// Compare the new recruitment time with application time
	if err = checkRecruitmentInBtoD(r, time.Now()); err != nil {
		return
	}

	// can't update other's application
	if app.CandidateID != uid {
		err = errors.New("you can't update other's application")
		return
	}

	filePath := ""
	if opts.Resume != nil {
		filePath = fmt.Sprintf("%s/%s/%s/%s", r.Name, opts.Group, uid, opts.Resume.Filename)
	}

	app, err = models.UpdateApplication(opts, filePath)
	return
}

// DeleteApplication delete candidate's application by applicationId
// DELETE applications/:aid
// only by application's candidate
func DeleteApplication(c *gin.Context) {
	var (
		app *pkg.Application
		err error
	)
	defer func() { common.Resp(c, app, err) }()

	aid := c.Param("aid")
	uid := common.GetUID(c)
	if aid == "" {
		err = fmt.Errorf("request body error, application id is nil")
		return
	}

	app, err = models.GetApplicationByIdForCandidate(aid)
	if err != nil {
		return
	}

	// can't delete other's application
	if app.CandidateID != uid {
		err = errors.New("you can't delete other's application")
		return
	}
	err = models.DeleteApplication(aid)
	return
}

// AbandonApplication abandon candidate's application by applicationId
// DELETE applications/:aid/abandoned
// only by the member of application's group
func AbandonApplication(c *gin.Context) {
	var (
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	aid := c.Param("aid")
	if aid == "" {
		err = fmt.Errorf("request param error, application id is nil")
		return
	}

	uid := common.GetUID(c)

	// check member's role to abandon application
	if err = checkMemberGroup(aid, uid); err != nil {
		return
	}

	err = models.AbandonApplication(aid)
	return
}

// GetResume Download resume by application's
// GET applications/:aid/resume
func GetResume(c *gin.Context) {
	var (
		app *pkg.Application
		err error
	)

	aid := c.Param("aid")
	if aid == "" {
		err = fmt.Errorf("request param error, application id is nil")
		return
	}

	app, err = models.GetApplicationByIdForCandidate(aid)
	if err != nil {
		common.Resp(c, nil, err)
		return
	}

	// don't have role to download file
	if !common.IsMember(c) && !(app.CandidateID == common.GetUID(c)) {
		err = fmt.Errorf("you don't have role to download file")
		common.Resp(c, nil, err)
		return
	}
	if app.Resume == "" {
		err = fmt.Errorf("you don't upload resume")
		common.Resp(c, nil, err)
		return
	}

	resp, err := global.GetCOSObjectResp(app.Resume)
	if err != nil {
		common.Resp(c, nil, err)
		return
	}

	reader := resp.Body
	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
}

// GetAllApplications get all applications by recruitmentId
// GET applications/recruitment/:rid
// member role
func GetAllApplications(c *gin.Context) {
	var (
		apps []pkg.Application
		err  error
	)
	defer func() { common.Resp(c, apps, err) }()

	rid := c.Param("rid")
	if rid == "" {
		err = fmt.Errorf("request body error, recruitment id is nil")
		return
	}

	apps, err = models.GetApplicationsByRid(rid)
	if err != nil {
		return
	}

	if len(apps) == 0 {
		return
	}

	// todo wwb
	// add grpc handler(get all user details)
	//var userIds []string
	//for _, app := range apps {
	//	userIds = append(userIds, app.CandidateID)
	//}
	for i, _ := range apps {
		apps[i].UserDetail, err = grpc.GetUserInfoByUID(apps[i].CandidateID)
		if err != nil {
			return
		}
	}
	return
}

// PUT applications/:aid/step
// only by the member of application's group
func SetApplicationStep(c *gin.Context) {
	var (
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.SetAppStepOpts{}
	opts.Aid = c.Param("aid")

	if err = c.ShouldBind(&opts); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	uid := common.GetUID(c)
	// check member's role to set application step
	if err = checkMemberGroup(opts.Aid, uid); err != nil {
		return
	}

	err = models.SetApplicationStepById(opts)
	return
}

// SetApplicationInterviewTimeById allocate application's group/team interview time
// PUT /:aid/interview/:type
// by the member of application's group
func SetApplicationInterviewTimeById(c *gin.Context) {
	var (
		app *pkg.Application
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.SetAppInterviewTimeOpts{}
	if err = c.ShouldBind(&opts); err != nil {
		return
	}

	opts.Aid = c.Param("aid")
	opts.InterviewType = c.Param("type")
	if err = opts.Validate(); err != nil {
		return
	}

	// check application's status such as abandoned
	app, err = models.GetApplicationById(opts.Aid)
	if err != nil {
		return
	}
	if err = checkApplyStatus(app); err != nil {
		return
	}

	// check member's role to set application interview time
	uid := common.GetUID(c)
	if err = checkMemberGroup(opts.Aid, uid); err != nil {
		return
	}

	// check update application time is between the start and the end
	r, err = models.GetRecruitmentById(app.RecruitmentID)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	if err = checkRecruitmentTimeInBtoE(r); err != nil {
		return
	}

	err = models.SetApplicationInterviewTime(opts)
	return
}

// GetInterviewsSlots get the interviews times candidates can select
// Follow the old HR code, this api will get all the interviews assigned by this group's member
// I think this api should get the interviews times candidate selected
// And the interviews selected by candidate can be got by GetApplicationById
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

// SelectInterviewSlots candidate select group/team interview time
// to save time, this api will not check Whether slotnum exceeds the limit
// PUT /:aid/slots/:type
// candidate role
func SelectInterviewSlots(c *gin.Context) {
	var (
		app *pkg.Application
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.SelectInterviewSlotsOpts{}
	if err = c.ShouldBind(&opts); err != nil {
		return
	}
	opts.Aid = c.Param("aid")
	opts.InterviewType = c.Param("type")
	if err = opts.Validate(); err != nil {
		return
	}

	app, err = models.GetApplicationById(opts.Aid)
	if err != nil {
		return
	}

	r, err = models.GetRecruitmentById(app.RecruitmentID)
	if err != nil {
		return
	}

	uid := common.GetUID(c)
	// check if user is the application's owner
	if app.CandidateID != uid {
		err = errors.New("you can't update other's application")
		return
	}

	if err = checkApplyStatus(app); err != nil {
		return
	}

	if err = checkRecruitmentTimeInBtoE(r); err != nil {
		return
	}

	if err = checkStepInInterviewSelectStatus(opts.InterviewType, app); err != nil {
		return
	}

	var name pkg.Group
	if opts.InterviewType == string(constants.InGroup) {
		name = pkg.GroupMap[app.Group]
	} else {
		name = "unique"
	}

	// ？？？？?
	// for _, interview := range application.InterviewSelections {
	// 	if interview.Name != name {
	// 		common.Error(c, rerror.ReselectInterviewError.WithData("application"))
	// 		return
	// 	}
	// }

	var ierrors []string
	var interviews []*pkg.Interview

	for _, iid := range opts.Iids {
		interview := &pkg.Interview{}
		var ierr error
		// check the select interview is in the recruitment
		interview, ierr = models.GetInterviewById(iid)
		if err != nil {
			ierrors = append(ierrors, fmt.Sprintf("[get interview %s in db failed, %s] ", interview.Uid, ierr.Error()))
			continue
		}
		// check the select interview name == param name
		if interview.Name != name {
			ierrors = append(ierrors,
				fmt.Sprintf("[the select interview %s name = %s, group/team name = %s , failed]", interview.Uid, interview.Name, name))
			continue
		}
		interviews = append(interviews, interview)
	}

	if updateErr := models.UpdateInterviewSelection(app, interviews); updateErr != nil {
		ierrors = append(ierrors, fmt.Sprintf("[%s]", updateErr.Error()))
	}
	if len(ierrors) != 0 {
		err = fmt.Errorf("there are %d error msg: %v", len(ierrors), ierrors)
		return
	}
	return
}

// checkRecruitmentInBtoD check whether the recruitment is between the start and the deadline
// such as summit the application/update the application
func checkRecruitmentInBtoD(r *pkg.Recruitment, now time.Time) error {
	if r.Beginning.After(now) {
		// submit too early
		return fmt.Errorf("recruitment %s has not started yet", r.Name)
	} else if r.Deadline.Before(now) {
		return fmt.Errorf("the application deadline of recruitment %s has already passed", r.Name)
	} else if r.End.Before(now) {
		return fmt.Errorf("recruitment %s has already ended", r.Name)
	}
	return nil
}

// checkRecruitmentInBtoE check whether the recruitment is between the start and the end
// such as move the application's step
func checkRecruitmentTimeInBtoE(recruitment *pkg.Recruitment) error {
	now := time.Now()
	if recruitment.Beginning.After(now) {
		return fmt.Errorf("recruitment %s has not started yet", recruitment.Name)
	} else if recruitment.End.Before(now) {
		return fmt.Errorf("recruitment %s has already ended", recruitment.Name)
	}
	return nil
}

// check application's status
// If the resume has already been rejected or abandoned return false
func checkApplyStatus(application *pkg.Application) error {
	if application.Rejected {
		return fmt.Errorf("application %s has already been rejected", application.Uid)
	}
	if application.Abandoned {
		return fmt.Errorf("application %s has already been abandoned ", application.Uid)
	}
	return nil
}

// check if application step is in interview select status
func checkStepInInterviewSelectStatus(interviewType string, app *pkg.Application) error {
	if interviewType == "group" && app.Step != string(constants.GroupTimeSelection) {
		return fmt.Errorf("you can't set group interview time now")
	}
	if interviewType == "team" && app.Step != string(constants.TeamTimeSelection) {
		return fmt.Errorf("you can't set team interview time now")
	}
	return nil
}

// check if the user is a member of group the application applied
func checkMemberGroup(aid string, uid string) (err error) {
	appToCheck, err := models.GetApplicationByIdForCandidate(aid)
	if err != nil {
		return err
	}

	member, err := grpc.GetUserInfoByUID(uid)
	if err != nil {
		return err
	}

	if utils.CheckInGroups(member.Groups, appToCheck.Group) {
		return nil
	}

	return errors.New("you and the candidate are not in the same group, " +
		"and you cannot manipulate other people’s application. ")
}
