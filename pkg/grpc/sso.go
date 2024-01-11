package grpc

import (
	"UniqueRecruitmentBackend/internal/constants"
	pb "UniqueRecruitmentBackend/pkg/proto/sso"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcSSOClient struct {
	pb.SSOServiceClient
}

var defaultGrpcClient *GrpcSSOClient

type UserDetail struct {
	UID         string           `json:"uid"`
	Phone       string           `json:"phone"`
	Email       string           `json:"email"`
	Password    string           `json:"password,omitempty"`
	Name        string           `json:"name"`
	AvatarURL   string           `json:"avatar_url"`
	Gender      constants.Gender `json:"gender"`
	JoinTime    string           `json:"join_time"`
	Groups      []string         `json:"groups"`
	LarkUnionID string           `json:"lark_union_id"`
}

func GetUserInfoByUID(uid string) (*UserDetail, error) {
	req := &pb.GetUserByUIDRequest{
		Uid: uid,
	}
	ctx := context.Background()
	resp, err := defaultGrpcClient.GetUserByUID(ctx, req)
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
		JoinTime:    resp.JoinTime,
		Gender:      constants.Gender(resp.Gender),
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
