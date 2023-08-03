package global

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/constants"
	"context"
	"github.com/imroc/req/v3"
	"net/http"
	"time"
)

type SSOClient struct {
	*req.Client
}

var defaultClient *SSOClient

func GetSSOClient() *SSOClient {
	return defaultClient
}

type CheckPermissionByRoleRequest struct {
	UID  string `json:"uid"`
	Role string `json:"role"`
}

type CheckPermissionByRoleResponse struct {
	Message string `json:"message"`
	Data    struct {
		OK bool `json:"ok"`
	} `json:"data"`
}

type UserDetail struct {
	UID         string           `json:"uid"`
	Phone       string           `json:"phone"`
	Email       string           `json:"email"`
	Password    string           `json:"password,omitempty"`
	Roles       []string         `json:"roles"`
	Name        string           `json:"name"`
	AvatarURL   string           `json:"avatar_url"`
	Gender      constants.Gender `json:"gender"`
	Groups      []string         `json:"groups"`
	LarkUnionID string           `json:"lark_union_id"`
}

func (client *SSOClient) GetUserInfoByUID(ctx context.Context, uid string) (*UserDetail, error) {
	var userDetail UserDetail

	path := "/rbac/user"
	err := client.Get(path).SetQueryParam("uid", uid).Do(ctx).Into(&userDetail)
	if err != nil {
		return nil, err
	}

	return &userDetail, nil
}
func makeUidCookie(uid string) *http.Cookie {
	return &http.Cookie{
		Name:    "uid",
		Value:   uid,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	}
}
func makeSSOCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "SSO_SESSION",
		Value:   "MTY5MTA2ODI0OXxOd3dBTkVSUVIwdFZVVWxhTjFneVIxSlpVVk5PVFZST05ETXlRek0xUmtOQ1ZWVlJSMDVHUkRkVFdrMU1Xa3BOUmxCV1NGSkdRa0U9fGWADr71d6ZtXU51aNvnBKZfUqYurIYSmf3lPEjX-I0r",
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/api/v1",
	}
}
func (client *SSOClient) CheckPermissionByRole(ctx context.Context, uid, role string) (bool, error) {
	var resp CheckPermissionByRoleResponse

	path := "/rbac/user/check_permission_by_role"
	req := CheckPermissionByRoleRequest{
		UID:  uid,
		Role: role,
	}
	/*
		Due to the permission control of sso
		HTTP request needs to carry cookie
	*/
	err := client.Post(path).SetBody(req).SetCookies(makeUidCookie(uid), makeSSOCookie()).Do(ctx).Into(&resp)
	if err != nil {
		return false, err
	}

	return resp.Data.OK, nil
}

func newSSOClient() *SSOClient {
	return &SSOClient{
		Client: req.C().
			SetBaseURL(configs.Config.SSO.Addr).
			SetCommonContentType("application/json"),
	}
}

func setupSSO() {
	defaultClient = newSSOClient()
}
