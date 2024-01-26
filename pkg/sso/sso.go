package sso

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"

	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/pkg"
	"UniqueRecruitmentBackend/pkg/logger"
)

const UniqueSessionName = "SSO_SESSION"

type SSOClient struct {
	*req.Client
}

var defaultClient *SSOClient

type UserDetailResponse struct {
	Message string `json:"message"`
	Data    struct {
		UserDetail
	} `json:"data"`
}

type UserDetail struct {
	UID         string     `json:"uid"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	Password    string     `json:"password,omitempty"`
	Roles       []string   `json:"roles"`
	Name        string     `json:"name"`
	AvatarURL   string     `json:"avatar_url"`
	Gender      pkg.Gender `json:"gender"`
	JoinTime    string     `json:"join_time"`
	Groups      []string   `json:"groups"`
	LarkUnionID string     `json:"lark_union_id"`
}

func makeSSOCookie(ctx *gin.Context) *http.Cookie {
	SSOSession, _ := ctx.Cookie(UniqueSessionName)
	return &http.Cookie{
		Name:    UniqueSessionName,
		Value:   SSOSession,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/api/v1",
	}
}
func GetUserInfoByUID(ctx *gin.Context, uid string) (*UserDetail, error) {
	var req UserDetailResponse

	path := "/rbac/user"
	err := defaultClient.Get(path).SetQueryParam("uid", uid).
		SetCookies(makeSSOCookie(ctx)).Do(ctx).Into(&req)

	if err != nil {
		return nil, err
	}
	return &req.Data.UserDetail, nil
}

func newSSOClient() *SSOClient {
	return &SSOClient{
		Client: req.C().
			SetBaseURL(configs.Config.SSO.Addr).
			SetCommonContentType("application/json"),
	}
}

func init() {
	logger.InfoF("[Init] set up http sso client")
	defaultClient = newSSOClient()
}
