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
	return db.Delete(&pkg.Comment{}, cid).Error
}
