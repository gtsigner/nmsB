package win

import (
	"testing"

	"./api"
)

func TestLookupPrivilegeName(t *testing.T) {
	privilegeName := "SeDebugPrivilege"

	luid, err := LookupPrivilegeValue("", privilegeName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	name, err := LookupPrivilegeName("", luid)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if name != privilegeName {
		t.Fatalf("privilege name not equal [ %s != %s ] by luid [ %d ]", privilegeName, name, luid)
	}
}

func TestEnableDebugPrivilege(t *testing.T) {
	err := SetCurrentProcessPrivilege(true, "SeShutdownPrivilege")

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	privilege, err := GetCurrentProcessPrivilege("SeShutdownPrivilege")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if privilege == nil || privilege.Attributes != api.SE_PRIVILEGE_ENABLED {
		t.Fatalf("fail to enable privilege 'SeShutdownPrivilege'")
		return
	}

}
