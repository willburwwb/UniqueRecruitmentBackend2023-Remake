package global

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/constants"
	"context"
	"github.com/imroc/req/v3"
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
	Ok bool `json:"ok"`
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

func (client *SSOClient) CheckPermissionByRole(ctx context.Context, req CheckPermissionByRoleRequest) (*CheckPermissionByRoleResponse, error) {
	var resp CheckPermissionByRoleResponse

	path := "/rbac/user/check_permission_by_role"
	err := client.Post(path).SetBody(req).Do(ctx).Into(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func newSSOClient() *SSOClient {
	return &SSOClient{
		Client: req.C().
			SetBaseURL(configs.Config.SSO.Addr).
			SetCommonContentType("application/json"),
	}
}

func init() {
	defaultClient = newSSOClient()
}
