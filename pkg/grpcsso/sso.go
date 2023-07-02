package grpcsso

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/proto/sso"
	"context"
)

func GetUserByUID(uid string) (*models.User, error) {
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
