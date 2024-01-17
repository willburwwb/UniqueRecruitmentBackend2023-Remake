package models

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"

	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/pkg"
)

func CreateApplication(req *pkg.CreateAppOpts, uid string, filePath string) error {
	db := global.GetDB()
	row := db.Where("'recruitmentId' = ?", req.RecruitmentID).
		Find(&pkg.Application{}).RowsAffected

	//check if user recruitment application's num >1
	if row != 0 {
		return errors.New("A candidate can only apply once at the same recruitment")
	}

	return db.Create(&pkg.Application{
		Grade:         req.Grade,
		Institute:     req.Institute,
		Major:         req.Major,
		Rank:          req.Rank,
		Group:         req.Group,
		Intro:         req.Intro,
		RecruitmentID: req.RecruitmentID,
		Referrer:      req.Referrer,
		IsQuick:       req.IsQuick,
		Resume:        filePath,
		CandidateID:   uid,
		Step:          string(constants.SignUp),
	}).Error
}

func GetApplicationByIdForCandidate(aid string) (*pkg.Application, error) {
	db := global.GetDB()

	var a pkg.Application
	if err := db.Preload("interview_selections").
		Where("uid = ?", aid).
		Find(&a).Error; err != nil {
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
		Where("uid = ?", aid).Find(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func UpdateApplication(aid string, filename string, req *pkg.UpdateAppOpts) error {
	req.Resume = nil
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var a pkg.Application
	if err := json.Unmarshal(bytes, &a); err != nil {
		return err
	}
	a.Uid = aid
	if filename != "" {
		a.Resume = filename
	}

	db := global.GetDB()
	return db.Updates(&a).Error
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
	application, err := GetApplicationById(aid)
	if err != nil {
		return err
	}
	application.Abandoned = true
	return db.Updates(&application).Error
}

func GetApplicationByRecruitmentId(rid string) ([]pkg.Application, error) {
	recruitment, err := GetFullRecruitmentById(rid)
	if err != nil {
		return nil, err
	}

	return recruitment.Applications, nil
}

func SetApplicationStepById(aid string, req *pkg.SetAppStepOpts) error {
	db := global.GetDB()
	application, err := GetApplicationById(aid)
	if err != nil {
		return err
	}
	if application.Step != req.From {
		return errors.New("the step doesn't match")
	}
	application.Step = req.To
	return db.Updates(&application).Error
}

func SetApplicationInterviewTime(aid, interviewType string, time time.Time) error {
	db := global.GetDB()
	application, err := GetApplicationById(aid)
	if err != nil {
		return err
	}

	switch interviewType {
	case "group":
		application.InterviewAllocationsGroup = time
	case "team":
		application.InterviewAllocationsTeam = time
	}

	return db.Updates(&application).Error
}

func UpdateInterviewSelection(application *pkg.Application, interviews []*pkg.Interview) error {
	db := global.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		if errClear := tx.Model(application).Association("InterviewSelections").Clear(); errClear != nil {
			return errClear
		}
		application.InterviewSelections = interviews
		if errUpdate := tx.Save(application).Error; errUpdate != nil {
			return errUpdate
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
