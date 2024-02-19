package controllers

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"fmt"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateRecruitment create recruitment
// @Id create_recruitment.
// @Summary create recruitment.
// @Description gcreate recruitment, only can be created by admin
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Param 	pkg.CreateRecOpts body pkg.CreateRecOpts true "create recruitment opts"
// @Success 200 {object} common.JSONResult{data=pkg.Recruitment} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments [post]
func CreateRecruitment(c *gin.Context) {
	var (
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, r, err) }()

	opts := &pkg.CreateRecOpts{}
	if err = c.ShouldBind(opts); err != nil {
		return
	}

	if err = opts.Validate(); err != nil {
		return
	}

	r, err = models.CreateRecruitment(opts)
	if err != nil {
		zapx.Error("save recruitment wrong", zap.Error(err))
		return
	}

	zapx.Info("success create recruitment")
	return
}

// UpdateRecruitment update recruitment
// @Id update_recruitment.
// @Summary update recruitment.
// @Description update recruitment, only can be updated by admin
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Param 	rid path string true "recruitment uid"
// @Param 	pkg.UpdateRecOpts body pkg.UpdateRecOpts true "update recruitment opts"
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments/{rid} [put]
func UpdateRecruitment(c *gin.Context) {
	var (
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.UpdateRecOpts{}
	opts.Rid = c.Param("rid")
	if err = c.ShouldBindJSON(opts); err != nil {
		return
	}
	if err = opts.Validate(); err != nil {
		return
	}

	if err := models.UpdateRecruitment(opts); err != nil {
		zapx.Error("update recruitment failed", zap.Error(err))
		return
	}
	zapx.Info("success update recruitment")
	return
}

// GetRecruitmentById get recruitment
// @Id get_recruitment.
// @Summary get recruitment.
// @Description get recruitment, member can only get the recruitment's detail(include application, interviews) after join in.
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Param 	rid path string true "recruitment uid"
// @Success 200 {object} common.JSONResult{data=pkg.Recruitment} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments/{rid} [get]
func GetRecruitmentById(c *gin.Context) {
	var (
		r    *pkg.Recruitment
		user *pkg.UserDetail
		err  error
	)
	defer func() { common.Resp(c, r, err) }()

	opts := &pkg.GetRecOpts{}
	if err = c.ShouldBindUri(opts); err != nil {
		return
	}

	// member role, return interviews + applications
	if common.IsMember(c) {
		user, err = grpc.GetUserInfoByUID(common.GetUID(c))
		if err != nil {
			return
		}

		r, err = models.GetRecruitmentById(opts.Rid)
		// todo(wwb) member join in after recruitment
		if !checkJoinTime(user.JoinTime, r.Beginning) {
			zapx.Warn("get old recruitment detail failed....")
		} else {
			r, err = models.GetFullRecruitmentById(opts.Rid)
			log.Println(r, err)
			//r.Statistics, err = models.GetRecruitmentStatistics(opts.Rid)
		}
	} else {
		r, err = models.GetRecruitmentById(opts.Rid)
	}
	return
}

// GetAllRecruitment get all recruitment
// @Id get_all_recruitment.
// @Summary get all recruitment.
// @Description get all recruitment, can only be got by member(not include applications and interviews).
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Success 200 {object} common.JSONResult{data=[]pkg.Recruitment} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments [get]
func GetAllRecruitment(c *gin.Context) {
	var (
		recruitments []pkg.Recruitment
		err          error
	)
	defer func() { common.Resp(c, recruitments, err) }()

	recruitments, err = models.GetAllRecruitment()
	for i := range recruitments {
		recruitments[i].Statistics, err = models.GetRecruitmentStatistics(recruitments[i].Uid)
		if err != nil {
			return
		}
	}
	return
}

// GetPendingRecruitment get pending recruitment
// @Id get_pending_recruitment.
// @Summary get pending recruitment.
// @Description get pending(the most recent) recruitment, member can only get the recruitment's detail(include application, interviews) after join in.
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Success 200 {object} common.JSONResult{data=pkg.Recruitment} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments [get]
func GetPendingRecruitment(c *gin.Context) {
	var (
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, r, err) }()

	r, err = models.GetPendingRecruitment()
	if err != nil {
		return
	}

	if common.IsMember(c) {
		r, err = models.GetFullRecruitmentById(r.Uid)
		r.Statistics, err = models.GetRecruitmentStatistics(r.Uid)
	} else {
		r, err = models.GetRecruitmentById(r.Uid)
	}

	return
}

// SetStressTestTime set StressTest time
// @Id set_stress_test_time
// @Summary set stress test start and end time.
// @Description set stress test start and end time.
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments/{rid}/stressTest [put]
func SetStressTestTime(c *gin.Context) {
	var (
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.SetStressTestTimeOpts{}
	if err = c.ShouldBind(opts); err != nil {
		return
	}
	opts.Rid = c.Param("rid")
	if err = opts.Validate(); err != nil {
		return
	}

	err = models.UpdateStressTestTime(opts)
	return
}

// UploadRecruitmentFile upload recruitment file
// @Id upload_recruitment_file
// @Summary upload recruitment file, such as written test.
// @Description upload recruitment file, such as written test.
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments/{rid}/file/{group}/{type} [put]
func UploadRecruitmentFile(c *gin.Context) {
	var (
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, nil, err) }()

	opts := &pkg.UploadRecruitmentFileOpts{}
	if err = c.ShouldBindUri(opts); err != nil {
		return
	}
	if err = c.ShouldBind(opts); err != nil {
		return
	}
	if err = opts.Validate(); err != nil {
		return
	}

	r, err = models.GetRecruitmentById(opts.Rid)
	if err != nil {
		return
	}

	// file path example: 2023秋(rname)/web(group)/WrittenTest.pdf(type)
	filePath := fmt.Sprintf("%s/%s/%s", r.Name, opts.Group, fmt.Sprintf("%s.pdf", opts.Type))
	err = global.UpLoadAndSaveFileToCos(opts.File, filePath)
	return
}

// DownloadRecruitmentFile download recruitment file
// @Id download_recruitment_file
// @Summary download recruitment file, such as written test.
// @Description upload recruitment file, such as written test.
// @Tags recruitment
// @Accept  json
// @Produce  json
// @Success 200 {object} common.JSONResult{} ""
// @Failure 400 {object} common.JSONResult{} "code is not 0 and msg not empty"
// @Router /recruitments/{rid}/file/{type}/{group} [get]
func DownloadRecruitmentFile(c *gin.Context) {
	var (
		r   *pkg.Recruitment
		err error
	)

	opts := &pkg.DownloadRecruitmentFileOpts{}
	if err = c.ShouldBindUri(opts); err != nil {
		common.Resp(c, nil, err)
		return
	}
	if err = opts.Validate(); err != nil {
		common.Resp(c, nil, err)
		return
	}

	r, err = models.GetRecruitmentById(opts.Rid)
	if err != nil {
		common.Resp(c, nil, err)
		return
	}

	// file path example: 2023秋(rname)/web(group)/WrittenTest.pdf(type)
	filePath := fmt.Sprintf("%s/%s/%s", r.Name, opts.Group, fmt.Sprintf("%s.pdf", opts.Type))
	resp, err := global.GetCOSObjectResp(filePath)
	if err != nil {
		common.Resp(c, nil, err)
		return
	}

	reader := resp.Body
	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
}

func checkJoinTime(joinTime string, recruitmentTime time.Time) bool {
	return true
}
