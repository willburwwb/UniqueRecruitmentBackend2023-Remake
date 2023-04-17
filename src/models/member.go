package models

import (
	"UniqueRecruitmentBackend/constants"
)

type MemberEntity struct {
	Common
	JoinTime  string
	IsCaptain bool
	IsAdmin   bool
	Group     constants.Group
	Avatar    string
	Comments  []CommentEntity
	Name      string
	Password  Password
	Phone     string
	Mail      string
	Gender    constants.Gender
}
