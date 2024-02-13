package controllers

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"errors"
	"fmt"
	"github.com/xylonx/zapx"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateApplication create application.
// @Id create_application.
// @Summary create an application for candidate.
// @Description create an application. Remember to submit data with form instead of json!!!
// @Tags application
// @Accept  multipart/form-data
// @Produce  json
// @Param pkg.CreateAppOpts body pkg.CreateAppOpts true "application detail"
// @Success 200 {object} common.JSONResult{data=pkg.Application} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications [post]
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
		// file path example: 2023秋(rname)/wwb(uid)/filename
		filePath = fmt.Sprintf("%s/%s/%s", r.Name, uid, opts.Resume.Filename)
	}

	//save application to database
	app, err = models.CreateApplication(opts, uid, filePath)
	return
}

// GetApplication get application.
// @Id get_application.
// @Summary get an application for candidate and member
// @Description get candidate's application by applicationId, candidate and member will see different views of application
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Success 200 {object} common.JSONResult{data=pkg.Application} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid} [get]
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

// UpdateApplication update application.
// @Id update_application.
// @Summary update candidate's application by applicationId
// @Description update candidate's application by applicationId, can only be modified by application's owner
// @Tags application
// @Accept  multipart/form-data
// @Produce  json
// @Param aid path string true "application id"
// @Param pkg.UpdateAppOpts body pkg.UpdateAppOpts true "update application opts"
// @Success 200 {object} common.JSONResult{data=pkg.Application} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid} [put]
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
		filePath = fmt.Sprintf("%s/%s/%s", r.Name, uid, opts.Resume.Filename)
	}

	app, err = models.UpdateApplication(opts, filePath)
	return
}

// DeleteApplication delete application.(DEPRECATED! Use abandon instead.)
// @Id delete_application.
// @Summary delete candidate's application by applicationId
// @Description delete candidate's application by applicationId, can only be deleted by application's owner
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Success 200 {object} common.JSONResult{data=pkg.Application} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid} [delete]
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

// AbandonApplication abandon application.
// @Id abandon_application.
// @Summary candidate abandon his/her application
// @Description candidate abandon his/her application, can only be abandoned by application's owner
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid}/abandoned [put]
func AbandonApplication(c *gin.Context) {
	var (
		app *pkg.Application
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	aid := c.Param("aid")
	if aid == "" {
		err = fmt.Errorf("request param error, application id is nil")
		return
	}

	uid := common.GetUID(c)
	app, err = models.GetApplicationByIdForCandidate(aid)
	if err != nil {
		return
	}

	if app.CandidateID != uid {
		err = errors.New("you can't abandon other's application")
		return
	}

	err = models.AbandonApplication(aid)
	return
}

// RejectApplication reject application.
// @Id reject_application.
// @Summary reject candidate's application by applicationId,
// @Description reject candidate's application by applicationId, can only be abandoned by member of the corresponding group
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid}/abandoned [put]
func RejectApplication(c *gin.Context) {
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

	err = models.RejectApplication(aid)
	return
}

// GetResume get application's resume.
// @Id get_resume.
// @Summary get application's resume by applicationId
// @Description get application's resume by applicationId, can only be got by member or application's owner
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid}/resume [get]
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

// GetAllApplications get all applications by recruitmentId.
// @Id get_all_applications.
// @Summary get all applications by recruitmentId.
// @Description get all applications by recruitmentId, can only be got by member, applications information included comments and interview selections.
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Success 200 {object} common.JSONResult{data=[]pkg.Application} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/recruitment/{rid} [get]
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

// SetApplicationStep set application step by applicationId.
// @Id set_application_step.
// @Summary set application step by applicationId.
// @Description get all applications by recruitmentId, can only be modified by member of the corresponding group
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid}/step [put]
func SetApplicationStep(c *gin.Context) {
	var (
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.SetAppStepOpts{}
	opts.Aid = c.Param("aid")

	if err = c.ShouldBind(&opts); err != nil {
		return
	}

	if err = opts.Validate(); err != nil {
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

// SetApplicationInterviewTime set_application_interview_time
// @Id set_application_interview_time.
// @Summary allocate application's group/team interview time.
// @Description allocate application's group/team interview time, can only be modified by member of the corresponding group
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Param	type path pkg.GroupOrTeam true "group or team"
// @Param	interview_id body string true "interview uid"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid}/interview/{type} [put]
func SetApplicationInterviewTime(c *gin.Context) {
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
	opts.InterviewType = pkg.GroupOrTeam(c.Param("type"))
	if err = opts.Validate(); err != nil {
		return
	}

	// check application's status such as abandoned
	app, err = models.GetApplicationByIdForCandidate(opts.Aid)
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
		return
	}
	if err = checkRecruitmentTimeInBtoE(r); err != nil {
		return
	}

	err = models.SetApplicationInterviewTime(opts)
	return
}

// GetInterviewsSlots set_application_interview_time
// @Id set_application_interview_time.
// @Summary allocate application's group/team interview time.
// @Description allocate application's group/team interview time, can only be modified by member of the corresponding group
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Param	type path pkg.GroupOrTeam true "group or team"
// @Success 200 {object} common.JSONResult{data=[]pkg.Interview} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid}/interview/{type} [get]
func GetInterviewsSlots(c *gin.Context) {
	var (
		interviews []pkg.Interview
		app        *pkg.Application
		r          *pkg.Recruitment
		err        error
	)
	defer func() { common.Resp(c, interviews, err) }()

	opts := &pkg.GetInterviewsSlotsOpts{}
	if err = c.ShouldBindUri(opts); err != nil {
		return
	}
	if err = opts.Validate(); err != nil {
		return
	}

	app, err = models.GetApplicationByIdForCandidate(opts.Aid)
	if err != nil {
		return
	}

	r, err = models.GetFullRecruitmentById(app.RecruitmentID)
	if err != nil {
		return
	}

	var name pkg.Group
	if opts.InterviewType == pkg.InGroup {
		name = app.Group
	} else {
		name = pkg.Unique
	}

	for _, interview := range r.Interviews {
		if interview.Name == name {
			interviews = append(interviews, interview)
		}
	}
	return
}

// SelectInterviewSlots select interview slots
// @Id select_interview_slots.
// @Summary candidate select group/team interview time.
// @Description candidate select group/team interview time, to save time, this api will not check Whether slot number exceeds the limit
// @Tags application
// @Accept  json
// @Produce  json
// @Param	aid path int true "application id"
// @Param	type path pkg.GroupOrTeam true "group or team"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /applications/{aid}/slots/{type} [put]
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
	opts.InterviewType = pkg.GroupOrTeam(c.Param("type"))
	if err = opts.Validate(); err != nil {
		return
	}

	app, err = models.GetApplicationByIdForCandidate(opts.Aid)
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
	if opts.InterviewType == pkg.InGroup {
		name = app.Group
	} else {
		name = pkg.Unique
	}

	var interviews []pkg.Interview
	interviews, err = models.GetInterviewsByIds(opts.Iids)
	if err != nil {
		return
	}

	var ierrors []string
	for _, interview := range interviews {
		// check the select interview name == param name
		if interview.Name != name {
			ierrors = append(ierrors,
				fmt.Sprintf("[the select interview %s name = %s, group/team name = %s , failed]", interview.Uid, interview.Name, name))
			continue
		}
		interviews = append(interviews, interview)
	}

	iidsToAdd, iidsToDel := getAddAndDelInterviews(app.InterviewSelections, opts.Iids)
	zapx.Infof("iidsToAdd %v, iidsToDel %v", iidsToAdd, iidsToDel)

	if updateErr := models.UpdateInterviewSelection(app, interviews, iidsToAdd, iidsToDel); updateErr != nil {
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
func checkStepInInterviewSelectStatus(interviewType pkg.GroupOrTeam, app *pkg.Application) error {
	if interviewType == pkg.InGroup && app.Step != pkg.GroupTimeSelection {
		return fmt.Errorf("you can't set group interview time now")
	}
	if interviewType == pkg.InTeam && app.Step != pkg.TeamTimeSelection {
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

func getAddAndDelInterviews(originInterviews []pkg.Interview, selectIids []string) (iidsToAdd []string, iidsToDel []string) {
	for i := range originInterviews {
		ok := false
		for _, selectIid := range selectIids {
			if originInterviews[i].Uid == selectIid {
				ok = true
				break
			}
		}
		if !ok {
			iidsToDel = append(iidsToDel, originInterviews[i].Uid)
		}
	}

	for _, selectIid := range selectIids {
		ok := false
		for i := range originInterviews {
			if originInterviews[i].Uid == selectIid {
				ok = true
				break
			}
		}
		if !ok {
			iidsToAdd = append(iidsToAdd, selectIid)
		}
	}
	return
}
