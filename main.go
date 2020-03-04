package main

import (
	"fmt"

	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"github.com/muxi-mini-project/2020-sharing-backend/router"
)


// @title sharing
// @version 1.0
// @description 资料共享平台

// @host 47.102.120.167:2333
// @BasePath /sharing/v1/

// @Schemas http

func main() {
	model.DB.Init()        // 初始化数据库
	defer model.DB.Close() // 记得关闭数据库

	router.InitRouter() // 初始化路由
	router.Router.Run() // 运行
	fmt.Println("Running... Successful!")
}
