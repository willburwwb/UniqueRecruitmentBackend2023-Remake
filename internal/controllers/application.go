package controllers

import (
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/pkg/msg"
	"UniqueRecruitmentBackend/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// CreateApplication create an application. Remember to transfer data with form instead of json!!!
// Accept role >=candidate
func CreateApplication(c *gin.Context) {
	var req request.CreateApplicationRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError(c, msg.RequestBodyError.WithDetail(err.Error()))
		return
	}
	recruitment, err := models.GetRecruitmentById(req.RecruitmentID)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("recruitment").WithDetail("when you submit the application"))
		return
	}
	// Compare the new recruitment time with application time
	if !checkApplyTime(c, recruitment, time.Now()) {
		return
	}

	//design must summit resume
	if req.Group == "Design" && req.Resume == nil {
		response.ResponseError(c,
			msg.RequestBodyError.WithDetail("candidate who sign up for the design group must have a resume"))
		return
	}

	//TODO(wwb)
	//when sso done,fix this filePath->user's uid
	// file path example: 2023秋(rname)/web(group)/wwb(userID)/filename
	filePath := fmt.Sprintf("%s/%s/%s/%s", recruitment.Name, req.Group, "thisisuserid", req.Resume.Filename)

	log.Println(filePath)
	//resume upload to COS
	err = upLoadAndSaveFileToCos(req.Resume, filePath)
	if err != nil {
		//TODO(wwb)
		//when sso done,fix this filePath->user's uid
		response.ResponseError(c, msg.UpLoadFileError.WithData("thisisuserid").WithDetail(err.Error()))
		return
	}

	//save application to database
	application, err := models.CreateAndSaveApplication(&req, filePath)
	if err != nil {
		response.ResponseError(c, msg.SaveDatabaseError.WithDetail(err.Error()))
		return
	}
	response.ResponseOK(c, "Success save application", application)
}

// GetApplicationByIdCandidate candidate views
// Get /applications/candidate
// Query aid
func GetApplicationByIdCandidate(c *gin.Context) {
	applicationId := c.Query("aid")
	// TODO(wwb)
	// check the role of user
	// then redirect to candidateGetApplicationById or memberGetApplicationById
	// :ForUser
	userID := utils.GetUserId(c)
	// check whether user has role to select application
	isHasRole, err := checkCandidateHasRole(applicationId, userID)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("application").
			WithDetail(err.Error()).WithDetail("when check user's role"))
		return
	}
	if !isHasRole {
		response.ResponseError(c, msg.RoleError.WithData(userID, "select application"))
		return
	}
	// select user's application
	application, err := models.FindOneByIdForCandidate(applicationId)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("application").
			WithDetail(err.Error()))
		return
	}
	response.ResponseOK(c, "candidate success get application", application)
	return

}

// GetApplicationByIdForMember member/admin views
// Get /applications/member
// Query aid
func GetApplicationByIdForMember(c *gin.Context) {
	//:ForMember
	applicationId := c.Query("aid")
	application, err := models.FindOneByIdForMember(applicationId)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("application").
			WithDetail(err.Error()))
		return
	}
	response.ResponseOK(c, "member success get application", application)
	return
}

// UpdateApplicationByCandidate this api can only be called by candidate!!
// Put /applications
// Query aid
func UpdateApplicationByCandidate(c *gin.Context) {
	var req request.UpApplicationRequest
	if err := c.ShouldBind(&req); err != nil {
		response.ResponseError(c, msg.RequestBodyError.WithDetail(err.Error()))
		return
	}
	recruitment, err := models.GetRecruitmentById(req.RecruitmentID)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("recruitment").WithDetail("when you update the application"))
		return
	}
	// Compare the new recruitment time with application time
	if !checkApplyTime(c, recruitment, time.Now()) {
		return
	}
	if req.Resume != nil {
		// Update the resume to cps
	}
	//application, err := models.CreateAndSaveApplication(&req, filePath)
	//if err != nil {
	//	response.ResponseError(c, msg.SaveDatabaseError.WithDetail(err.Error()))
	//	return
	//}
	//response.ResponseOK(c, "Success save application", application)
}

func DeleteApplicationById(c *gin.Context) {

}
func checkCandidateHasRole(aid string, uid string) (bool, error) {
	resultBool, err := models.FindOneByIdAndUid(aid, uid)
	return resultBool, err
}
func checkApplyTime(c *gin.Context, recruitment *models.RecruitmentEntity, now time.Time) bool {
	if recruitment.Beginning.After(now) {
		// submit too early
		response.ResponseError(c, msg.RecruitmentNotReady.WithData(recruitment.Name))
		return false
	} else if recruitment.Deadline.Before(now) {
		log.Println(recruitment.Deadline, now)
		response.ResponseError(c, msg.RecruitmentStopped.WithData(recruitment.Name))
		return false
	} else if recruitment.End.Before(now) {
		response.ResponseError(c, msg.RecruitmentEnd.WithData(recruitment.Name))
		return false
	}
	return true
}
