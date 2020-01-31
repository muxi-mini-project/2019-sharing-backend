package main

import (
    "fmt"

    "github.com/muxi-mini-project/2020-sharing-backend/router"
    "github.com/muxi-mini-project/2020-sharing-backend/model"
)

func main() {
    model.DB.Init()        // 初始化数据库
    defer model.DB.Close() // 记得关闭数据库

    router.InitRouter()    // 初始化路由
    router.Router.Run()    // 运行
    fmt.Println("Running... Successful!")
}