package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	error2 "UniqueRecruitmentBackend/internal/error"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/utils"

	"github.com/gin-gonic/gin"
)
// Create recruitment interviews for candidate to select 
// POST /recruitment/:rid/interviews/:name
// member group
func CreateRecruitmentInterviews(c *gin.Context) {
	rid := c.Param("rid")
	name := c.Param("name")
	var interviews []request.CreateInterview
	if err := c.ShouldBind(&interviews); err != nil {
		common.Error(c, error2.RequestBodyError.WithData(err.Error()))
		return
	}

	if !checkGroupName(c, name) {
		return
	}

	if err := models.CreateAndSaveInterview(rid, name, interviews); err != nil {
		common.Error(c, error2.SaveDatabaseError.WithData("interview").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success save interviews database", nil)

}
// Update recruitment interviews 
// PUT /recruitment/:rid/interviews/:name
// member group
func UpdateRecruitmentInterviews(c *gin.Context) {
	rid := c.Param("rid")
	name := c.Param("name")
	var interviews []request.UpdateInterview
	if err := c.ShouldBind(&interviews); err != nil {
		common.Error(c, error2.RequestBodyError.WithData(err.Error()))
		return
	}

	// check user's group == name
	if !checkGroupName(c, name) {
		return
	}

	if err := models.UpdateInterviews(rid, name, interviews); err != nil {
		common.Error(c, error2.SaveDatabaseError.WithData("interview").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success update interviews database", nil)
}
func DeleteRecruitmentInterviews(c *gin.Context) {
	name := c.Param("name")
	var interviews []request.DeleteInterviewUID
	if err := c.ShouldBind(&interviews); err != nil {
		common.Error(c, error2.RequestBodyError.WithData(err.Error()))
		return
	}
	// check user's group == name
	if !checkGroupName(c, name) {
		return
	}
	if err := models.DeleteInterviews(name, interviews); err != nil {
		common.Error(c, error2.SaveDatabaseError.WithData("interview").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success delete interviews database", nil)
}

// check user's group == name
func checkGroupName(c *gin.Context, name string) bool {
	if name != "unique" {
		uid := common.GetUID(c)
		userInfo, err := getUserInfoByUID(c, uid)
		if err != nil {
			common.Error(c, error2.CheckPermissionError.WithDetail(err.Error()))
			return false
		}
		if !utils.CheckInArrary(name, userInfo.Groups) {
			common.Error(c, error2.CheckPermissionError.WithDetail("you are not in this group"))
			return false
		}
	}
	return true
}

// SetRecruitmentInterviews set group/team all interview times
// func SetRecruitmentInterviews(c *gin.Context) {
// 	// todo (get member info
// 	rid := c.Param("rid")
// 	name := c.Param("name")

// 	var interviews []request.UpdateInterviewRequest
// 	if err := c.ShouldBind(&interviews); err != nil {
// 		common.Error(c, error2.RequestBodyError.WithData(err.Error()))
// 		return
// 	}

// 	// judge whether the recruitment has expired
// 	resp, err := models.GetRecruitmentById(rid, constants.CandidateRole)
// 	if err != nil {
// 		common.Error(c, error2.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
// 		return
// 	}
// 	if resp.End.Before(time.Now()) {
// 		common.Error(c, error2.RecruitmentEnd.WithData(resp.Name))
// 		return
// 	}

// 	// member can only update his group's interview or team interview (组面/群面
// 	// todo (get member' group
// 	//if name != constants.InTeam && member.Group != name {
// 	//	response.Error(c, msg.GroupNotMatch)
// 	//}

// 	var interviewsToAdd []*request.UpdateInterviewRequest
// 	interviewsToUpdate := make(map[string]*request.UpdateInterviewRequest)
// 	for _, interview := range interviews {
// 		if interview.Uid != "" {
// 			// update
// 			interviewsToUpdate[interview.Uid] = &request.UpdateInterviewRequest{
// 				Date:       interview.Date,
// 				Period:     interview.Period,
// 				SlotNumber: interview.SlotNumber,
// 				Uid:        interview.Uid,
// 			}
// 		} else {
// 			// add
// 			interviewsToAdd = append(interviewsToAdd, &request.UpdateInterviewRequest{
// 				Date:       interview.Date,
// 				Period:     interview.Period,
// 				SlotNumber: interview.SlotNumber,
// 			})
// 		}
// 	}

// 	originInterviews, err := models.GetInterviewsByRidAndName(rid, name)
// 	if err != nil {
// 		common.Error(c, error2.GetDatabaseError.WithData("interviews").WithDetail("when you update interviews"))
// 		return
// 	}

// 	for _, origin := range *originInterviews {
// 		value, ok := interviewsToUpdate[origin.Uid]
// 		if ok {
// 			if len(origin.Applications) != 0 && (origin.Date != value.Date || origin.Period != value.Period) {
// 				common.Error(c, error2.InterviewUpdateError.WithData("the interview time has been selected"))
// 				return
// 			} else {
// 				origin.Date = value.Date
// 				origin.SlotNumber = value.SlotNumber
// 				origin.Period = value.Period
// 				if err := models.UpdateInterview(&origin); err != nil {
// 					common.Error(c, error2.UpdateDatabaseError.WithData("interview").WithDetail(err.Error()))
// 					return
// 				}
// 			}
// 		} else {
// 			if len(origin.Applications) != 0 {
// 				// when some candidates have selected this interview time, abort delete
// 				common.Error(c, error2.InterviewHasBeenSelected.WithData("interview"))
// 				return
// 			} else {
// 				if err := models.RemoveInterviewByID(origin.Uid); err != nil {
// 					common.Error(c, error2.UpdateDatabaseError.WithData("interview").WithDetail(err.Error()))
// 					return
// 				}
// 			}
// 		}
// 	}

// 	for _, interview := range interviewsToAdd {
// 		if err := models.CreateAndSaveInterview(interview); err != nil {
// 			common.Error(c, error2.SaveDatabaseError.WithData("interview"))
// 			return
// 		}
// 	}

// 	common.Success(c, "Update interviews success", nil)
// }
