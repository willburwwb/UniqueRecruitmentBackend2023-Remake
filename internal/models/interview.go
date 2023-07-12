package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/request"
	"time"
)

type InterviewEntity struct {
	Common
	Date          time.Time            `gorm:"not null;uniqueIndex:interviews_all"`
	Period        string               `gorm:"not null;uniqueIndex:interviews_all"` //constants.Period
	Name          string               `gorm:"not null;uniqueIndex:interviews_all"` //constants.GroupOrTeam
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
	if err := db.Model(&InterviewEntity{}).Where("rid = ? AND name = ?", rid, name).Find(&res).Error; err != nil {
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

	return db.Model(&InterviewEntity{}).Updates(&ui).Error
}

func CreateAndSaveInterview(interview *request.UpdateInterviewRequest) error {
	return nil
}
