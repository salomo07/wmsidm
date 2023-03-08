package routers

import (
	"log"

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
	// r.Any("", func(c *gin.Context) {
	// 	log.Println("Ini dari : ", c.Request.Host+c.Request.URL.String())
	// })
	// r.Any("/idm/api/v1/authenticate", func(c *gin.Context) {
	// 	c.JSON(200, map[string]interface{}{"Lalala": "Yeyeye"})
	// 	log.Println("Ini dari : ", c.Request.Host+c.Request.URL.String())
	// })
	// r.Any("/*path", func(c *gin.Context) {
	// 	log.Println("dfgx")
	// })
	// r.GET("/idm/api/v1/authenticate", func(c *gin.Context) {
	// 	log.Println("df", c.Request.Host+c.Request.URL.String())
	// })
	idmv1 := r.Group("/idm/api/v1/")
	{
		idmv1.GET("authenticate?:name", func(c *gin.Context) {
			log.Println("df", c.Query("name"), c.Request.Host+c.Request.URL.String())
			c.JSON(200, map[string]interface{}{"Lalala": "Yeyeye"})
			log.Println("Ini dari : ", c.Request.Host+c.Request.URL.String())
		})
	}
}
