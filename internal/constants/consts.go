package constants

type Group string

const (
	Web     Group = "web"
	Lab     Group = "lab"
	Ai      Group = "ai"
	Game    Group = "game"
	Android Group = "android"
	Ios     Group = "ios"
	Design  Group = "design"
	Pm      Group = "pm"
	Unique  Group = "unique" // for group interview
)

var GroupMap = map[string]Group{
	"web":     "web",
	"lab":     "lab",
	"ai":      "ai",
	"game":    "game",
	"android": "android",
	"ios":     "ios",
	"design":  "design",
	"pm":      "pm",
}

type Period string

const (
	Morning   Period = "morning"
	Afternoon Period = "afternoon"
	Evening   Period = "evening"
)

type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
	Oth    Gender = 3
)

type Grade string

const (
	Freshman  Grade = "freshman"
	Sophomore Grade = "sophomore"
	Junior    Grade = "junior"
	Senior    Grade = "senior"
	Graduate  Grade = "graduate"
)

type Step string

const (
	SignUp               Step = "SignUp"               //报名
	WrittenTest          Step = "WrittenTest"          //笔试
	GroupTimeSelection   Step = "GroupTimeSelection"   //组面时间选择
	GroupInterview       Step = "GroupInterview"       //组面
	OnlineGroupInterview Step = "OnlineGroupInterview" //在线组面
	StressTest           Step = "StressTest"           // 熬测
	TeamTimeSelection    Step = "TeamTimeSelection"    // 面试时间选择
	TeamInterview        Step = "TeamInterview"        // 群面
	OnlineTeamInterview  Step = "OnlineTeamInterview"  // 在线群面
	Pass                 Step = "Pass"                 // 通过
)

var StepMap = map[Step]string{
	"SignUp":               "SignUp",
	"WrittenTest":          "WrittenTest",
	"GroupTimeSelection":   "GroupTimeSelection",
	"GroupInterview":       "GroupInterview",
	"OnlineGroupInterview": "OnlineGroupInterview",
	"StressTest":           "StressTest",
	"TeamTimeSelection":    "TeamTimeSelection",
	"TeamInterview":        "TeamInterview",
	"OnlineTeamInterview":  "OnlineTeamInterview",
	"Pass":                 "Pass",
}

type GroupOrTeam string

const (
	InGroup GroupOrTeam = "group"
	InTeam  GroupOrTeam = "team"
)

type Role string

const (
	Admin         Role = "admin"
	MemberRole    Role = "member"
	CandidateRole Role = "candidate"
)

type SMSType string

const (
	Accept SMSType = "Accept"
	Reject SMSType = "Reject"
)

type SMSTemplateType string

const (
	StateChange          SMSTemplateType = "stateChange"
	VerificationCode     SMSTemplateType = "verificationCode"
	Interview            SMSTemplateType = "interview"
	PassSMS              SMSTemplateType = "pass"
	Delay                SMSTemplateType = "delay"
	OnLineGroupInterview SMSTemplateType = "onlineGroupInterview"
	OnLineTeamInterview  SMSTemplateType = "onlineTeamInterview"
)

var SMSTemplateMap = map[SMSTemplateType]string{
	StateChange:          "1092770",
	VerificationCode:     "1092824",
	Interview:            "1113517",
	PassSMS:              "1092767",
	Delay:                "1092765",
	OnLineGroupInterview: "1533304",
	OnLineTeamInterview:  "1533302",
}
