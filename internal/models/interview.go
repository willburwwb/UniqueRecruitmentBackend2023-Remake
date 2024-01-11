package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/request"
	"encoding/json"
	"time"
)

type Interview struct {
	Common
	Date          time.Time        `json:"date" gorm:"not null;uniqueIndex:interviews_all"`
	Period        constants.Period `json:"period" gorm:"not null;uniqueIndex:interviews_all"` //constants.Period
	Name          constants.Group  `json:"name" gorm:"not null;uniqueIndex:interviews_all"`   //constants.Group
	SlotNumber    int              `json:"slotNumber" gorm:"column:slotNumber;not null"`
	RecruitmentID string           `json:"recruitmentID" gorm:"column:recruitmentId;type:uuid;uniqueIndex:interviews_all"` //manytoone
	Applications  []*Application   `json:"applications,omitempty" gorm:"many2many:interview_selections"`                   //manytomany
}

func (c Interview) TableName() string {
	return "interviews"
}

func GetInterviewsByRidAndName(rid string, name string) (*[]Interview, error) {
	db := global.GetDB()
	var res []Interview
	if err := db.Model(&Interview{}).Preload("Applications").Where("\"recruitmentId\" = ? AND name = ?", rid, name).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
func GetInterviewById(iid string) (*Interview, error) {
	db := global.GetDB()
	var res Interview
	if err := db.Where("uid = ?", iid).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil

}

func UpdateInterview(interview *Interview) error {
	db := global.GetDB()

	return db.Updates(interview).Error
}
func CreateAndSaveInterview(interview *request.UpdateInterview) error {
	var interviewEntity Interview
	bytes, err := json.Marshal(interview)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &interviewEntity)
	db := global.GetDB()
	return db.Create(&interviewEntity).Error
}

// func CreateAndSaveInterview(interviews []request.UpdateInterview) rerror {
// 	var interviewEntitys []Interview
// 	for _, interview := range interviews {
// 		var interviewEntity Interview
// 		bytes, err := json.Marshal(interview)
// 		if err != nil {
// 			return err
// 		}
// 		json.Unmarshal(bytes, &interviewEntity)
// 		interviewEntitys = append(interviewEntitys, interviewEntity)
// 	}
// 	db := global.GetDB()
// 	return db.Create(&interviewEntitys).Error
// }

//	func UpdateInterviews(rid string, name string, interviews []request.UpdateInterview) rerror {
//		var interviewEntitys []Interview
//		for _, interview := range interviews {
//			var interviewEntity Interview
//			bytes, err := json.Marshal(interview)
//			if err != nil {
//				return err
//			}
//			json.Unmarshal(bytes, &interviewEntity)
//			interviewEntity.RecruitmentID = rid
//			interviewEntity.Name = constants.Group(name)
//			interviewEntitys = append(interviewEntitys, interviewEntity)
//		}
//		db := global.GetDB()
//		err := db.Transaction(func(tx *gorm.DB) rerror {
//			for _, interviewEntity := range interviewEntitys {
//				errUpdate := tx.Updates(interviewEntity).Error
//				if errUpdate != nil {
//					return errUpdate
//				}
//			}
//			return nil
//		})
//		return err
//	}
//
//	func DeleteInterviews(name string, interviews []request.DeleteInterviewUID) rerror {
//		db := global.GetDB()
//		err := db.Transaction(func(tx *gorm.DB) rerror {
//			for _, interview := range interviews {
//				if errDelete := tx.Delete(&Interview{}, interview).Error; errDelete != nil {
//					return errDelete
//				}
//			}
//			return nil
//		})
//		return err
//	}
func CreateInterviewsInBatches(interviews []Interview) error {
	db := global.GetDB()
	return db.Create(&interviews).Error
}

func RemoveInterviewByID(iid string) error {
	db := global.GetDB()
	return db.Where("uid = ?", iid).Delete(&Interview{}).Error
}
