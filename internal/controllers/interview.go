package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// SetRecruitmentInterviews set group/team all interview times
// PUT /recruitment/:rid/interviews/:name
// Use put to prevent resource are duplicated
func SetRecruitmentInterviews(c *gin.Context) {
	var (
		r    *pkg.Recruitment
		user *pkg.UserDetail
		err  error
	)

	defer func() { common.Resp(c, nil, err) }()

	rid := c.Param("rid")
	name := c.Param("name")
	uid := common.GetUID(c)
	if rid == "" || name == "" {
		err = fmt.Errorf("request param wrong, you should set rid and name")
		return
	}

	var interviews []pkg.UpdateInterviewOpts
	if err := c.ShouldBind(&interviews); err != nil {
		common.Error(c, rerror.RequestBodyError.WithData(err.Error()))
		return
	}

	// judge whether the recruitment has expired
	r, err = models.GetRecruitmentById(rid)
	if err != nil {
		return
	}
	if r.End.Before(time.Now()) {
		err = fmt.Errorf("recruitment %s has already ended", r.Name)
		return
	}

	user, err = grpc.GetUserInfoByUID(uid)
	if err != nil {
		return
	}

	// member can only update his group's interview or team interview (组面/群面
	if err = checkInterviewGroupName(user, name); err != nil {
		return
	}

	var interviewsToAdd []*pkg.UpdateInterviewOpts
	interviewsToUpdate := make(map[string]*pkg.UpdateInterviewOpts)
	for _, interview := range interviews {
		if interview.Uid != "" {
			// update
			interviewsToUpdate[interview.Uid] = &pkg.UpdateInterviewOpts{
				Date:       interview.Date,
				Period:     interview.Period,
				SlotNumber: interview.SlotNumber,
				Uid:        interview.Uid,
			}
		} else {
			// add
			interviewsToAdd = append(interviewsToAdd, &pkg.UpdateInterviewOpts{
				Date:       interview.Date,
				Period:     interview.Period,
				SlotNumber: interview.SlotNumber,
			})
		}
	}

	originInterviews, err := models.GetInterviewsByRidAndName(rid, name)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("interviews").WithDetail("when you update interviews"))
		return
	}

	var errors []string

	for _, origin := range originInterviews {
		value, ok := interviewsToUpdate[origin.Uid]
		if ok {
			// check whether the interview time has been selected by candidates
			if len(origin.Applications) != 0 && (!utils.ComPareTimeHour(origin.Date, value.Date) || origin.Period != value.Period) {
				errors = append(errors, fmt.Sprintf("interview %v have been selected", origin))
			} else {
				origin.Date = value.Date
				origin.SlotNumber = value.SlotNumber
				origin.Period = value.Period
				if err := models.UpdateInterview(&origin); err != nil {
					errors = append(errors, fmt.Sprintf("update interview %v on db failed, err: %v", origin, err))
				}
			}
		} else {
			if len(origin.Applications) != 0 {
				// when some candidates have selected this interview time, abort delete
				errors = append(errors, fmt.Sprintf("interview %v have been selected", origin))
			} else {
				if err := models.RemoveInterviewByID(origin.Uid); err != nil {
					errors = append(errors, fmt.Sprintf("delete interview %v on db failed, err: %v", origin, err))
				}
			}
		}
	}

	for _, interview := range interviewsToAdd {
		if err := models.CreateAndSaveInterview(interview); err != nil {
			errors = append(errors, fmt.Sprintf("save interview %v on db failed, err: %v", interview, err))
		}
	}
	if len(errors) != 0 {
		err = fmt.Errorf("%v", errors)
		return
	}
	return
}

// check user's group == name
func checkInterviewGroupName(user *pkg.UserDetail, name string) error {
	if name != "unique" {
		if !utils.CheckInGroups(user.Groups, name) {
			err := fmt.Errorf("you can't set other group's interview time")
			return err
		}
	}
	return nil
}
