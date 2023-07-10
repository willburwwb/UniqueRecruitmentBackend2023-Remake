package grpcsso

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/proto/sso"
	"context"
)

func GetUserInfoByUID(uid string) (*models.User, error) {
	req, err := global.SSOClient.GetUserByUID(context.Background(), &sso.GetUserByUIDRequest{Uid: uid})
	if err != nil {
		return nil, err
	}
	return &models.User{
		Uid:         req.Uid,
		Phone:       req.Phone,
		Email:       req.Email,
		Name:        req.Name,
		JoinTime:    req.JoinTime,
		AvatarUrl:   req.AvatarUrl,
		Gender:      string(req.Gender),
		Groups:      req.Groups,
		LarkUnionId: req.LarkUnionId,
	}, nil
}

func CheckRole(uid string, action string, resoure string) (bool, error) {
	_, err := global.SSOClient.CheckPermission(context.Background(), &sso.CheckPermissionRequest{
		Uid: uid,
		Object: &sso.Object{
			Action:   action,  //"Push",
			Resource: resoure, //"OpenPlatform::SMS",
		},
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// Here I only consider getting the user role
// example: admin/member/candidate
