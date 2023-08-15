package utils

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/constants"
)

func CheckRoleByUserDetail(userDetail *global.UserDetail, roles ...constants.Role) bool {
	for _, v := range userDetail.Roles {
		for _, role := range roles {
			if v == string(role) {
				return true
			}
		}
	}
	return false
}
func CheckInArrary(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
