package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	models "github.com/app1-server/models"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"password",
		"10.104.230.57",
		"app1"))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

func GetArticle(id int) (*models.Article, error) {
	var article models.Article
	err := db.Where("id = ? AND status = ? ", id, 1).First(&article).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}