package routers

import (
	"log"
	"wmsidm/controllers"

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

	idmv1 := r.Group("/idm/api/v1/")
	{
		idmv1.POST("/auth", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			jsonData, err := c.GetRawData()
			if c.Query("k") != controllers.MD5_KEY {
				c.JSON(400, map[string]interface{}{"error": "Need API Key"})
			} else if err != nil {
				c.JSON(400, map[string]interface{}{"error": "Bad request body"})
			} else {
				c.JSON(200, map[string]interface{}{"result": "Ini contoh response"})
				log.Println("",jsonData)
			}
		})
		// idmv1.POST("/auth", func(c *gin.Context) {
		// 	c.Header("Content-Type", "application/json; charset=utf-8")
		// 	jsonData, err := c.GetRawData()
		// 	if c.Query("k") != controllers.MD5_KEY {
		// 		c.JSON(400, map[string]interface{}{"error": "Need API Key"})
		// 	} else if err != nil {
		// 		c.JSON(400, map[string]interface{}{"error": "Bad request body"})
		// 	} else {
		// 		log.Println("",jsonData)
		// 	}
		// })
	}
}
