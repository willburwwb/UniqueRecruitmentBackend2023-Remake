package models

import (
	"gorm.io/gorm"

	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func GetInterviewsByRidAndNameWithoutApp(rid string, name pkg.Group) ([]pkg.Interview, error) {
	db := global.GetDB()
	var res []pkg.Interview
	if err := db.Model(&pkg.Interview{}).
		Omit("\"selectNumber\", \"slotNumber\"").
		Where("\"selectNumber\" < \"slotNumber\"").
		Where("\"recruitmentId\" = ? AND name = ?", rid, name).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetInterviewsByRidAndNameWithoutAppByMember(rid string, name pkg.Group) ([]pkg.Interview, error) {
	db := global.GetDB()
	var res []pkg.Interview
	if err := db.Model(&pkg.Interview{}).
		Where("\"recruitmentId\" = ? AND name = ?", rid, name).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetInterviewsByRidAndName(rid string, name pkg.Group) ([]pkg.Interview, error) {
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

func UpdateInterview(interview *pkg.Interview) error {
	db := global.GetDB()
	if err := db.Model(&pkg.Interview{}).
		Where("\"uid\" = ?", interview.Uid).
		Updates(map[string]interface{}{
			"date":       interview.Date,
			"period":     interview.Period,
			"start":      interview.Start,
			"end":        interview.End,
			"slotNumber": interview.SlotNumber,
			"name":       interview.Name,
		}).Error; err != nil {
		return err
	}
	return nil
}

func AddAndDeleteInterviews(interviewsToAdd []pkg.Interview, interviewIdsToDel []string) (err error) {
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
		return nil
	}); err != nil {
		return
	}
	return
}
