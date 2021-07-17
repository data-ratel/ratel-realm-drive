package device

import (
	"testing"

	"github.com/RatelData/ratel-drive-core/common/auth"
	"github.com/RatelData/ratel-drive-core/common/util"
)

func TestGenDeviceID(t *testing.T) {
	t.Log("UUID: " + genDeviceID())
}

func TestRegisterDevice(t *testing.T) {
	util.SetAppConfigFilePath("../../config/app.json")

	loginResult, err := auth.Login("test", "test123456")
	if err != nil || !RegisterDevice(loginResult) {
		t.Fail()
	}
}
