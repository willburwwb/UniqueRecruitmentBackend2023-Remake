package controllers

import (
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/pkg/msg"
	"github.com/gin-gonic/gin"
	"time"
)

// 原先接口是 /recruitment/:rid/interviews/:name

// PUT :rid/interviews/:name
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

}
