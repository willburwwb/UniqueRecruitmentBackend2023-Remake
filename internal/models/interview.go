package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func GetInterviewById(iid string) (*pkg.Interview, error) {
	db := global.GetDB()
	var interview pkg.Interview
	if err := db.Model(&pkg.Interview{}).
		Where("uid = ?", iid).
		First(&interview).Error; err != nil {
		return nil, err
	}
	return &interview, nil
}

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

func GetInterviewsByIdsAndName(ids []string, name pkg.Group) ([]pkg.Interview, error) {
	db := global.GetDB()
	var interviews []pkg.Interview
	if err := db.Where("uid in ?", ids).
		Where("name = ?", name).
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
	if len(interviewsToAdd) != 0 {
		if errCreate := db.Create(interviewsToAdd).Error; errCreate != nil {
			return errCreate
		}
	}
	if len(interviewIdsToDel) != 0 {
		if errDelete := db.Delete(&pkg.Interview{}, "uid in ?", interviewIdsToDel).Error; errDelete != nil {
			return errDelete
		}
	}
	return
}

func GetInterviewsCannotBeUpdate(iids []string) (map[string]struct{}, error) {
	db := global.GetDB()
	interviewsCannotBeUpdate := make(map[string]struct{})
	res := []string{}

	if len(iids) == 0 {
		return interviewsCannotBeUpdate, nil
	}

	// get the interview uid that has been selected by the application
	if err := db.Table("interview_selections").
		Select("DISTINCT interview_uid").
		Where("interview_uid IN ?", iids).
		Find(&res).Error; err != nil {
		return nil, err
	}
	for _, val := range res {
		interviewsCannotBeUpdate[val] = struct{}{}
	}

	// get the interview uid that has been allocated by the application
	if err := db.Model(&pkg.Application{}).
		Select("DISTINCT \"interviewAllocationsGroupId\"").
		Where("\"interviewAllocationsGroupId\" IN ?", iids).
		Find(&res).Error; err != nil {
		return nil, err
	}
	for _, val := range res {
		interviewsCannotBeUpdate[val] = struct{}{}
	}

	if err := db.Model(&pkg.Application{}).
		Select("DISTINCT \"interviewAllocationsTeamId\"").
		Where("\"interviewAllocationsTeamId\" IN ?", iids).
		Find(&res).Error; err != nil {
		return nil, err
	}
	for _, val := range res {
		interviewsCannotBeUpdate[val] = struct{}{}
	}
	return interviewsCannotBeUpdate, nil
}
