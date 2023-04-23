package models

type CommentEntity struct {
	Common
	ApplicationID uint
	//Application ApplicationEntity
	//Member      MemberEntity?为什么member有comment
	Content    string
	Evaluation string
}
