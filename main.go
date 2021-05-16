package main

import (
	"fmt"

	_ "github.com/RatelData/ratel-drive-core/docs"
	"github.com/gin-contrib/static"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/RatelData/ratel-drive-core/common/util/config"
	"github.com/RatelData/ratel-drive-core/common/util/misc"
	"github.com/RatelData/ratel-drive-core/service/storage"
	"github.com/RatelData/ratel-drive-core/service/users"
	"github.com/gin-gonic/gin"
)

// @title RatelDriveCore API
// @version 1.0
// @description RatelDriveCore server
// @termsOfService https://drive.rateldata.com/terms/

// @contact.name API Support
// @contact.url https://drive.rateldata.com/support
// @contact.email support@rateldata.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8666
// @BasePath /
func main() {
	appConfig := config.GetServerConfig()
	gin.SetMode(appConfig.GetServerMode())

	misc.CheckCreateDataDirectory()

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(static.Serve("/app", static.LocalFile("./ui/build", false)))

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api")

	v1_storage := v1.Group("/storage")
	storage.RegisterAllRouters(v1_storage)

	users.UsersRoutesRegister(v1)

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
