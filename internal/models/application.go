package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
	"errors"
	"fmt"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateApplication(opts *pkg.CreateAppOpts, uid string, filePath string) (*pkg.Application, error) {
	db := global.GetDB()
	app := &pkg.Application{
		Grade:         opts.Grade,
		Institute:     opts.Institute,
		Major:         opts.Major,
		Rank:          opts.Rank,
		Group:         opts.Group,
		Intro:         opts.Intro,
		RecruitmentID: opts.RecruitmentID,
		Referrer:      opts.Referrer,
		IsQuick:       opts.IsQuick,
		Resume:        filePath,
		CandidateID:   uid,
		Step:          pkg.SignUp,
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if errdb := tx.Create(app).Error; errdb != nil {
			return errdb
		}
		//upload resume to COS
		if filePath != "" {
			errfile := global.UpLoadAndSaveFileToCos(opts.Resume, filePath)
			if errfile != nil {
				zapx.Error("upload resume to tos failed", zap.String("filepath", filePath))
				return errfile
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return app, nil
}

func GetApplicationByIdForCandidate(aid string) (*pkg.Application, error) {
	db := global.GetDB()
	var a pkg.Application
	if err := db.Preload("InterviewSelections", func(db *gorm.DB) *gorm.DB {
		return db.Omit("selectNumber") // omit selectNumber when candidate get
	}).Preload("InterviewAllocationsGroup", func(db *gorm.DB) *gorm.DB {
		return db.Omit("selectNumber") // omit selectNumber when candidate get
	}).Preload("InterviewAllocationsTeam", func(db *gorm.DB) *gorm.DB {
		return db.Omit("selectNumber") // omit selectNumber when candidate get
	}).
		Where("uid = ?", aid).
		First(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

// GetApplicationById For member
func GetApplicationById(aid string) (*pkg.Application, error) {
	db := global.GetDB()
	var a pkg.Application
	if err := db.Preload("Comments").
		Preload("InterviewSelections").
		Preload("InterviewAllocationsGroup").
		Preload("InterviewAllocationsTeam").
		Where("uid = ?", aid).
		First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func UpdateApplication(opts *pkg.UpdateAppOpts, filePath string) (*pkg.Application, error) {
	db := global.GetDB()

	var a pkg.Application
	if err := db.Model(&pkg.Application{}).
		Where("uid = ?", opts.Aid).
		First(&a).Error; err != nil {
		return nil, err
	}

	if opts.Grade != "" {
		a.Grade = opts.Grade
	}
	if opts.Institute != "" {
		a.Institute = opts.Institute
	}
	if opts.Major != "" {
		a.Major = opts.Major
	}
	if opts.Rank != "" {
		a.Rank = opts.Rank
	}
	if opts.Group != "" {
		a.Group = opts.Group
	}
	if opts.Intro != "" {
		a.Intro = opts.Intro
	}
	if opts.IsQuick != nil {
		a.IsQuick = *opts.IsQuick
	}
	if opts.Referrer != "" {
		a.Referrer = opts.Referrer
	}
	if opts.Resume != nil {
		a.Resume = filePath
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if errdb := tx.Save(&a).Error; errdb != nil {
			return errdb
		}

		//upload resume to COS
		if opts.Resume != nil {
			if errfile := global.UpLoadAndSaveFileToCos(opts.Resume, filePath); errfile != nil {
				zapx.Error("upload resume to tos failed", zap.String("filepath", filePath))
				return errfile
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &a, nil
}

func UpdateApplicationStep(aid string, step string) error {
	db := global.GetDB()
	if err := db.Model(&pkg.Application{}).Where("uid = ?", aid).Update("step", step).Error; err != nil {
		return err
	}
	return nil
}

func DeleteApplication(aid string) error {
	db := global.GetDB()
	return db.Where("uid = ?", aid).Delete(&pkg.Application{}).Error
}

func AbandonApplication(aid string) error {
	db := global.GetDB()
	application, err := GetApplicationByIdForCandidate(aid)
	if err != nil {
		return err
	}
	application.Abandoned = true
	return db.Updates(&application).Error
}

func RejectApplication(aid string) error {
	db := global.GetDB()
	application, err := GetApplicationByIdForCandidate(aid)
	if err != nil {
		return err
	}
	application.Rejected = true
	return db.Updates(&application).Error
}

func GetApplicationsByRid(rid string) ([]pkg.Application, error) {
	recruitment, err := GetFullRecruitmentById(rid)
	if err != nil {
		return nil, err
	}

	return recruitment.Applications, nil
}

func SetApplicationStepById(opts *pkg.SetAppStepOpts) error {
	db := global.GetDB()
	app, err := GetApplicationByIdForCandidate(opts.Aid)
	if err != nil {
		return err
	}

	if app.Step != opts.From {
		return errors.New("the step doesn't match")
	}
	if app.Abandoned || app.Rejected {
		return fmt.Errorf("application of %s has already been abandoned/reject", app.Uid)
	}
	return db.Model(&pkg.Application{}).
		Where("uid = ?", app.Uid).
		Updates(map[string]interface{}{
			"step": opts.To,
		}).Error
}

func SetApplicationInterviewTime(opts *pkg.SetAppInterviewTimeOpts) error {
	db := global.GetDB()
	if _, err := GetInterviewById(opts.InterviewId); err != nil {
		return err
	}

	application, err := GetApplicationByIdForCandidate(opts.Aid)
	if err != nil {
		return err
	}

	switch opts.InterviewType {
	case pkg.InGroup:
		err = db.Model(&pkg.Application{}).
			Where("uid = ?", application.Uid).
			Update("\"interviewAllocationsGroupId\"", opts.InterviewId).Error
	case pkg.InTeam:
		err = db.Model(&pkg.Application{}).
			Where("uid = ?", application.Uid).
			Update("\"interviewAllocationsTeamId\"", opts.InterviewId).Error
	}
	return err
}

func UpdateInterviewSelection(app *pkg.Application, interviews []pkg.Interview, iidsToAdd, iidsToDel []string) error {
	db := global.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		if errDb := tx.Model(app).
			Association("InterviewSelections").
			Clear(); errDb != nil {
			return errDb
		}

		app.InterviewSelections = interviews
		if errDb := tx.Save(app).Error; errDb != nil {
			return errDb
		}

		if errDb := tx.Model(&pkg.Interview{}).
			Where("uid IN ?", iidsToAdd).
			Updates(map[string]interface{}{
				"selectNumber": gorm.Expr("\"selectNumber\" + ?", 1),
			}).Error; errDb != nil {
			return errDb
		}

		if errDb := tx.Model(&pkg.Interview{}).
			Where("uid IN ?", iidsToDel).
			Updates(map[string]interface{}{
				"selectNumber": gorm.Expr("\"selectNumber\" - ?", 1),
			}).Error; errDb != nil {
			return errDb
		}
		return nil
	})
	return err
}

// TODO 上面的几个更新函数统一改调这个
func UpdateApplicationInfo(application *pkg.Application) error {
	db := global.GetDB()
	return db.Updates(&application).Error
}

func GetApplicationsByUserId(userId string) (*[]pkg.Application, error) {
	db := global.GetDB()
	var apps []pkg.Application
	if err := db.Preload("InterviewSelections", func(db *gorm.DB) *gorm.DB {
		return db.Omit("selectNumber", "slotNumber") // omit selectNumber when candidate get
	}).Preload("InterviewAllocationsGroup", func(db *gorm.DB) *gorm.DB {
		return db.Omit("selectNumber", "slotNumber") // omit selectNumber when candidate get
	}).Preload("InterviewAllocationsTeam", func(db *gorm.DB) *gorm.DB {
		return db.Omit("selectNumber", "slotNumber") // omit selectNumber when candidate get
	}).
		Where("\"candidateId\" = ?", userId).
		Find(&apps).Error; err != nil {
		return nil, err
	}
	return &apps, nil
}
