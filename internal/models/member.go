package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"
)

// not use

type MemberEntity struct {
	Common
	Name string `gorm:"not null"`

	Phone        string           `gorm:"not null;unique;unique"`
	Mail         string           `gorm:"unique"`
	Gender       constants.Gender `gorm:"not null"` //constants.Gender
	WeChatID     string           `gorm:"column:weChatID;not null;unique"`
	JoinTime     string           `gorm:"column:joinTime;not null"`
	IsCaptain    bool             `gorm:"column:isCaptain;not null;default:false"`
	IsAdmin      bool             `gorm:"column:isAdmin;not null;default:false"`
	Group        constants.Group  `gorm:"not null"` //constants.Group
	Avatar       string
	Comments     []Comment `gorm:"foreignKey:MemberID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //onetomany
	PasswordSalt string    `gorm:"column:passwordSalt;not null"`
	PasswordHash string    `gorm:"column:passwordHash;not null;unique"`
}

func (c MemberEntity) TableName() string {
	return "members"
}

type MemberOut struct {
	Common
	Name string `gorm:"not null"`

	Phone     string           `gorm:"not null;unique;unique"`
	Mail      string           `gorm:"unique"`
	Gender    constants.Gender `gorm:"not null"` //constants.Gender
	JoinTime  string           `gorm:"column:joinTime;not null"`
	IsCaptain bool             `gorm:"column:isCaptain;not null;default:false"`
	IsAdmin   bool             `gorm:"column:isAdmin;not null;default:false"`
	Group     constants.Group  `gorm:"not null"` //constants.Group
}

func GetMemberById(mid string) (*MemberEntity, error) {
	db := global.GetDB()
	var m MemberEntity
	err := db.Where("uid = ?", mid).Find(&m).Error
	return &m, err
}
