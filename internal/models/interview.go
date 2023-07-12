package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/request"
	"time"
)

type InterviewEntity struct {
	Common
	Date          time.Time            `gorm:"not null;uniqueIndex:interviews_all"`
	Period        constants.Period     `gorm:"not null;uniqueIndex:interviews_all"` //constants.Period
	Name          constants.Group      `gorm:"not null;uniqueIndex:interviews_all"` //constants.Group
	SlotNumber    int                  `gorm:"column:slotNumber;not null"`
	RecruitmentID string               `gorm:"column:recruitmentId;type:uuid;uniqueIndex:interviews_all"` //manytoone
	Applications  []*ApplicationEntity `gorm:"many2many:interview_selections"`                            //manytomany
}

func (c InterviewEntity) TableName() string {
	return "interviews"
}

func GetInterviewsByRidAndName(rid string, name string) (*[]InterviewEntity, error) {
	db := global.GetDB()
	var res []InterviewEntity
	if err := db.Model(&InterviewEntity{}).Preload("Applications").Where("rid = ? AND name = ?", rid, name).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func UpdateInterview(interview *InterviewEntity) error {
	ui := InterviewEntity{
		Date:       interview.Date,
		Period:     interview.Period,
		SlotNumber: interview.SlotNumber,
	}

	db := global.GetDB()

	return db.Updates(&ui).Error
}

func CreateAndSaveInterview(interview *request.UpdateInterviewRequest) error {
	return nil
}

func CreateInterviewsInBatches(interviews []InterviewEntity) error {
	db := global.GetDB()
	return db.Create(&interviews).Error
}

func RemoveInterviewByID(iid string) error {
	db := global.GetDB()
	return db.Delete(&InterviewEntity{}, iid).Error
}
