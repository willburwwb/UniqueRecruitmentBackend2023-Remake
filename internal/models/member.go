package models

import (
	"UniqueRecruitmentBackend/internal/constants"
)

type MemberEntity struct {
	Common
	JoinTime  string
	IsCaptain bool
	IsAdmin   bool
	Group     constants.Group
	Avatar    string
	//Comments  []CommentEntity //onetomany

	Name     string
	Password string
	Phone    string
	Mail     string
	Gender   constants.Gender
}
