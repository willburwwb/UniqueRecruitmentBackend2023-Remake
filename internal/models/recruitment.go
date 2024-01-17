package models

import (
	"encoding/json"
	"errors"
	"log"

	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func CreateRecruitment(opts *pkg.CreateRecOpts) (r *pkg.Recruitment, err error) {
	db := global.GetDB()
	if db.Model(&pkg.Recruitment{}).
		Where("name = ?", opts.Name).
		Find(r).RowsAffected > 0 {
		return nil, errors.New("recruitment with the same name cannot be created")
	}

	r = &pkg.Recruitment{
		Name:      opts.Name,
		Beginning: opts.Beginning,
		Deadline:  opts.Deadline,
		End:       opts.End,
	}
	log.Println("create ", r)
	err = db.Model(&pkg.Recruitment{}).Create(r).Error
	return
}

func UpdateRecruitment(opts *pkg.UpdateRecOpts) error {
	bytes, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	var r pkg.Recruitment
	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}
	r.Uid = opts.Rid

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
