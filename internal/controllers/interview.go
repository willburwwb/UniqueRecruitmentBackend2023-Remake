package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
)

func GetRecruitmentInterviews(c *gin.Context) {
	var (
		interviews []pkg.Interview
		err        error
	)

	defer func() { common.Resp(c, interviews, err) }()

	rid := c.Param("rid")
	name := c.Param("name")
	if rid == "" || name == "" {
		err = fmt.Errorf("request param wrong, you should set rid and name")
		return
	}

	interviews, err = models.GetInterviewsByRidAndNameWithoutApp(rid, name)
	if err != nil {
		return
	}

	return
}

// SetRecruitmentInterviews set group/team all interview times
// PUT /recruitment/:rid/interviews/:name
// Use put to prevent resource are duplicated
func SetRecruitmentInterviews(c *gin.Context) {
	var (
		r                *pkg.Recruitment
		user             *pkg.UserDetail
		originInterviews []pkg.Interview
		err              error
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
	if err = c.ShouldBind(&interviews); err != nil {
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

	var interviewsToAdd []pkg.Interview
	var interviewIdsToDel []string
	interviewsToUpdate := make(map[string]pkg.Interview)
	for _, interview := range interviews {
		if interview.Uid != "" {
			// update
			interviewsToUpdate[interview.Uid] = pkg.Interview{
				Name:          name,
				RecruitmentID: rid,
				Date:          interview.Date,
				Period:        interview.Period,
				SlotNumber:    interview.SlotNumber,
			}
		} else {
			// add
			interviewsToAdd = append(interviewsToAdd, pkg.Interview{
				Name:          name,
				RecruitmentID: rid,
				Date:          interview.Date,
				Period:        interview.Period,
				SlotNumber:    interview.SlotNumber,
			})
		}
	}

	originInterviews, err = models.GetInterviewsByRidAndName(rid, name)
	if err != nil {
		return
	}

	var errors []string

	for _, origin := range originInterviews {
		value, ok := interviewsToUpdate[origin.Uid]
		if ok {
			// check whether the interview time has been selected by candidates
			if len(origin.Applications) != 0 && (!utils.ComPareTimeHour(origin.Date, value.Date) || origin.Period != value.Period) {
				errors = append(errors, fmt.Sprintf("interview %v have been selected", origin))
			}
		} else {
			// delete
			if len(origin.Applications) != 0 {
				// when some candidates have selected this interview time, abort delete
				errors = append(errors, fmt.Sprintf("interview %v have been selected", origin))
			} else {
				interviewIdsToDel = append(interviewIdsToDel, origin.Uid)
			}
		}
	}

	if errdb := models.UpdateInterview(interviewsToAdd, interviewIdsToDel, interviewsToUpdate); errdb != nil {
		errors = append(errors, fmt.Sprintf("update interviews db failed, err: %s", errdb.Error()))
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
