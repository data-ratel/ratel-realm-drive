package device

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/ratel-drive-core/common/auth"
	"github.com/ratel-drive-core/common/util/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RegisterDeviceReq struct {
	Device struct {
		DeviceID    string `json:"device_id"`
		UserID      uint   `json:"user_id"`
		ServicePort int    `json:"service_port"`
	} `json:"device"`
}

func RegisterDevice(loginResult auth.LoginResult) bool {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	token := loginResult.User.Token
	userID := loginResult.User.UserID

	endpoint := "/api/devices/register"

	registerDevReq := RegisterDeviceReq{}
	registerDevReq.Device.DeviceID = genDeviceID()
	registerDevReq.Device.UserID = userID
	registerDevReq.Device.ServicePort = config.GetServerConfig().ServerPort

	body, _ := json.Marshal(registerDevReq)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		SetBody(body).
		Post(config.CentralHost() + endpoint)

	return handleRegisterDeviceResult(resp, err)
}

// Generate a unique device id
func genDeviceID() string {
	return uuid.NewString()
}

func handleRegisterDeviceResult(resp *resty.Response, err error) bool {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err == nil && resp.StatusCode() == 201 {
		logger.Info("Register device succeed!",
			zap.String("body", resp.String()),
		)
		return true
	}

	var errInfo zapcore.Field
	if err != nil {
		errInfo = zap.String("error", err.Error())
	} else {
		errInfo = zap.String("status", resp.Status())
	}

	logger.Error("Register device failed!",
		errInfo,
	)

	return false
}
