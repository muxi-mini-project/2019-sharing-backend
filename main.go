package main

import (
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"github.com/muxi-mini-project/2020-sharing-backend/router"
	"log"
)

func main() {
	model.DB.Init()
	defer model.DB.Close()
	router.InitRouter()
	if err := router.Router.Run(":8080"); err != nil {
		log.Print("路由错误")
		log.Fatal(err)
	}
}
