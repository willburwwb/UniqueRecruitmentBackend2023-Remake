package models

import (
	"UniqueRecruitmentBackend/global"
)

func SetupTables() {
	db := global.GetDB()
	db.AutoMigrate(&Common{})
	db.AutoMigrate(&RecruitmentEntity{})
	db.AutoMigrate(&CandidateEntity{})
	db.AutoMigrate(&MemberEntity{})
	db.AutoMigrate(&ApplicationEntity{})
	db.AutoMigrate(&InterviewEntity{})

	db.AutoMigrate(&CommentEntity{})
}
