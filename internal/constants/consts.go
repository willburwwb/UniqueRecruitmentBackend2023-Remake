package constants

type Group string

const (
	Web     Group = "web"
	lab     Group = "lab"
	ai      Group = "ai"
	game    Group = "game"
	android Group = "android"
	ios     Group = "ios"
	design  Group = "design"
	pm      Group = "pm"
	unique  Group = "unique"
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
const ()
