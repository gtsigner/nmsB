package win

import (
	"log"
	"testing"
)

func TestLookupPrivilegeName(t *testing.T) {
	privilegeName := "SeDebugPrivilege"

	luid, err := LookupPrivilegeValue("", privilegeName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	log.Println(*luid)

	name, err := LookupPrivilegeName("", *luid)
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

	names, err := GetCurrentProcessPrivileges()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	log.Println(names)

}
