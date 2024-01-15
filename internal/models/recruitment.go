package models

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/request"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/pgtype"
)

type Recruitment struct {
	Common
	Name       string       `gorm:"not null;unique" json:"name"`
	Beginning  time.Time    `gorm:"not null" json:"beginning"`
	Deadline   time.Time    `gorm:"not null" json:"deadline"`
	End        time.Time    `gorm:"not null" json:"end"`
	Statistics pgtype.JSONB `gorm:"type:jsonb"`

	Applications []Application `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //一个hr->简历 ;级联删除
	Interviews   []Interview   `gorm:"foreignKey:RecruitmentID;references:Uid;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //一个hr->面试 ;级联删除
}

func (r Recruitment) TableName() string {
	return "recruitments"
}

func (r *Recruitment) GetInterviews(name string) []Interview {
	reInterviews := make([]Interview, 0)
	for _, interview := range r.Interviews {
		if string(interview.Name) == name {
			reInterviews = append(reInterviews, interview)
		}
	}
	return reInterviews
}

func CreateRecruitment(req *request.CreateRecruitment) (string, error) {
	db := global.GetDB()
	r := &Recruitment{
		Name:      req.Name,
		Beginning: req.Beginning,
		Deadline:  req.Deadline,
		End:       req.End,
	}
	err := db.Model(&Recruitment{}).Create(r).Error
	return r.Uid, err
}

func UpdateRecruitment(req *request.UpdateRecruitment) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	var r Recruitment
	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}
	r.Uid = req.Rid

	db := global.GetDB()
	return db.Updates(&r).Error
}

func GetRecruitmentById(rid string) (*Recruitment, error) {
	db := global.GetDB()
	var r Recruitment
	if err := db.Model(&Recruitment{}).Where("uid = ?", rid).Find(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func GetFullRecruitmentById(rid string) (*Recruitment, error) {
	db := global.GetDB()
	var r Recruitment
	//remember preload need the struct filed name
	var err error
	if err = db.Model(&Recruitment{}).
		Preload("Applications").
		Preload("Interviews").
		Where("uid = ?", rid).Find(&r).Error; err != nil {
		err = db.Model(&Recruitment{}).Where("uid = ?", rid).Find(&r).Error
	}
	return &r, err
}

func GetAllRecruitment() ([]Recruitment, error) {
	db := global.GetDB()
	var r []Recruitment
	err := db.Model(&Recruitment{}).Order("beginning DESC").Find(&r).Error
	return r, err
}

// GetPendingRecruitment get the latest recruitment
func GetPendingRecruitment() (*Recruitment, error) {
	db := global.GetDB()
	var r Recruitment
	//if err := db.Model(&Recruitment{}).
	//	Select("uid").
	//	Where("? BETWEEN \"beginning\" AND \"end\"", time.Now()).
	//	First(&r).Error; err != nil {
	//	return nil, err
	//}
	if err := db.Model(&Recruitment{}).
		Select("uid").
		Order("beginning DESC").
		Limit(1).
		Find(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}
