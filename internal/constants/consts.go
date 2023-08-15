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
	SignUp             Step = "SignUp"             //报名
	WrittenTest        Step = "WrittenTest"        //笔试
	GroupTimeSelection Step = "GroupTimeSelection" //组面时间选择
	GroupInterview     Step = "GroupInterview"     //组面
	StressTest         Step = "StressTest"         // 熬测
	TeamTimeSelection  Step = "TeamTimeSelection"  // 面试时间选择
	TeamInterview      Step = "TeamInterview"      // 群面
	Pass               Step = "Pass"               // 通过
)

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
