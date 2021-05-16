package users

import (
	"github.com/RatelData/ratel-drive-core/common/util/misc"
	"github.com/gin-gonic/gin"
)

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
}

func (umv *LoginValidator) Bind(c *gin.Context) error {
	err := misc.Bind(c, umv)
	if err != nil {
		return err
	}
	return nil
}
