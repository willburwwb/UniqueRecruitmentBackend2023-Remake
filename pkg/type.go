package pkg

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"
)

type Common struct {
	Uid       string    `gorm:"column:uid;type:uuid;default:gen_random_uuid();primaryKey" json:"uid"`
	CreatedAt time.Time `gorm:"column:createdAt;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;index" json:"updated_at"`
}

type UserDetail struct {
	UID         string   `json:"uid"`
	Phone       string   `json:"phone"`
	Email       string   `json:"email"`
	Password    string   `json:"password,omitempty"`
	Name        string   `json:"name"`
	AvatarURL   string   `json:"avatar_url"`
	Gender      Gender   `json:"gender"`
	JoinTime    string   `json:"join_time"`
	Groups      []string `json:"groups"`
	LarkUnionID string   `json:"lark_union_id"`
}

type UserDetailResp struct {
	UserDetail
	Applications []Application `json:"applications"`
}

type Recruitment struct {
	Common
	Name       string    `gorm:"not null;unique" json:"name"`
	Beginning  time.Time `gorm:"not null" json:"beginning"`
	Deadline   time.Time `gorm:"not null" json:"deadline"`
	End        time.Time `gorm:"not null" json:"end"`
	Statistics string    `gorm:"not null" json:"statistics"`

	Applications []Application `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"applications"` //一个hr->简历 ;级联删除
	Interviews   []Interview   `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"interviews"`   //一个hr->面试 ;级联删除
}

func (r Recruitment) TableName() string {
	return "recruitments"
}

func (r Recruitment) GetInterviews(name string) []Interview {
	reInterviews := make([]Interview, 0)
	for _, interview := range r.Interviews {
		if interview.Name == name {
			reInterviews = append(reInterviews, interview)
		}
	}
	return reInterviews
}

type CreateRecOpts struct {
	Name      string    `json:"name" binding:"required"`
	Beginning time.Time `json:"beginning" binding:"required"`
	Deadline  time.Time `json:"deadline" binding:"required"`
	End       time.Time `json:"end" binding:"required"`
}

func (r *CreateRecOpts) Validate() error {
	if r.Beginning.After(r.Deadline) || r.Deadline.After(r.End) {
		return errors.New("time set up wrong")
	}
	return nil
}

type UpdateRecOpts struct {
	Rid       string    `json:"rid"`
	Name      string    `json:"name"`
	Beginning time.Time `json:"beginning"`
	Deadline  time.Time `json:"deadline"`
	End       time.Time `json:"end"`
}

func (r *UpdateRecOpts) Validate() error {
	if r.Rid == "" {
		return errors.New("recruitment id is null")
	}
	return nil
}

type GetRecOpts struct {
	Rid string `uri:"rid" binding:"required"`
}

type InterviewInfo struct {
	Id         string    `json:"id"`
	Date       time.Time `json:"date"`
	Period     Period    `json:"period"`
	SlotNumber int       `json:"slot_number"`
}

type SetRecInterviewTimeOpts struct {
	Interviews []InterviewInfo
}

// Application records the detail of application for candidate
// uniqueIndex(CandidateID,RecruitmentID)
type Application struct {
	Common
	Grade                     string    `gorm:"not null" json:"grade"` //pkg.Grade
	Institute                 string    `gorm:"not null" json:"institute"`
	Major                     string    `gorm:"not null" json:"major"`
	Rank                      string    `gorm:"not null" json:"rank"`
	Group                     string    `gorm:"not null" json:"group"` //pkg.Group
	Intro                     string    `gorm:"not null" json:"intro"`
	IsQuick                   bool      `gorm:"column:isQuick;not null" json:"is_quick"`
	Referrer                  string    `json:"referrer"`
	Resume                    string    `json:"resume"`
	Abandoned                 bool      `gorm:"not null; default false" json:"abandoned"`
	Rejected                  bool      `gorm:"not null; default false" json:"rejected"`
	Step                      string    `gorm:"not null" json:"step"`                                                                          //pkg.Step
	CandidateID               string    `gorm:"column:candidateId;type:uuid;uniqueIndex:UQ_CandidateID_RecruitmentID" json:"candidate_id"`     //manytoone
	RecruitmentID             string    `gorm:"column:recruitmentId;type:uuid;uniqueIndex:UQ_CandidateID_RecruitmentID" json:"recruitment_id"` //manytoone
	InterviewAllocationsGroup time.Time `gorm:"column:interviewAllocationsGroup;" json:"interview_allocations_group"`
	InterviewAllocationsTeam  time.Time `gorm:"column:interviewAllocationsTeam;" json:"interview_allocations_team"`

	UserDetail          *UserDetail `gorm:"-" json:"user_detail"`                                                                                     // get from sso
	InterviewSelections []Interview `gorm:"many2many:interview_selections;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"interview_selections"` //manytomany
	Comments            []Comment   `gorm:"foreignKey:ApplicationID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"comments"`    //onetomany
}

func (a Application) TableName() string {
	return "applications"
}

type CreateAppOpts struct {
	Grade         string `form:"grade" json:"grade" binding:"required"`
	Institute     string `form:"institute" json:"institute" binding:"required"`
	Major         string `form:"major" json:"major" binding:"required"`
	Rank          string `form:"rank" json:"rank" binding:"required"`
	Group         string `form:"group" json:"group" binding:"required"`
	Intro         string `form:"intro" json:"intro" binding:"required"` //自我介绍
	RecruitmentID string `form:"recruitment_id" json:"recruitment_id" binding:"required"`
	Referrer      string `form:"referrer" json:"referrer"` //推荐人
	IsQuick       bool   `form:"is_quick" json:"is_quick"` //速通

	Resume *multipart.FileHeader `form:"resume" json:"resume"` //简历
}

func (opts *CreateAppOpts) Validate() (err error) {
	if GroupMap[opts.Group] == "" {
		return errors.New("request body error, group set wrong")
	}
	return
}

type UpdateAppOpts struct {
	Aid string

	Grade     string `form:"grade" json:"grade,omitempty"`
	Institute string `form:"institute" json:"institute,omitempty"`
	Major     string `form:"major" json:"major,omitempty"`
	Rank      string `form:"rank" json:"rank,omitempty"`
	Group     string `form:"group" json:"group,omitempty"`
	Intro     string `form:"intro" json:"intro,omitempty"`       //自我介绍
	Referrer  string `form:"referrer" json:"referrer,omitempty"` //推荐人
	IsQuick   *bool  `form:"is_quick" json:"is_quick"`           //速通

	Resume *multipart.FileHeader `form:"resume" json:"resume,omitempty"` //简历
}

func (opts *UpdateAppOpts) Validate() (err error) {
	if opts.Group != "" && GroupMap[opts.Group] == "" {
		return errors.New("request body error, group set wrong")
	}
	if opts.Aid == "" {
		return errors.New("request body error, application id is nil")
	}
	return
}

type SetAppStepOpts struct {
	Aid string

	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
}

func (opts *SetAppStepOpts) Validate() (err error) {
	fromRank, ok := StepRanks[Step(opts.From)]
	if !ok {
		return fmt.Errorf("request body error, from step %s set wrong", opts.From)
	}

	toRank, ok := StepRanks[Step(opts.To)]
	if !ok {
		return fmt.Errorf("request body error, to step %s set wrong", opts.To)
	}

	if fromRank >= toRank {
		return fmt.Errorf("request body error, %s should be set after %s", opts.From, opts.To)
	}

	if opts.Aid == "" {
		return errors.New("request body error, application id is nil")
	}
	return
}

type SetAppInterviewTimeOpts struct {
	Aid           string
	InterviewType string

	Time time.Time `json:"time" binding:"required"`
}

func (opts *SetAppInterviewTimeOpts) Validate() (err error) {
	if opts.InterviewType != "group" && opts.InterviewType != "team" {
		return fmt.Errorf("request param rerror, type should be group/team")
	}
	if opts.Aid == "" {
		return errors.New("request param error, application id is nil")
	}
	return
}

type SelectInterviewSlotsOpts struct {
	Aid           string
	InterviewType string

	Iids []string `json:"iids" binding:"required"`
}

func (opts *SelectInterviewSlotsOpts) Validate() (err error) {
	if opts.InterviewType != "group" && opts.InterviewType != "team" {
		return fmt.Errorf("request param rerror, type should be group/team")
	}
	if opts.Aid == "" {
		return errors.New("request param error, application id is nil")
	}
	if len(opts.Iids) == 0 {
		return errors.New("request body error, len of interview ids is 0")
	}
	return
}

type Interview struct {
	Common
	Date          time.Time     `json:"date" gorm:"not null;uniqueIndex:interviews_all"`
	Period        Period        `json:"period" gorm:"not null;uniqueIndex:interviews_all"` //pkg.Period
	Name          string        `json:"name" gorm:"not null;uniqueIndex:interviews_all"`   //pkg.Group
	SlotNumber    int           `json:"slot_number" gorm:"column:slotNumber;not null"`
	RecruitmentID string        `json:"recruitment_id" gorm:"not null;column:recruitmentId;type:uuid;uniqueIndex:interviews_all"` //manytoone
	Applications  []Application `json:"applications,omitempty" gorm:"many2many:interview_selections"`                             //manytomany
}

func (c Interview) TableName() string {
	return "interviews"
}

type UpdateInterviewOpts struct {
	Uid        string    `json:"uid" form:"uid"`
	Date       time.Time `json:"date" form:"date" binding:"required"`
	Period     Period    `json:"period" form:"period" binding:"required" `
	SlotNumber int       `json:"slot_number" form:"slot_number" binding:"required"`
}

type DeleteInterviewUID string

type Evaluation int

const (
	Good Evaluation = iota
	Normal
	Bad
)

type Comment struct {
	Common
	ApplicationID string     `gorm:"column:applicationId;type:uuid;" json:"application_id"` //manytoone
	MemberID      string     `gorm:"column:memberId;type:uuid;index" json:"member_id"`      //manytoone
	Content       string     `gorm:"column:content;not null" json:"content"`
	Evaluation    Evaluation `gorm:"column:evaluation;type:int;not null" json:"evaluation"`
}

func (c Comment) TableName() string {
	return "comments"
}

type CreateCommentOpts struct {
	MemberID string `json:"member_id"`

	ApplicationID string `json:"application_id" binding:"required"`
	Content       string `json:"content" binding:"required"`
	Evaluation    int    `json:"evaluation"`
}

type SendSMSOpts struct {
	Type      SMSType  `json:"type" binding:"required"`    // the candidate status : Pass or Fail
	Current   Step     `json:"current" binding:"required"` // the application current step
	Next      Step     `json:"next" binding:"required"`    // the application next step
	Time      string   `json:"time"`                       // the next step(interview/test) time
	Place     string   `json:"place"`                      // the next step(interview/test) place
	MeetingId string   `json:"meeting_id"`
	Rest      string   `json:"rest"`
	Aids      []string `json:"aids"` // the applications will be sent sms
}

func (opts *SendSMSOpts) Validate() (err error) {
	if opts.Type != Accept && opts.Type != Reject {
		err = fmt.Errorf("sms type is invalid")
		return
	}
	if _, ok := ZhToEnStepMap[string(opts.Next)]; ok {
		opts.Next = ZhToEnStepMap[string(opts.Next)]
	}
	if _, ok := ZhToEnStepMap[string(opts.Current)]; ok {
		opts.Current = ZhToEnStepMap[string(opts.Current)]
	}
	if len(opts.Aids) == 0 {
		err = fmt.Errorf("request body error, aids is nil")
		return
	}
	if _, ok := EnToZhStepMap[opts.Next]; !ok {
		err = fmt.Errorf("request body error, next is invalid")
		return
	}
	if _, ok := EnToZhStepMap[opts.Current]; !ok {
		err = fmt.Errorf("request body error, current is invalid")
		return
	}
	return
}
