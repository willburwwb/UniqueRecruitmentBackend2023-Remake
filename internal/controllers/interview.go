package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"time"

	"github.com/gin-gonic/gin"
)

// SetRecruitmentInterviews set group/team all interview times
// PUT /recruitment/:rid/interviews/:name
// Use put to prevent resource are duplicated
func SetRecruitmentInterviews(c *gin.Context) {
	// todo (get member info
	rid := c.Param("rid")
	name := c.Param("name")

	var interviews []pkg.UpdateInterviewOpts
	if err := c.ShouldBind(&interviews); err != nil {
		common.Error(c, rerror.RequestBodyError.WithData(err.Error()))
		return
	}

	// judge whether the recruitment has expired
	resp, err := models.GetRecruitmentById(rid)
	if err != nil {
		common.Error(c, rerror.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	if resp.End.Before(time.Now()) {
		common.Error(c, rerror.RecruitmentEnd.WithData(resp.Name))
		return
	}

	// member can only update his group's interview or team interview (组面/群面
	// todo (get member' group
	if !checkGroupName(c, name) {
		common.Error(c, rerror.CheckPermissionError.WithDetail("you are not in this group"))
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

	for _, origin := range *originInterviews {
		value, ok := interviewsToUpdate[origin.Uid]
		if ok {
			// check whether the interview time has been selected by candidates

			if len(origin.Applications) != 0 && (!utils.ComPareTimeHour(origin.Date, value.Date) || origin.Period != value.Period) {
				//	common.Error(c, rerror.InterviewUpdateError.WithData("the interview time has been selected"))
				//	return
				errors = append(errors, rerror.InterviewHasBeenSelected.WithData(origin.Uid).Msg())
			} else {
				origin.Date = value.Date
				origin.SlotNumber = value.SlotNumber
				origin.Period = value.Period
				if err := models.UpdateInterview(&origin); err != nil {
					//	common.Error(c, rerror.UpdateDatabaseError.WithData("interview").WithDetail(err.Error()))
					//	return
					errors = append(errors, rerror.UpdateDatabaseError.WithData("interview").Msg()+err.Error())
				}
			}
		} else {
			if len(origin.Applications) != 0 {
				// when some candidates have selected this interview time, abort delete
				//	common.Error(c, rerror.InterviewHasBeenSelected.WithData("interview"))
				//	return
				errors = append(errors, rerror.InterviewHasBeenSelected.WithData(origin.Uid).Msg())
			} else {
				if err := models.RemoveInterviewByID(origin.Uid); err != nil {
					//		common.Error(c, rerror.UpdateDatabaseError.WithData("interview").WithDetail(err.Error()))
					//		return
					errors = append(errors, rerror.RemoveDatabaseError.WithData("interview").Msg()+err.Error())
				}
			}
		}
	}

	for _, interview := range interviewsToAdd {
		if err := models.CreateAndSaveInterview(interview); err != nil {
			//	common.Error(c, rerror.SaveDatabaseError.WithData("interview"))
			//	return
			errors = append(errors, rerror.SaveDatabaseError.WithData("interview").Msg())

		}
	}
	if len(errors) != 0 {
		common.Error(c, rerror.UpdateDatabaseError.WithData("interview").WithDetail(errors...))
		return
	}
	common.Success(c, "Update interviews success", nil)
}

// check user's group == name
func checkGroupName(c *gin.Context, name string) bool {
	if name != "unique" {
		uid := common.GetUID(c)
		userInfo, err := grpc.GetUserInfoByUID(uid)
		if err != nil {
			common.Error(c, rerror.CheckPermissionError.WithDetail(err.Error()))
			return false
		}
		if !utils.CheckInGroups(userInfo.Groups, name) {
			common.Error(c, rerror.CheckPermissionError.WithDetail("you are not in this group"))
			return false
		}
	}
	return true
}
