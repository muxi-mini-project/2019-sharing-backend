package router

import (
    "github.com/gin-gonic/gin"
    "github.com/muxi-mini-project/2020-sharing-backend/handler"
    "github.com/muxi-mini-project/2020-sharing-backend/handler/user/background"
    "github.com/muxi-mini-project/2020-sharing-backend/handler/user/following"
    "github.com/muxi-mini-project/2020-sharing-backend/handler/user/image"
    "github.com/muxi-mini-project/2020-sharing-backend/handler/user/login"
    "github.com/muxi-mini-project/2020-sharing-backend/handler/user/register"
    "github.com/muxi-mini-project/2020-sharing-backend/handler/user/signture"
    "github.com/muxi-mini-project/2020-sharing-backend/handler/user/view"
)
var Router *gin.Engine


func InitRouter() {
    Router = gin.Default()

    Router.POST("/login", login.Login)
    Router.POST("/register", register.Register)
    Router.GET("/view", view.View)
    Router.PUT("/background", background.Background)
    Router.PUT("/image", image.Image)
    Router.PUT("/signture", signture.Signture)
    Router.POST("/following", following.Following)

    Router.GET("/file:fileid", handler.GetFileInfo)
    Router.DELETE("/file", handler.DeleteFile)



    return
}
