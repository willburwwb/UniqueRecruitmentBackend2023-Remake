package models

type CommentEntity struct {
	Common
	ApplicationID string `gorm:"columns:applicationId"` //manytoone
	MemberID      string `gorm:"columns:memberId"`      //manytoone
	Content       string `gorm:"not null"`
	Evaluation    string `gorm:"not null"`
}

func (c CommentEntity) TableName() string {
	return "comments"
}
