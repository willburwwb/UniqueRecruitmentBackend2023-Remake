package utils

import (
	"UniqueRecruitmentBackend/pkg"
)

func CheckRoles(userRoles []string, roles ...pkg.Role) bool {
	for _, v := range userRoles {
		for _, role := range roles {
			if v == string(role) {
				return true
			}
		}
	}
	return false
}
func CheckInGroups(arr []string, group pkg.Group) bool {
	for _, v := range arr {
		if v == string(group) {
			return true
		}
	}
	return false
}

func GetMaxRole(roles []string) pkg.Role {
	for _, v := range roles {
		if pkg.Role(v) == pkg.Admin {
			return pkg.Admin
		}
	}
	for _, v := range roles {
		if pkg.Role(v) == pkg.MemberRole {
			return pkg.MemberRole
		}
	}
	for _, v := range roles {
		if pkg.Role(v) == pkg.CandidateRole {
			return pkg.CandidateRole
		}
	}
	return ""
}
