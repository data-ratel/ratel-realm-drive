package auth

import (
	"testing"

	"github.com/RatelData/ratel-drive-core/common/util/config"
)

func TestLogin(t *testing.T) {
	config.SetAppConfigFilePath("../../config/app.json")
	_, err := Login("test@test.com", "test123456")
	if err != nil {
		t.Fail()
	}
}
