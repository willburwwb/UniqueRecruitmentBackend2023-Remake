package global

import (
	"UniqueRecruitmentBackend/configs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db.Session(&gorm.Session{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
}

func setupPgsql() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		configs.Config.Pgsql.Host, configs.Config.Pgsql.User, configs.Config.Pgsql.Dbname, configs.Config.Pgsql.Port, configs.Config.Pgsql.Password)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("connect to db rerror, %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("get db rerror, %v", err))
	}
	sqlDB.SetMaxIdleConns(configs.Config.Pgsql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(configs.Config.Pgsql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.Config.Pgsql.MaxLifeSeconds) * time.Second)
}
