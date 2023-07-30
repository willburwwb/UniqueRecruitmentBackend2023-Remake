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

func GetUID(c *gin.Context) string {
	get, ok := c.Get("X-UID")
	if !ok {
		return ""
	}

	uid, ok := get.(string)
	if !ok {
		return ""
	}

	return uid
}
