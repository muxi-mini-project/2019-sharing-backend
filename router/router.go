package router

import (
    "github.com/gin-gonic/gin"
    "github.com/jepril/sharing/handler/user/register"
    "github.com/jepril/sharing/handler/user/login"
    "github.com/jepril/sharing/handler/user/view"
    "github.com/jepril/sharing/handler/user/background"
    "github.com/jepril/sharing/handler/user/image"
    "github.com/jepril/sharing/handler/user/signture"
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

    return
}
