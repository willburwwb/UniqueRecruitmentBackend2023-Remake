package global

import (
	"UniqueRecruitmentBackend/configs"
	sso "UniqueRecruitmentBackend/internal/proto/sso"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db        *gorm.DB
	SSOConn   *grpc.ClientConn
	SSOClient sso.SSOServiceClient
	err       error
)
var (
	PgsqlConfig    *configs.PsqlConfigs
	ServerConfig   *configs.ServerConfigs
	SSOGrpcConfigs *configs.SSOGrpcConfigs
)

func Setup() error {
	Db, err = setupPgsql()
	if err != nil {
		return fmt.Errorf("pgsql setup failed :%s", err)
	}
	log.Println("Pgsql Connect Successed")

	SSOClient, err = setupSSOGrpc()
	if err != nil {
		return fmt.Errorf("sso grpc setup failed :%s", err)
	}
	return nil
}
func setupPgsql() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai ",
		PgsqlConfig.Host, PgsqlConfig.User, PgsqlConfig.Dbname, PgsqlConfig.Port)
	if PgsqlConfig.Password != "" {
		dsn = dsn + fmt.Sprintf("password=%s", PgsqlConfig.Password)
	}
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(PgsqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(PgsqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(PgsqlConfig.MaxLifeSeconds) * time.Second)
	return
}
func setupSSOGrpc() (sSOClient sso.SSOServiceClient, err error) {
	SSOConn, err = grpc.Dial(
		SSOGrpcConfigs.Addr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return
	}
	sSOClient = sso.NewSSOServiceClient(SSOConn)
	return
}

// func setupRedis() (*redis.Client, error) {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     RedisConfig.Addr,
// 		Password: RedisConfig.Password,
// 		DB:       RedisConfig.DB,
// 	})
// 	if err := rdb.Ping().Err(); err != nil {
// 		return nil, err
// 	}
// 	return rdb, nil
// }
// func setupSession() (sessions.Store, error) {
// 	sess, err := sredis.NewStoreWithDB(10, "tcp",
// 		RedisConfig.Addr, RedisConfig.Password,
// 		strconv.Itoa(RedisConfig.DB), []byte(SessConfig.Secret),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sess.Options(sessions.Options{Path: "/", Domain: SessConfig.Domain, HttpOnly: true})
// 	return sess, nil
// }
