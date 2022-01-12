package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"simple-rest-curd/component"
	"simple-rest-curd/modules/user/transport/ginuser"
)

func main() {
	dsn := os.Getenv("MySqlConnect")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err) //log.fatall neu co se dung chuong trinh luon - khac fmt o la co in them thoi gian loi
	}

	appCtx := component.NewAppContext(db)

	r := gin.Default()
	v1 := r.Group("/v1/users")
	{
		v1.POST("", ginuser.CreateUser(appCtx))
		v1.GET("", ginuser.ListUser(appCtx))
	}
	r.Run()
}
