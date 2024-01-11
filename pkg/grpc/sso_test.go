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
