package common

import (
	"UniqueRecruitmentBackend/internal/constants"
	"context"
	"github.com/gin-gonic/gin"
)

type contextKey string

const XUID contextKey = "X-UID"
const Role contextKey = "role"

func CtxWithUID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, XUID, uid)
}

func CtxWithRole(ctx context.Context, role constants.Role) context.Context {
	return context.WithValue(ctx, Role, role)
}

func IsCandidate(c *gin.Context) bool {
	return getValue(c, "role") == string(constants.CandidateRole)
}

func IsMember(c *gin.Context) bool {
	return getValue(c, "role") == string(constants.MemberRole) || getValue(c, "role") == string(constants.Admin)
}

func IsAdmin(c *gin.Context) bool {
	return getValue(c, "role") == string(constants.Admin)
}

func GetUID(c *gin.Context) string {
	return getValue(c, "X-UID")
}

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
