package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"server_common/graphql-example/conf"
	"server_common/graphql-example/model"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB db
var DB *gorm.DB

// Init db
func Init(){

	var err error
	var connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.Config.DB.UserName,
		conf.Config.DB.PWD, conf.Config.DB.Host, conf.Config.DB.Port, conf.Config.DB.DBName)

	DB, err = gorm.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	initTable()
}


func initTable(){
	// 自动迁移模式
	DB.AutoMigrate(&model.Auth{}, &model.Comment{}, &model.Follow{}, &model.User{}, &model.Post{})
}
