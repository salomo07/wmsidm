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
	r.Any("/idm/api/v1/*path", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"Lalala": "Yeyeye"})
		log.Println("Ini dari : ", c.Request.Host+c.Request.URL.String())
	})
	test := r.Group("/test")
	{
		test.POST("/write", func(c *gin.Context) {
			// jsonData, _ := c.GetRawData()
			xxx := string(c.Request.URL.Path)
			req := RequestLog{Ip: c.ClientIP(), UserAgent: c.Request.Header.Get("User-Agent"), Url: xxx}
			// log.Println(c.ClientIP(),c.Request.URL,c.Request.Header.Get("User-Agent"))

			c.Header("Content-Type", "application/json; charset=utf-8")
			// succ,_:=controllers.SaveRedis("testx",string(jsonData))
			// c.String(200, succ)
			log.Println(req)
		})
		test.GET("/read", func(c *gin.Context) {
			q, _ := c.GetQuery("key")
			succ := controllers.GetRedis(q)
			c.String(200, succ)
		})
		test.GET("/socket", func(c *gin.Context) {
			log.Println("Ini socket")
			c.HTML(200, "index.html", gin.H{})
			log.Println("Ini socketxxx")
		})
	}
}
