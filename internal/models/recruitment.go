package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"
	"UniqueRecruitmentBackend/internal/request"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/pgtype"
)

type RecruitmentEntity struct {
	Common
	Name       string       `gorm:"not null;unique" json:"name"`
	Beginning  time.Time    `gorm:"not null" json:"beginning"`
	Deadline   time.Time    `gorm:"not null" json:"deadline"`
	End        time.Time    `gorm:"not null" json:"end"`
	Statistics pgtype.JSONB `gorm:"type:jsonb"`

	Applications []ApplicationEntity `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //一个hr->简历 ;级联删除
	Interviews   []InterviewEntity   `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //一个hr->面试 ;级联删除
}

func (c RecruitmentEntity) TableName() string {
	return "recruitments"
}

func CreateRecruitment(req *request.CreateRecruitment) (string, error) {
	db := global.GetDB()
	r := &RecruitmentEntity{
		Name:      req.Name,
		Beginning: req.Beginning,
		Deadline:  req.Deadline,
		End:       req.End,
	}
	err := db.Model(&RecruitmentEntity{}).Create(r).Error
	return r.Uid, err
}

func UpdateRecruitment(rid string, req *request.UpdateRecruitment) error {
	
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	var r RecruitmentEntity
	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}
	r.Uid = rid
	
	db := global.GetDB()
	return db.Updates(&r).Error
}

func GetRecruitmentById(rid string, role constants.Role) (*RecruitmentEntity, error) {
	db := global.GetDB()
	var r RecruitmentEntity
	//remember preload need the struct filed name
	var err error
	if role == constants.MemberRole || role == constants.Admin {
		err = db.Model(&RecruitmentEntity{}).Preload("Applications").Preload("Interviews").
			Where("uid = ?", rid).Find(&r).Error
	} else {
		err = db.Model(&RecruitmentEntity{}).Where("uid = ?", rid).Find(&r).Error
	}
	return &r, err
}

func GetAllRecruitment() ([]RecruitmentEntity, error) {
	db := global.GetDB()
	var r []RecruitmentEntity
	err := db.Model(&RecruitmentEntity{}).Order("beginning DESC").Find(&r).Error
	return r, err
}

func GetPendingRecruitment(role constants.Role) (*RecruitmentEntity, error) {
	db := global.GetDB()
	var r RecruitmentEntity
	var err error
	if role == constants.MemberRole || role == constants.Admin {
		err = db.Model(&RecruitmentEntity{}).Preload("Applications").Preload("Interviews").
			Where("? BETWEEN beginning AND end", time.Now()).Find(&r).Error
	} else {
		err = db.Model(&RecruitmentEntity{}).Where("? BETWEEN beginning AND end", time.Now()).Find(&r).Error
	}
	return &r, err
}
