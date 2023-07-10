package models

type CommentEntity struct {
	Common
	ApplicationID string `gorm:"column:applicationId;type:uuid;"` //manytoone
	MemberID      string `gorm:"column:memberId;type:uuid;"`      //manytoone
	Content       string `gorm:"not null"`
	Evaluation    string `gorm:"not null"`
}

func (c CommentEntity) TableName() string {
	return "comments"
}
