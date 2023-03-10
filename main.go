package main

import (
	_ "log"
	"wmsidm/routers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	port := "7890"
	// os.Getenv("PORT")
	r := gin.Default()

	routers.IDMRouter(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Endpoint isn't found"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// r.Run()
	r.Run(":" + port)
}
