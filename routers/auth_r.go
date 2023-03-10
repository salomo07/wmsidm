package routers

import (
	"wmsidm/controllers"
	docs "wmsidm/docs"

	"github.com/gin-gonic/gin"
)

type RequestLog struct {
	Ip            string `json:"ip"`
	UserAgent     string `json:"devicename"`
	Url           string `json:"url"`
	Authorization string `json:"authorization"`
	Payload       string `json:"payload"`
}

func IDMRouter(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	idmv1 := r.Group("/idm/api/v1/")
	{
		idmv1.POST("/setredis?:key", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			jsonData, err := c.GetRawData()
			if err != nil {
				c.JSON(400, map[string]interface{}{"error": "Bad request body"})
			} else if c.Query("key") == "" {
				c.JSON(400, map[string]interface{}{"error": "Key is not found"})
			} else {
				xxx, errSave := controllers.SaveRedis(c.Query("key"), string(jsonData))
				if errSave != "" {
					c.JSON(500, map[string]interface{}{"error": "Error occured when saving to redis server"})
				} else {
					c.String(200, xxx)
				}
			}

		})
		idmv1.GET("/getredis?:key", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			if c.Query("key") == "" {
				c.JSON(400, map[string]interface{}{"error": "Key is not found"})
			} else {
				xxx, errGet := controllers.GetRedis(c.Query("key"))
				if errGet != "" {
					c.JSON(500, map[string]interface{}{"error": errGet})
				} else {
					c.String(200, xxx)
				}
			}
		})
	}
}
