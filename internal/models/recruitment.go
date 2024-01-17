package models

import (
	"encoding/json"

	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func CreateRecruitment(req *pkg.CreateRecOpts) (string, error) {
	db := global.GetDB()
	r := &pkg.Recruitment{
		Name:      req.Name,
		Beginning: req.Beginning,
		Deadline:  req.Deadline,
		End:       req.End,
	}
	err := db.Model(&pkg.Recruitment{}).Create(r).Error
	return r.Uid, err
}

func UpdateRecruitment(req *pkg.UpdateRecOpts) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	var r pkg.Recruitment
	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}
	r.Uid = req.Rid

	db := global.GetDB()
	return db.Updates(&r).Error
}

func GetRecruitmentById(rid string) (*pkg.Recruitment, error) {
	db := global.GetDB()
	var r pkg.Recruitment
	if err := db.Model(&pkg.Recruitment{}).
		Where("uid = ?", rid).
		Find(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func GetFullRecruitmentById(rid string) (*pkg.Recruitment, error) {
	db := global.GetDB()
	var r pkg.Recruitment
	//remember preload need the struct filed name
	var err error
	if err = db.Model(&pkg.Recruitment{}).
		Preload("Applications").
		Preload("Interviews").
		Where("uid = ?", rid).Find(&r).Error; err != nil {
		err = db.Model(&pkg.Recruitment{}).Where("uid = ?", rid).Find(&r).Error
	}
	return &r, err
}

func GetAllRecruitment() ([]pkg.Recruitment, error) {
	db := global.GetDB()
	var r []pkg.Recruitment
	err := db.Model(&pkg.Recruitment{}).
		Order("beginning DESC").
		Find(&r).Error
	return r, err
}

// GetPendingRecruitment get the latest recruitment
func GetPendingRecruitment() (*pkg.Recruitment, error) {
	db := global.GetDB()
	var r pkg.Recruitment
	//if err := db.Model(&Recruitment{}).
	//	Select("uid").
	//	Where("? BETWEEN \"beginning\" AND \"end\"", time.Now()).
	//	First(&r).Error; err != nil {
	//	return nil, err
	//}
	if err := db.Model(&pkg.Recruitment{}).
		Select("uid").
		Order("beginning DESC").
		Limit(1).
		Find(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}
