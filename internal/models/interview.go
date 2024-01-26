package models

import (
	"gorm.io/gorm"

	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func GetInterviewsByRidAndNameWithoutApp(rid string, name string) ([]pkg.Interview, error) {
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

func GetInterviewsByIds(ids []string) ([]pkg.Interview, error) {
	db := global.GetDB()
	var interviews []pkg.Interview
	if err := db.Where("uid in ?", ids).
		Find(&interviews).Error; err != nil {
		return nil, err
	}
	return interviews, nil
}

func UpdateInterview(interviewsToAdd []pkg.Interview, interviewIdsToDel []string, interviewsToUpdate map[string]pkg.Interview) (err error) {
	db := global.GetDB()
	if err = db.Transaction(func(tx *gorm.DB) error {
		if len(interviewsToAdd) != 0 {
			if errCreate := tx.Create(interviewsToAdd).Error; errCreate != nil {
				return errCreate
			}
		}
		if len(interviewIdsToDel) != 0 {
			if errDelete := tx.Delete(&pkg.Interview{}, "uid in ?", interviewIdsToDel).Error; errDelete != nil {
				return errDelete
			}
		}

		for uid, interviewToUpdate := range interviewsToUpdate {
			if err := tx.Model(&pkg.Interview{}).Where("uid = ?", uid).Updates(map[string]interface{}{
				"date":       interviewToUpdate.Date,
				"period":     interviewToUpdate.Period,
				"slotNumber": interviewToUpdate.SlotNumber,
				"name":       interviewToUpdate.Name,
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
