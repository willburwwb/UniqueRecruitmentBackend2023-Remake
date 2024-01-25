package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/pkg"
)

func CreateComment(req *pkg.CreateCommentOpts) (*pkg.Comment, error) {
	db := global.GetDB()
	c := &pkg.Comment{
		ApplicationID: req.ApplicationID,
		MemberID:      req.MemberID,
		Content:       req.Content,
		Evaluation:    pkg.Evaluation(req.Evaluation),
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
