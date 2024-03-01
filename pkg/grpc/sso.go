package grpc

import (
	"UniqueRecruitmentBackend/pkg"
	pb "UniqueRecruitmentBackend/pkg/proto/sso"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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

func GetUsers(uids []string) ([]pkg.UserDetail, error) {
	req := &pb.GetUsersRequest{
		Uid: uids,
	}
	ctx := context.Background()
	resp, err := defaultGrpcClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	users := make([]pkg.UserDetail, 0)
	for _, user := range resp.Users {
		users = append(users, pkg.UserDetail{
			UID:         user.Uid,
			Name:        user.Name,
			Email:       user.Email,
			Phone:       user.Phone,
			AvatarURL:   user.AvatarUrl,
			Groups:      user.Groups,
			JoinTime:    user.JoinTime,
			Gender:      pkg.Gender(user.Gender),
			LarkUnionID: user.LarkUnionId,
		})
	}
	return users, nil
}

func GetGroupsDetail() (map[string]int, error) {
	ctx := context.Background()
	resp, err := defaultGrpcClient.GetGroupsDetail(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	groupsDetail := make(map[string]int)
	for group, count := range resp.Groups.Fields {
		groupsDetail[group] = int(count.GetNumberValue())
	}
	return groupsDetail, nil
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
