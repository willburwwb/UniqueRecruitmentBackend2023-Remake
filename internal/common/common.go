package common

import "github.com/gin-gonic/gin"

func IsCandidate(uid string) bool {
	// TODO wait for sso
	return true
}

func IsMember(uid string) bool {
	// TODO wait for sso
	return true
}

func IsAdmin(uid string) bool {
	// TODO wait for sso
	return true
}

func GetUserID(c *gin.Context) string {
	// TODO wait for sso
	return "thisisuserid"
}
