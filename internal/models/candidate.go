package models

// not use

type CandidateEntity struct {
	Common
	Applications []Application `gorm:"foreignKey:CandidateID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //onetomany
	Name         string        `gorm:"not null"`
	Phone        string        `gorm:"not null;unique"`
	Mail         string        `gorm:"unique"`
	Gender       string        `gorm:"not null"` //constants.Gender
	PasswordSalt string        `gorm:"column:passwordSalt;not null"`
	PasswordHash string        `gorm:"column:passwordHash;not null;unique"`
}

func (c CandidateEntity) TableName() string {
	return "candidates"
}
