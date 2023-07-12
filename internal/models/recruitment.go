package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/request"
	"time"

	"github.com/jackc/pgx/pgtype"
)

type RecruitmentEntity struct {
	Common
	Name       string       `gorm:"not null;unique"`
	Beginning  time.Time    `gorm:"not null"`
	Deadline   time.Time    `gorm:"not null"`
	End        time.Time    `gorm:"not null"`
	Statistics pgtype.JSONB `gorm:"type:jsonb"`

	Applications []ApplicationEntity `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //一个hr->简历 ;级联删除
	Interviews   []InterviewEntity   `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //一个hr->面试 ;级联删除
}

func (c RecruitmentEntity) TableName() string {
	return "recruitments"
}

func CreateRecruitment(r *request.CreateRecruitmentRequest) (string, error) {
	db := global.GetDB()
	ri := RecruitmentEntity{
		Name:      r.Name,
		Beginning: r.Beginning,
		Deadline:  r.Deadline,
		End:       r.End,
	}
	err := db.Model(&RecruitmentEntity{}).Create(&ri).Error
	return ri.Uid, err
}

func UpdateRecruitment(rid string, r *request.UpdateRecruitmentRequest) error {
	db := global.GetDB()
	return db.Model(&RecruitmentEntity{}).Where("uid = ?", rid).Updates(&RecruitmentEntity{
		Beginning: r.Beginning,
		Deadline:  r.Deadline,
		End:       r.End,
	}).Error
}

func GetRecruitmentById(rid string) (*RecruitmentEntity, error) {
	db := global.GetDB()
	var r RecruitmentEntity
	//remember preload need the struct filed name
	err := db.Model(&RecruitmentEntity{}).Preload("Applications").Preload("Interviews").Where("uid = ?", rid).Find(&r).Error
	return &r, err
}

func GetAllRecruitment() ([]RecruitmentEntity, error) {
	db := global.GetDB()
	var r []RecruitmentEntity
	err := db.Model(&RecruitmentEntity{}).Preload("Interviews").Find(&r).Error
	return r, err
}
