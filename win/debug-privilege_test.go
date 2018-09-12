package win

import (
	"testing"
)

func TestEnableDebugPrivilege(t *testing.T) {
	err := EnableDebugPrivilege()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

}
