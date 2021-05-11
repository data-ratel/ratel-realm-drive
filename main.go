package main

import (
	"fmt"

	_ "github.com/ratel-drive-core/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ratel-drive-core/common/util/config"
	"github.com/ratel-drive-core/common/util/misc"
	"github.com/ratel-drive-core/service/storage"
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

// @host drive.rateldata.com
// @BasePath /api
func main() {
	appConfig := config.GetServerConfig()
	gin.SetMode(appConfig.GetServerMode())

	misc.CheckCreateDataDirectory()

	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(static.Serve("/", static.LocalFile("./ui/build", false)))

	url := ginSwagger.URL("/docs/doc.json") // The url pointing to API definition
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := r.Group("/api")
	v1_storage := v1.Group("/storage")
	storage.RegisterAllRouters(v1_storage)

	r.Run(fmt.Sprintf(":%d", appConfig.ServerPort))
}
