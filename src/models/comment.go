package models

type CommentEntity struct {
	Common
	Application ApplicationEntity
	Member      MemberEntity
	Content     string
	Evaluation  string
}
