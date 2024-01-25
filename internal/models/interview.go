package models

import (
	"gorm.io/gorm"

	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func GetInterviewsByRidAndName(rid string, name string) ([]pkg.Interview, error) {
	db := global.GetDB()
	var res []pkg.Interview
	if err := db.Model(&pkg.Interview{}).
		Preload("Applications").
		Where("\"recruitmentId\" = ? AND name = ?", rid, name).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
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

func UpdateInterview(interviewsToAdd []pkg.Interview, interviewIdsToDel []string, interviewsToUpdate map[string]pkg.Interview) (err error) {
	db := global.GetDB()
	if err = db.Transaction(func(tx *gorm.DB) error {
		if errCreate := tx.Create(interviewsToAdd).Error; errCreate != nil {
			return errCreate
		}

		if errDelete := tx.Delete(&pkg.Interview{}, "uid in ?", interviewIdsToDel).Error; errDelete != nil {
			return errDelete
		}

		for uid, interviewToUpdate := range interviewsToUpdate {
			if err := tx.Where("uid = ?", uid).Updates(map[string]interface{}{
				"date":        interviewToUpdate.Date,
				"period":      interviewToUpdate.Period,
				"slot_number": interviewToUpdate.SlotNumber,
				"name":        interviewToUpdate.Name,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return
	}
	return
}
