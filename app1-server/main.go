package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luozui/app1-server/view"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.String(200, "hello HPU")
	})
	r.GET("/getarticle", view.GetArticle)
	r.GET("/gethostname", view.GetHostName)
	return r
}

func main() {
	// db := models.SetupDB()
	// defer db.Close()
	r := setupRouter()
	r.Run(":80")
}
