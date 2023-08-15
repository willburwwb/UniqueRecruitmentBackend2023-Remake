package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/request"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type InterviewEntity struct {
	Common
	Date          time.Time            `json:"date" gorm:"not null;uniqueIndex:interviews_all"`
	Period        constants.Period     `json:"period" gorm:"not null;uniqueIndex:interviews_all"` //constants.Period
	Name          constants.Group      `json:"name" gorm:"not null;uniqueIndex:interviews_all"`   //constants.Group
	SlotNumber    int                  `json:"slotNumber" gorm:"column:slotNumber;not null"`
	RecruitmentID string               `json:"recruitmentID" gorm:"column:recruitmentId;type:uuid;uniqueIndex:interviews_all"` //manytoone
	Applications  []*ApplicationEntity `json:"applications,omitempty" gorm:"many2many:interview_selections"`                   //manytomany
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

// func UpdateInterview(interview *InterviewEntity) error {
// 	ui := InterviewEntity{
// 		Date:       interview.Date,
// 		Period:     interview.Period,
// 		SlotNumber: interview.SlotNumber,
// 	}

// 	db := global.GetDB()

//		return db.Updates(&ui).Error
//	}
func CreateAndSaveInterview(rid string, name string, interviews []request.CreateInterview) error {
	var interviewEntitys []InterviewEntity
	for _, interview := range interviews {
		var interviewEntity InterviewEntity
		bytes, err := json.Marshal(interview)
		if err != nil {
			return err
		}
		json.Unmarshal(bytes, &interviewEntity)
		interviewEntity.RecruitmentID = rid
		interviewEntity.Name = constants.Group(name)
		interviewEntitys = append(interviewEntitys, interviewEntity)
	}
	db := global.GetDB()
	return db.Create(&interviewEntitys).Error
}
func UpdateInterviews(rid string, name string, interviews []request.UpdateInterview) error {
	var interviewEntitys []InterviewEntity
	for _, interview := range interviews {
		var interviewEntity InterviewEntity
		bytes, err := json.Marshal(interview)
		if err != nil {
			return err
		}
		json.Unmarshal(bytes, &interviewEntity)
		interviewEntity.RecruitmentID = rid
		interviewEntity.Name = constants.Group(name)
		interviewEntitys = append(interviewEntitys, interviewEntity)
	}
	db := global.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, interviewEntity := range interviewEntitys {
			errUpdate := tx.Updates(interviewEntity).Error
			if errUpdate != nil {
				return errUpdate
			}
		}
		return nil
	})
	return err
}
func DeleteInterviews(name string, interviews []request.DeleteInterviewUID) error {
	db := global.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, interview := range interviews {
			if errDelete := tx.Delete(&InterviewEntity{}, interview).Error; errDelete != nil {
				return errDelete
			}
		}
		return nil
	})
	return err
}
func CreateInterviewsInBatches(interviews []InterviewEntity) error {
	db := global.GetDB()
	return db.Create(&interviews).Error
}

func RemoveInterviewByID(iid string) error {
	db := global.GetDB()
	return db.Delete(&InterviewEntity{}, iid).Error
}
