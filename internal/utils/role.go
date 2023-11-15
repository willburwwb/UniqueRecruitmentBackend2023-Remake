package utils

import (
	"UniqueRecruitmentBackend/internal/constants"
)

func CheckRoles(userRoles []string, roles ...constants.Role) bool {
	for _, v := range userRoles {
		for _, role := range roles {
			if v == string(role) {
				return true
			}
		}
	}
	return false
}
func CheckInGroups(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func GetMaxRole(roles []string) constants.Role {
	for _, v := range roles {
		if constants.Role(v) == constants.Admin {
			return constants.Admin
		}
	}
	for _, v := range roles {
		if constants.Role(v) == constants.MemberRole {
			return constants.MemberRole
		}
	}
	for _, v := range roles {
		if constants.Role(v) == constants.CandidateRole {
			return constants.CandidateRole
		}
	}
	return ""
}
