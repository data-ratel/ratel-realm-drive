package main

import (
	"fmt"

	_ "github.com/RatelData/ratel-drive-core/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/RatelData/ratel-drive-core/common/util"
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
	util.InitLogger()
	defer util.GetLogger().Sync()

	appConfig := util.GetServerConfig()
	gin.SetMode(appConfig.GetServerMode())

	util.CheckCreateDataDirectory()

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	if appConfig.IsDebugMode() {
		r.Use(cors.Default())
	} else {
		r.Use(static.Serve("/app", static.LocalFile("./ui", false)))
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	api_storage := api.Group("/storage")
	storage.RegisterAllRouters(api_storage)

	users.UsersRoutesRegister(api)

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
