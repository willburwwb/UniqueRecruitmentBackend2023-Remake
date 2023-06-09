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

type Period string

const (
	Morning   Period = "morning"
	Afternoon Period = "afternoon"
	Evening   Period = "evening"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Oth    Gender = "oth"
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
type GroupOrTeam string

const (
	InGroup GroupOrTeam = "group"
	InTeam  GroupOrTeam = "team"
)

type Role uint

const (
	Admin         Role = 1
	MemberRole    Role = 2
	CandidateRole Role = 3
)
