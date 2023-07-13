package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/request"
	"encoding/json"
	"errors"
	"time"
)

// used for insert data without sso
const fakeCandidateId = "b234d3f4-1e74-11ee-8b78-b69bc9af8fe4"

type ApplicationEntity struct {
	Common
	Grade     string `gorm:"not null"` //constants.Grade
	Institute string `gorm:"not null"`
	Major     string `gorm:"not null"`
	Rank      string `gorm:"not null"` //constants.Rank
	Group     string `gorm:"not null"` //constants.Group
	Intro     string `gorm:"not null"`
	IsQuick   bool   `gorm:"column:isQuick;not null"`
	Referrer  string

	Resume string

	Abandoned                 bool               `gorm:"not null; default false" `
	Rejected                  bool               `gorm:"not null; default false"`
	Step                      string             `gorm:"not null"`                                                                //constants.Step
	CandidateID               string             `gorm:"column:candidateId;type:uuid;uniqueIndex:UQ_CandidateID_RecruitmentID"`   //manytoone
	RecruitmentID             string             `gorm:"column:recruitmentId;type:uuid;uniqueIndex:UQ_CandidateID_RecruitmentID"` //manytoone
	InterviewAllocationsGroup time.Time          `gorm:"column:interviewAllocationsGroup;"`
	InterviewAllocationsTeam  time.Time          `gorm:"column:interviewAllocationsTeam;"`
	InterviewSelections       []*InterviewEntity `gorm:"many2many:interview_selections;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`          //manytomany
	Comments                  []CommentEntity    `gorm:"foreignKey:ApplicationID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //onetomany
}

func (a ApplicationEntity) TableName() string {
	return "applications"
}

type ApplicationForCandidate struct {
	Common
	Grade         string
	Institute     string
	Major         string
	Rank          string
	Group         string
	Intro         string
	Referrer      string
	Resume        string
	Step          string
	RecruitmentID string
}

func CreateAndSaveApplication(req *request.CreateApplicationRequest, filename string) (*ApplicationEntity, error) {
	db := global.GetDB()
	row := db.Where("'recruitmentId' = ?", req.RecruitmentID).Find(&ApplicationEntity{}).RowsAffected

	//check now user's recruitment application >1
	if row != 0 {
		return nil, errors.New("A candidate can only apply once at the same recruitment")
	}
	a := ApplicationEntity{
		Grade:         req.Grade,
		Institute:     req.Institute,
		Major:         req.Major,
		Rank:          req.Rank,
		Group:         req.Group,
		Intro:         req.Intro,
		RecruitmentID: req.RecruitmentID,
		Referrer:      req.Referrer,
		IsQuick:       req.IsQuick,
		Resume:        filename,
		CandidateID:   fakeCandidateId,
		// TODO(wwb)
		// Add step status
		Step: "",
	}
	err := db.Create(&a).Error
	return &a, err
}

func GetApplicationByIdForCandidate(aid string) (*ApplicationForCandidate, error) {
	db := global.GetDB()
	var a ApplicationEntity

	if err := db.Where("uid = ?", aid).Find(&a).Error; err != nil {
		return nil, err
	}

	var afc ApplicationForCandidate
	bytes, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &afc); err != nil {
		return nil, err
	}

	return &afc, err
}

// GetApplicationById For member
func GetApplicationById(aid string) (*ApplicationEntity, error) {
	db := global.GetDB()
	var a ApplicationEntity
	if err := db.Preload("Comments").Where("uid = ?", aid).Find(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func UpdateApplication(aid string, filename string, req *request.UpdateApplicationRequest) error {
	req.Resume = nil
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	var a ApplicationEntity
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
	if err := db.Model(&ApplicationEntity{}).Where("uid = ?", aid).Update("step", step).Error; err != nil {
		return err
	}
	return nil
}

func DeleteApplication(aid string) error {
	db := global.GetDB()
	return db.Delete(&ApplicationEntity{}, aid).Error
}

func AbandonApplication(aid string) error {
	db := global.GetDB()
	applicationEntity, err := GetApplicationById(aid)
	if err != nil {
		return err
	}
	applicationEntity.Abandoned = true
	return db.Updates(&applicationEntity).Error
}

func GetApplicationByRecruitmentId(rid string) ([]ApplicationEntity, error) {
	recruitmentById, err := GetRecruitmentById(rid)
	if err != nil {
		return nil, err
	}

	return recruitmentById.Applications, nil
}

func SetApplicationStepById(aid string, req *request.SetApplicationStepRequest) error {
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

// TODO 上面的几个更新函数统一改调这个
func UpdateApplicationInfo(application *ApplicationEntity) error {
	db := global.GetDB()
	return db.Updates(&application).Error
}

/*

 */
