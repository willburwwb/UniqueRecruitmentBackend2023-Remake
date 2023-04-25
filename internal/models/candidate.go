package models

type CandidateEntity struct {
	Common
	Applications []ApplicationEntity `gorm:"foreignKey:CandidateID;references:Uid;constraint:OnDelete:CASCADE;"` //onetomany
	Name         string              `gorm:"not null"`
	Phone        string              `gorm:"not null;unique"`
	Mail         string              `gorm:"unique"`
	Gender       string              `gorm:"not null"` //constants.Gender
	PasswordSalt string              `gorm:"columns:passwordSalt;not null"`
	PasswordHash string              `gorm:"columns:passwordHash;not null;unique"`
}

func (c CandidateEntity) TableName() string {
	return "candidates"
}
