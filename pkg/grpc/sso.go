package grpc

import (
	"UniqueRecruitmentBackend/pkg"
	pb "UniqueRecruitmentBackend/pkg/proto/sso"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcSSOClient struct {
	pb.SSOServiceClient
}

var defaultGrpcClient *GrpcSSOClient

func GetUserInfoByUID(uid string) (*pkg.UserDetail, error) {
	req := &pb.GetUserByUIDRequest{
		Uid: uid,
	}
	ctx := context.Background()
	resp, err := defaultGrpcClient.GetUserByUID(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pkg.UserDetail{
		UID:         resp.Uid,
		Name:        resp.Name,
		Email:       resp.Email,
		Phone:       resp.Phone,
		AvatarURL:   resp.AvatarUrl,
		Groups:      resp.Groups,
		JoinTime:    resp.JoinTime,
		Gender:      pkg.Gender(resp.Gender),
		LarkUnionID: resp.LarkUnionId,
	}, nil
}

func GetRolesByUID(uid string) ([]string, error) {
	req := &pb.GetRolesByUIDRequest{
		Uid: uid,
	}
	ctx := context.Background()
	resp, err := defaultGrpcClient.GetRolesByUID(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetRoles(), nil
}

func init() {
	var err error
	defaultGrpcClient, err = setupSSOGrpc()
	if err != nil {
		return
	}
}

func setupSSOGrpc() (*GrpcSSOClient, error) {
	ssoConn, err := grpc.Dial(
		//configs.Config.Grpc.Addr,
		"dev.back.sso.hustunique.com:50000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	grpcClient := pb.NewSSOServiceClient(ssoConn)
	return &GrpcSSOClient{grpcClient}, err
}
