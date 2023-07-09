package controllers

import (
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/pkg/msg"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// CreateApplication create an application. Remember to transfer data with form instead of json!!!
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

func GetApplicationById(c *gin.Context) {
	//这里区分两种权限，选手和member会看到不同数据。
	//var applicationId string
	//applicationId = c.Query("applicationId")

}

func UpdateApplicationById(c *gin.Context) {

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
