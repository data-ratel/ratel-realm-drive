package auth

import (
	"testing"

	"github.com/ratel-drive-core/common/util/config"
)

func TestLogin(t *testing.T) {
	config.SetAppConfigFilePath("../../config/app.json")
	_, err := Login("test", "test123456")
	if err != nil {
		t.Fail()
	}
}
