package models

import (
	"UniqueRecruitmentBackend/global"
)

func SetupTables() {
	db := global.Db
	db.AutoMigrate(&Common{})
	db.AutoMigrate(&RecruitmentEntity{})
	db.AutoMigrate(&CandidateEntity{})
	db.AutoMigrate(&MemberEntity{})
	db.AutoMigrate(&CommentEntity{})
	db.AutoMigrate(&ApplicationEntity{})
	db.AutoMigrate(&InterviewEntity{})
}
