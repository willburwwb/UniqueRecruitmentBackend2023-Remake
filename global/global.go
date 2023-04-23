package global

import (
	"UniqueRecruitmentBackend/configs"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	sredis "github.com/gin-contrib/sessions/redis"
	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Rdb       *redis.Client
	SessStore sessions.Store
)
var (
	PgsqlConfig  *configs.PsqlConfigs
	RedisConfig  *configs.RedisConfigs
	SessConfig   *configs.SessConfigs
	ServerConfig *configs.ServerConfigs
)

func Setup() error {
	_, err := setupPgsql()
	if err != nil {
		return fmt.Errorf("pgsql setup failed %s", err)
	}
	log.Println("psql setup succeed")
	_, err = setupRedis()
	if err != nil {
		return fmt.Errorf("rdb setup failed %s", err)
	}
	log.Println("redus setup succeed")
	_, err = setupSession()
	if err != nil {
		return fmt.Errorf("session setup failed %s", err)
	}
	log.Println("session setup succeed")
	return nil
}
func setupPgsql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		PgsqlConfig.Host, PgsqlConfig.User, PgsqlConfig.Password, PgsqlConfig.Dbname, PgsqlConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(PgsqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(PgsqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(PgsqlConfig.MaxLifeSeconds) * time.Second)
	return db, nil
}
func setupRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisConfig.Addr,
		Password: RedisConfig.Password,
		DB:       RedisConfig.DB,
	})
	if err := rdb.Ping().Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
func setupSession() (sessions.Store, error) {
	sess, err := sredis.NewStoreWithDB(10, "tcp",
		RedisConfig.Addr, RedisConfig.Password,
		strconv.Itoa(RedisConfig.DB), []byte(SessConfig.Secret),
	)
	if err != nil {
		return nil, err
	}
	sess.Options(sessions.Options{Path: "/", Domain: SessConfig.Domain, HttpOnly: true})
	return sess, nil
}
