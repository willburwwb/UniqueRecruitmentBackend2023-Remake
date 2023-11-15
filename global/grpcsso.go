package global

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/constants"
	pb "UniqueRecruitmentBackend/pkg/proto/sso"
	"context"
	"crypto/tls"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func setupSSOGrpc() (*GrpcSSOClient, error) {
	ssoConn, err := grpc.Dial(
		configs.Config.SSO.Addr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	grpcClient := pb.NewSSOServiceClient(ssoConn)
	return &GrpcSSOClient{grpcClient}, err
}

type GrpcSSOClient struct {
	pb.SSOServiceClient
}

var defaultGrpcClient *GrpcSSOClient

func GetGrpcSSOClient() *GrpcSSOClient {
	if defaultClient == nil {
		defaultGrpcClient, _ = setupSSOGrpc()
	}
	return defaultGrpcClient
}
func (client *GrpcSSOClient) GetUserInfoByUID(uid string) (*UserDetail, error) {
	req := &pb.GetUserByUIDRequest{
		Uid: uid,
	}
	ctx := context.Background()
	resp, err := client.GetUserByUID(ctx, req)
	if err != nil {
		return nil, err
	}
	return &UserDetail{
		UID:         resp.Uid,
		Name:        resp.Name,
		Email:       resp.Email,
		Phone:       resp.Phone,
		AvatarURL:   resp.AvatarUrl,
		Groups:      resp.Groups,
		Gender:      constants.Gender(resp.Gender),
		LarkUnionID: resp.LarkUnionId,
	}, nil
}

// TODO:
// Add grpc permission check
// func (client *GrpcSSOClient) CheckPermissionByRole(uid, role string) (bool, rerror) {
// 	req := &pb.CheckPermissionRequest{
// 		Uid:  uid,
// 		Object: ,
// 	}
// 	ctx := context.Background()
// 	resp, err := client.CheckPermissionByRole(ctx, req)
// 	if err != nil {
// 		return false, err
// 	}
// 	return resp.Ok, nil
// }
