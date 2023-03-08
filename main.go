package main

import (
	_ "log"
	"wmsidm/routers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// @title           Kukuruyuk.com
	// @version         1.0
	// @description     Ini description, tolong dibaca dulu.
	// @termsOfService  http://swagger.io/terms/

	// @contact.name   API Support
	// @contact.url    http://www.swagger.io/support
	// @contact.email  support@swagger.io

	// @license.name  Apache 2.0
	// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

	// @host      localhost:8080
	// @BasePath  /api/v1

	// @securityDefinitions.basic  BasicAuth

	// @externalDocs.description  OpenAPI
	// @externalDocs.url          https://swagger.io/resources/open-api/

	port := "7890"
	// os.Getenv("PORT")
	r := gin.Default()

	routers.IDMRouter(r)
	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"code": "404", "message": "Endpoint is not found"})
	// })

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// r.Run()
	r.Run(":" + port)
}
