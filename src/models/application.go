package models

import (
	"UniqueRecruitmentBackend/constants"
)

type ApplicationEntity struct {
	Common
	Grade                constants.Grade
	Institute            string
	Major                string
	Rank                 string
	Group                constants.Group
	Intro                string
	IsQuick              bool
	Referrer             string
	Resume               string
	Abandoned            bool
	Rejected             bool
	Step                 string
	InterviewSelections  []InterviewEntity
	InterviewAllocations []InterviewAllocation
	Candidate            CandidateEntity
	Recruitment          RecruitmentEntity
	Comments             []CommentEntity
}
