package view

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetHostName(c *gin.Context) {
	hostname, _ := os.Hostname()
	resp := map[string]string {
		"hostname": hostname,
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, resp)
}