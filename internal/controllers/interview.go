package controllers

import (
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
	"time"
)

// PUT /recruitment/:rid/interviews/:name
// member group

// SetRecruitmentInterviews set group/team interview time
func SetRecruitmentInterviews(c *gin.Context) {
	// todo (get member info
	rid := c.Param("rid")
	name := c.Param("name")

	var interviews []request.UpdateInterviewRequest
	if err := c.ShouldBind(&interviews); err != nil {
		response.ResponseError(c, msg.RequestBodyError.WithData(err.Error()))
		return
	}

	// judge whether the recruitment has expired
	resp, err := models.GetRecruitmentById(rid)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	if resp.End.After(time.Now()) {
		response.ResponseError(c, msg.RecruitmentEnd.WithData(resp.Name))
		return
	}

	// member can only update his group's interview or team interview (组面/群面
	// todo (get member' group
	//if name != constants.InTeam && member.Group != name {
	//	response.ResponseError(c, msg.GroupNotMatch)
	//}

	var interviewsToAdd []*request.UpdateInterviewRequest
	interviewsToUpdate := make(map[string]*request.UpdateInterviewRequest)
	for _, interview := range interviews {
		if interview.Uid != "" {
			// update
			interviewsToUpdate[interview.Uid] = &request.UpdateInterviewRequest{
				Date:       interview.Date,
				Period:     interview.Period,
				SlotNumber: interview.SlotNumber,
				Uid:        interview.Uid,
			}
		} else {
			// add
			interviewsToAdd = append(interviewsToAdd, &request.UpdateInterviewRequest{
				Date:       interview.Date,
				Period:     interview.Period,
				SlotNumber: interview.SlotNumber,
			})
		}
	}

	originInterviews, err := models.GetInterviewsByRidAndName(rid, name)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("interviews").WithDetail("when you update interviews"))
		return
	}

	for _, origin := range *originInterviews {
		value, ok := interviewsToUpdate[origin.Uid]
		if ok {
			if len(origin.Applications) != 0 && (origin.Date != value.Date || origin.Period != value.Period) {
				response.ResponseError(c, msg.InterviewUpdateError.WithData("the interview time has been selected"))
				return
			} else {
				origin.Date = value.Date
				origin.SlotNumber = value.SlotNumber
				origin.Period = value.Period
				if err := models.UpdateInterview(&origin); err != nil {
					response.ResponseError(c, msg.UpdateDatabaseError.WithData("interview").WithDetail(err.Error()))
					return
				}
			}
		} else {
			if len(origin.Applications) != 0 {
				// when some candidates have selected this interview time, abort delete
				response.ResponseError(c, msg.InterviewHasBeenSelected.WithData("interview"))
				return
			} else {
				if err := models.RemoveInterviewByID(origin.Uid); err != nil {
					response.ResponseError(c, msg.UpdateDatabaseError.WithData("interview").WithDetail(err.Error()))
					return
				}
			}
		}
	}

	for _, interview := range interviewsToAdd {
		if err := models.CreateAndSaveInterview(interview); err != nil {
			response.ResponseError(c, msg.SaveDatabaseError.WithData("interview"))
			return
		}
	}

	response.ResponseOK(c, "Update interviews success", nil)
}
