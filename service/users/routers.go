package users

import (
	"net/http"

	"github.com/RatelData/ratel-drive-core/common/auth"
	"github.com/RatelData/ratel-drive-core/common/errors"
	"github.com/RatelData/ratel-drive-core/service/device"
	"github.com/gin-gonic/gin"
)

func UsersRoutesRegister(router *gin.RouterGroup) {
	router.POST("/login", UserLogin)
}

func UserLogin(c *gin.Context) {
	lv := LoginValidator{}
	if err := lv.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.NewValidatorError(err))
		return
	}

	loginResult, loginErr := auth.Login(lv.User.Email, lv.User.Password)
	if loginErr != nil {
		c.JSON(http.StatusUnauthorized, errors.NewValidatorError(loginErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"token": loginResult.User.Token,
		},
	})

	device.RegisterDevice(loginResult)
}
