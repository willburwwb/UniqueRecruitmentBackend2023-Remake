package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/request"
)

type Evaluation int

const (
	Good Evaluation = iota
	Normal
	Bad
)

type CommentEntity struct {
	Common
	ApplicationID string     `gorm:"column:applicationId;type:uuid;"` //manytoone
	MemberID      string     `gorm:"column:memberId;type:uuid;"`      //manytoone
	Content       string     `gorm:"not null"`
	Evaluation    Evaluation `gorm:"column:evaluation;type:int;not null"`
}

func (c CommentEntity) TableName() string {
	return "comments"
}

func CreateComment(req *request.CreateCommentRequest) (string, error) {
	db := global.GetDB()
	c := CommentEntity{
		ApplicationID: req.ApplicationID,
		MemberID:      req.MemberID,
		Content:       req.Content,
		Evaluation:    req.Evaluation,
	}
	err := db.Create(&c).Error
	return c.Uid, err
}

func DeleteCommentById(cid string) error {
	db := global.GetDB()
	return db.Delete(&CommentEntity{}, cid).Error
}
