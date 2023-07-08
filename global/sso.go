package global

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/proto/sso"
	"crypto/tls"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	SSOConn   *grpc.ClientConn
	SSOClient sso.SSOServiceClient
)

func setupSSO() {
	var err error
	SSOClient, err = setupSSOGrpc()
	if err != nil {
		panic(fmt.Sprintf("setup sso error, %v", err))
	}
	return
}

func setupSSOGrpc() (sSOClient sso.SSOServiceClient, err error) {
	//log.Println("get sso addr", configs.Config.SSOGrpc.Addr)
	SSOConn, err = grpc.Dial(
		configs.Config.SSOGrpc.Addr,
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
