package pkg

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

var StepRanks = map[Step]int{
	"SignUp":               1,
	"WrittenTest":          2,
	"GroupTimeSelection":   3,
	"GroupInterview":       4,
	"OnlineGroupInterview": 4,
	"StressTest":           5,
	"TeamTimeSelection":    6,
	"TeamInterview":        7,
	"OnlineTeamInterview":  7,
	"Pass":                 8,
}
var EnToZhStepMap = map[string]string{
	"SignUp":               "报名",
	"WrittenTest":          "笔试",
	"GroupTimeSelection":   "组面时间选择",
	"GroupInterview":       "组面",
	"OnlineGroupInterview": "在线组面",
	"StressTest":           "熬测",
	"TeamTimeSelection":    "群面时间选择",
	"TeamInterview":        "群面",
	"OnlineTeamInterview":  "在线群面",
	"Pass":                 "通过",
}

var ZhToEnStepMap = map[string]string{
	"报名":     "SignUp",
	"笔试":     "WrittenTest",
	"组面时间选择": "GroupTimeSelection",
	"组面":     "GroupInterview",
	"在线组面":   "OnlineGroupInterview",
	"熬测":     "StressTest",
	"群面时间选择": "TeamTimeSelection",
	"群面":     "TeamInterview",
	"在线群面":   "OnlineTeamInterview",
	"通过":     "Pass",
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
	StateChange             SMSTemplateType = "stateChange"
	VerificationCode        SMSTemplateType = "verificationCode"
	Interviews              SMSTemplateType = "interviews"
	PassSMS                 SMSTemplateType = "pass"
	Delay                   SMSTemplateType = "delay"
	OnLineGroupInterviewSMS SMSTemplateType = "onlineGroupInterview"
	OnLineTeamInterviewSMS  SMSTemplateType = "onlineTeamInterview"
)

var SMSTemplateMap = map[SMSTemplateType]uint{
	StateChange:             1092770,
	VerificationCode:        1092824,
	Interviews:              1113517,
	PassSMS:                 1092767,
	Delay:                   1092765,
	OnLineGroupInterviewSMS: 1533304,
	OnLineTeamInterviewSMS:  1533302,
}

const (
	SessionNameUID       = "X-UniqueSSO-UID"
	SessionMaxAgeSeconds = 4 * 60 * 60 // seconds

	SSORedirectBackQueryKey = "unique_sso_redirect_back_uri"
)
