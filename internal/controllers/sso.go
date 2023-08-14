package controllers

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"

	"github.com/gin-gonic/gin"
)
func getUserInfoByUID(c *gin.Context, uid string) (*global.UserDetail, error) {
	s := global.GetSSOClient()
	userDetail, err := s.GetUserInfoByUID(c, uid)
	if err != nil {
		return nil, err
	}
	return userDetail, nil
}
func getUserRoleByUID(c *gin.Context, uid string) (constants.Role, error) {
	s := global.GetSSOClient()
	userDetail, err := s.GetUserInfoByUID(c, uid)
	if err != nil {
		return "", err
	}
	roles := userDetail.Roles
	for _, v := range roles {
		if v == "admin" {
			return constants.Admin, nil
		}
	}
	for _, v := range roles {
		if v == "member" {
			return constants.MemberRole, nil
		}
	}
	return constants.CandidateRole, nil
}
