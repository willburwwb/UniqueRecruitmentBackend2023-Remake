package models

// TODO(wwb)
// fix memberEntity definition

type MemberEntity struct {
	Common
	Name string `gorm:"not null"`

	Phone        string `gorm:"not null;unique;unique"`
	Mail         string `gorm:"unique"`
	Gender       string `gorm:"not null"` //constants.Gender
	WeChatID     string `gorm:"column:weChatID;not null;unique"`
	JoinTime     string `gorm:"column:joinTime;not null"`
	IsCaptain    bool   `gorm:"column:isCaptain;not null;default:false"`
	IsAdmin      bool   `gorm:"column:isAdmin;not null;default:false"`
	Group        string `gorm:"not null"` //constants.Group
	Avatar       string
	Comments     []CommentEntity `gorm:"foreignKey:MemberID;references:Uid;constraint:OnDelete:CASCADE，OnUpdate:CASCADE;"` //onetomany
	PasswordSalt string          `gorm:"column:passwordSalt;not null"`
	PasswordHash string          `gorm:"column:passwordHash;not null;unique"`
}

func (c MemberEntity) TableName() string {
	return "members"
}
