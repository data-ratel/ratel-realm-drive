package misc

import (
	"os"

	"github.com/ratel-drive-core/common/util/config"
	"go.uber.org/zap"
)

// HTTP json result wrapper
type JSONResult struct {
	Data string `json:"data"`
}

func IsPathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}

	return false
}

func IsPathDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func CheckCreateDataDirectory() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	rootDir := config.GetStorageConfig().StorageRootDir
	if err := os.MkdirAll(rootDir, os.ModePerm); err != nil {
		logger.Error("Create data directory failed!",
			zap.String("error", "Please check if you have the permission!"),
		)
		return
	}
}
