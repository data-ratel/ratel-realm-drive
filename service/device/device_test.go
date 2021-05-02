package device

import (
	"testing"

	"github.com/ratel-drive-core/common/auth"
	"github.com/ratel-drive-core/common/util/config"
)

func TestGenDeviceID(t *testing.T) {
	t.Log("UUID: " + genDeviceID())
}

func TestRegisterDevice(t *testing.T) {
	config.SetAppConfigFilePath("../../config/app.json")

	loginResult, err := auth.Login("test", "test123456")
	if err != nil || !RegisterDevice(loginResult) {
		t.Fail()
	}
}
