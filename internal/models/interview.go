package models

import (
	"encoding/json"

	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func GetInterviewsByRidAndName(rid string, name string) (*[]pkg.Interview, error) {
	db := global.GetDB()
	var res []pkg.Interview
	if err := db.Model(&pkg.Interview{}).
		Preload("Applications").
		Where("\"recruitmentId\" = ? AND name = ?", rid, name).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
func GetInterviewById(iid string) (*pkg.Interview, error) {
	db := global.GetDB()
	var res pkg.Interview
	if err := db.Where("uid = ?", iid).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil

}

func UpdateInterview(interview *pkg.Interview) error {
	db := global.GetDB()

	return db.Updates(interview).Error
}
func CreateAndSaveInterview(interview *pkg.UpdateInterviewOpts) error {
	var interviewEntity pkg.Interview
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
func CreateInterviewsInBatches(interviews []pkg.Interview) error {
	db := global.GetDB()
	return db.Create(&interviews).Error
}

func RemoveInterviewByID(iid string) error {
	db := global.GetDB()
	return db.Where("uid = ?", iid).Delete(&pkg.Interview{}).Error
}
