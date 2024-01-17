package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"UniqueRecruitmentBackend/pkg/rerror"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"time"

	"github.com/gin-gonic/gin"
)

// Post recruitment/
// admin role

// CreateRecruitment create new recruitment
func CreateRecruitment(c *gin.Context) {
	var req pkg.CreateRecOpts
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	if time.Now().After(req.Beginning) || req.Beginning.After(req.Deadline) || req.Deadline.After(req.End) {
		zapx.Error("time set up wrong", zap.String("uid", common.GetUID(c)))
		common.Error(c, rerror.RequestBodyError.WithDetail("set up time failed"))
		return
	}

	recruitmentId, err := models.CreateRecruitment(&req)
	if err != nil {
		zapx.Error("save recruitment wrong", zap.Error(err))
		common.Error(c, rerror.SaveDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}

	zapx.Info("success create recruitment")
	common.Success(c, "Success create recruitment", map[string]interface{}{
		"rid": recruitmentId,
	})
}

// Post recruitment/:rid/schedule
// admin role

// UpdateRecruitment update recruitment details
func UpdateRecruitment(c *gin.Context) {
	var req pkg.UpdateRecOpts
	req.Rid = c.Param("rid")
	if req.Rid == "" {
		common.Error(c, rerror.RequestBodyError.WithDetail("recruitment id is null"))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, rerror.RequestBodyError.WithDetail(err.Error()))
		return
	}

	if err := models.UpdateRecruitment(&req); err != nil {
		zapx.Error("update recruitment failed", zap.Error(err))
		common.Error(c, rerror.UpdateDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	zapx.Info("success update recruitment")
	common.Success(c, "Success update recruitment", map[string]interface{}{
		"rid": req.Rid,
	})
}

// Get recruitment/:rid

// GetRecruitmentById get recruitment details by id
// member can only get the recruitments' detail after join in
func GetRecruitmentById(c *gin.Context) {
	recruitmentId := c.Param("rid")
	if recruitmentId == "" {
		common.Error(c, rerror.RequestBodyError.WithDetail("lost http query params [rid]"))
		return
	}

	// member role, return interviews + applications
	if common.IsMember(c) {
		user, err := grpc.GetUserInfoByUID(common.GetUID(c))
		if err != nil {
			zapx.Error("get recruitment failed", zap.Error(err))
			common.Error(c, rerror.SSOError.WithData("recruitment").WithDetail(err.Error()))
			return
		}

		resp, err := models.GetRecruitmentById(recruitmentId)
		if !checkJoinTime(user.JoinTime, resp.Beginning) {
			zapx.Warn("get old recruitment detail failed", zap.Error(err))
			common.Success(c, "Success get recruitment, but don't have role to get old recruitment detail", resp)
			return
		}

		respfull, err := models.GetFullRecruitmentById(recruitmentId)
		if err != nil {
			zapx.Error("get recruitment failed", zap.Error(err))
			common.Error(c, rerror.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
			return
		}
		common.Success(c, "Success get recruitment by member role", respfull)
	} else {
		resp, err := models.GetRecruitmentById(recruitmentId)
		if err != nil {
			zapx.Error("get recruitment failed", zap.Error(err))
			common.Error(c, rerror.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
			return
		}
		common.Success(c, "Success get recruitment by candidate role", resp)
	}
}

// Get recruitment/

// GetAllRecruitment get all recruitment details
func GetAllRecruitment(c *gin.Context) {
	resp, err := models.GetAllRecruitment()
	if err != nil {
		zapx.Error("get all recruitment error", zap.Error(err))
		common.Error(c, rerror.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success get all recruitment", resp)
}

// Get recruitment/:rid/pending

// GetPendingRecruitment get pending recruitment details
func GetPendingRecruitment(c *gin.Context) {
	var err error
	resp, err := models.GetPendingRecruitment()
	if err != nil {
		zapx.Error("get pending recruitment id error", zap.Error(err))
		common.Error(c, rerror.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}

	if common.IsMember(c) {
		resp, err = models.GetFullRecruitmentById(resp.Uid)
	} else {
		resp, err = models.GetRecruitmentById(resp.Uid)
	}

	if err != nil {
		zapx.Error("get pending recruitment error", zap.Error(err))
		common.Error(c, rerror.GetDatabaseError.WithData("recruitment").WithDetail(err.Error()))
		return
	}
	common.Success(c, "Success get pending recruitment", resp)
}

func checkJoinTime(joinTime string, recruitmentTime time.Time) bool {
	return true
}
