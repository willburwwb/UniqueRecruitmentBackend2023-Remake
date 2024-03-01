package grpc

import "testing"

func TestGetUserInfoByUID(t *testing.T) {
	userInfo, err := GetUserInfoByUID("555f3016-3b01-4dcd-bbac-cb033d19caf6")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("Get UserInfo Success")
	t.Logf("%#v", userInfo)
}

func TestGetRolesByUID(t *testing.T) {
	userRoles, err := GetRolesByUID("c4fb1c23-e9de-40a6-b1d4-b4bc2df0a625")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("Get UserRoles Success")
	t.Logf("%#v", userRoles)
}

func TestGetUsers(t *testing.T) {
	users, err := GetUsers([]string{
		"036e59aa-3fb7-4ae3-91a3-6c439f2576cc",
		"fb00b144-f67c-490a-81ab-01201bfb6974",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("Get Users Success")
	t.Logf("%#v", users)
}

func TestGetGroupsDetail(t *testing.T) {
	groupsDetail, err := GetGroupsDetail()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("Get GroupsDetail Success")
	t.Logf("%#v", groupsDetail)
}
