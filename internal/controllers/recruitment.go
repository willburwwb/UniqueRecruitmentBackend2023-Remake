package controllers

import (
	"UniqueRecruitmentBackend/internal/common"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/grpc"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"time"

	"github.com/gin-gonic/gin"
)

// Post recruitment/
// admin role

// CreateRecruitment create new recruitment
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

// Post recruitment/:rid/schedule
// admin role

// UpdateRecruitment update recruitment details
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

// Get recruitment/:rid

// GetRecruitmentById get recruitment details by id
// member can only get the recruitments' detail after join in
func GetRecruitmentById(c *gin.Context) {
	var (
		r   *pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, r, err) }()

	opts := &pkg.GetRecOpts{}
	if err = c.BindUri(opts); err != nil {
		return
	}

	// member role, return interviews + applications
	if common.IsMember(c) {
		user, err := grpc.GetUserInfoByUID(common.GetUID(c))
		if err != nil {
			return
		}

		r, err = models.GetRecruitmentById(opts.Rid)
		// todo(wwb) member join in after recruitment
		if !checkJoinTime(user.JoinTime, r.Beginning) {
			zapx.Warn("get old recruitment detail failed....")
		} else {
			r, err = models.GetFullRecruitmentById(opts.Rid)
		}
	} else {
		r, err = models.GetRecruitmentById(opts.Rid)
	}

	return
}

// Get recruitment/

// GetAllRecruitment get all recruitment details
func GetAllRecruitment(c *gin.Context) {
	var (
		r   []pkg.Recruitment
		err error
	)
	defer func() { common.Resp(c, r, err) }()

	r, err = models.GetAllRecruitment()
	return
}

// Get recruitment/:rid/pending

// GetPendingRecruitment get pending recruitment details
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
	} else {
		r, err = models.GetRecruitmentById(r.Uid)
	}
	return
}

func checkJoinTime(joinTime string, recruitmentTime time.Time) bool {
	return true
}
