package view

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	redis "github.com/luozui/app1-server/dal/cache"
	"github.com/luozui/app1-server/dal/db"
	"github.com/luozui/app1-server/models"
)

func AddArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//article.Create()
	c.JSON(http.StatusCreated, article)
}

func GetArticle(c *gin.Context) {
	hostname, _ := os.Hostname()
	strID := c.DefaultQuery("id", "1")
	ID, _ := strconv.Atoi(strID)
	c.Header("Access-Control-Allow-Origin", "*")
	article, err := db.GetArticle(ID)
	if err != nil {
		resp := map[string]interface{}{
			"title":    "error",
			"content":  "error",
			"view_cnt": 0,
			"hostname": hostname,
		}
		c.JSON(500, resp)
		log.Println("DB error: ", err)
		return
	}
	viewCnt, err := redis.GetCntAndInc(fmt.Sprint("article_%d", ID))
	resp := map[string]interface{}{
		"title":    article.Title,
		"content":  article.Content,
		"view_cnt": viewCnt,
		"hostname": hostname,
	}
	c.JSON(http.StatusOK, resp)
}
