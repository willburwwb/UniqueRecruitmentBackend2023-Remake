package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/request"
	"UniqueRecruitmentBackend/internal/response"
	"UniqueRecruitmentBackend/internal/utils"
	"UniqueRecruitmentBackend/pkg/msg"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	aid := c.Param("aid")
	if common.IsCandidate("") {
		application, err := models.GetApplicationByIdForCandidate(aid)
		if err != nil {
			response.ResponseError(c, msg.GetDatabaseError.WithData("application").WithDetail("Get application info fail"))
			return
		}
		response.ResponseOK(c, "Get application success", application)
	} else {
		application, err := models.GetApplicationById(aid)
		if err != nil {
			response.ResponseError(c, msg.GetDatabaseError.WithData("application").WithDetail("Get application info fail"))
			return
		}
		response.ResponseOK(c, "Get application success", application)
	}
}

func UpdateApplicationById(c *gin.Context) {
	aid := c.Param("aid")
	var req request.UpdateApplicationRequest
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

	filePath := ""
	if req.Resume != nil {
		filePath = fmt.Sprintf("%s/%s/%s/%s", recruitment.Name, req.Group, "thisisuserid", req.Resume.Filename)
		if err := upLoadAndSaveFileToCos(req.Resume, filePath); err != nil {
			//TODO(wwb)
			//when sso done,fix this filePath->user's uid
			response.ResponseError(c, msg.UpLoadFileError.WithData("thisisuserid").WithDetail(err.Error()))
			return
		}
	}

	if err := models.UpdateApplication(aid, filePath, &req); err != nil {
		response.ResponseError(c, msg.UpdateDatabaseError.WithData("application").WithDetail(err.Error()))
		return
	}
	response.ResponseOK(c, "update application success", nil)
	return
}

func DeleteApplicationById(c *gin.Context) {
	aid := c.Param("aid")
	if err := models.DeleteApplication(aid); err != nil {
		response.ResponseError(c, msg.SaveDatabaseError.WithData("application"))
		return
	}
	response.ResponseOK(c, "delete application success", nil)
}

func AbandonApplicationById(c *gin.Context) {
	aid := c.Param("aid")
	if err := models.AbandonApplication(aid); err != nil {
		response.ResponseError(c, msg.SaveDatabaseError.WithData("application"))
		return
	}
	response.ResponseOK(c, "delete application success", nil)
}

func GetResumeById(c *gin.Context) {
	aid := c.Param("aid")
	application, err := models.GetApplicationById(aid)
	if err != nil {
		response.ResponseError(c, msg.GetDatabaseError.WithData("application").WithDetail("Get application info fail"))
		return
	}
	resp, err := utils.GetCOSObjectResp(application.Resume)
	if err != nil {
		response.ResponseError(c, msg.DownloadFileError.WithData("application").WithDetail("download resume fail"))
		return
	}

	reader := resp.Body
	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
	return
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
