package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func CreateComment(opts *pkg.CreateCommentOpts) (*pkg.Comment, error) {
	db := global.GetDB()
	c := &pkg.Comment{
		ApplicationID: opts.ApplicationID,
		MemberName:    opts.MemberName,
		MemberID:      opts.MemberID,
		Content:       opts.Content,
		Evaluation:    opts.Evaluation,
	}
	err := db.Create(c).Error
	return c, err
}

func DeleteCommentById(cid string) error {
	db := global.GetDB()
	return db.Delete(&pkg.Comment{}, "uid = ?", cid).Error
}

func GetCommentById(cid string) (*pkg.Comment, error) {
	db := global.GetDB()
	var c pkg.Comment
	if err := db.Model(&pkg.Comment{}).
		Where("uid = ?", cid).
		First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
