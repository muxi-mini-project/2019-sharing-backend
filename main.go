package main

import (
    "fmt"

    "github.com/jepril/sharing/router"
    "github.com/jepril/sharing/model"
)

func main() {
    model.DB.Init()        // 初始化数据库
    defer model.DB.Close() // 记得关闭数据库

    router.InitRouter()    // 初始化路由
    router.Router.Run()    // 运行
    fmt.Println("Running... Successful!")
}