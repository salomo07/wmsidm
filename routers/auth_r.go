package routers

import (
	"github.com/gin-gonic/gin"
	docs "wmsidm/docs"
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
		// ShowAccount godoc
		// @Summary      Show an account
		// @Description  get string by ID
		// @Tags         accounts
		// @Accept       json
		// @Produce      json
		// @Param        id   path      int  true  "Account ID"
		// @Success      200  {object}  model.Account
		// @Failure      400  {object}  httputil.HTTPError
		// @Failure      404  {object}  httputil.HTTPError
		// @Failure      500  {object}  httputil.HTTPError
		// @Router       /accounts/{id} [get]
		idmv1.GET("authenticate?:name", func(c *gin.Context) {
			// log.Println("df", c.Query("name"), c.Request.Host+c.Request.URL.String())
			c.JSON(200, map[string]interface{}{"Lalala": "Yeyeye", "datauser": c.Query("qw")})
			// log.Println("Ini dari : ", c.Request.Host+c.Request.URL.String())
		})
	}
}
