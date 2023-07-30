package common

import (
	"UniqueRecruitmentBackend/internal/constants"
	"github.com/gin-gonic/gin"
)

func getValue(c *gin.Context, key string) string {
	get, ok := c.Get(key)
	if !ok {
		return ""
	}

	value, ok := get.(string)
	if !ok {
		return ""
	}

	return value
}

func IsCandidate(c *gin.Context) bool {
	return getValue(c, "role") == string(constants.CandidateRole)
}

func IsMember(c *gin.Context) bool {
	return getValue(c, "role") == string(constants.MemberRole)
}

func IsAdmin(c *gin.Context) bool {
	return getValue(c, "role") == string(constants.Admin)
}

func GetUID(c *gin.Context) string {
	return getValue(c, "X-UID")
}
