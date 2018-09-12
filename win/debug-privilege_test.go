package win

import (
	"log"
	"testing"

	"./api"

	"golang.org/x/sys/windows"
)

func TestEnableDebugPrivilege(t *testing.T) {
	err := EnableDebugPrivilege()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	token, err := windows.OpenCurrentProcessToken()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	privileges, err := api.GetTokenPrivileges(token)
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, privilege := range privileges {
		log.Println(privilege.Luid)
	}

}
