package router

import (
    "github.com/gin-gonic/gin"
    "github.com/jepril/sharing/handler/user/register"
    "github.com/jepril/sharing/handler/user/login"
    "github.com/jepril/sharing/handler/user/view"
)

var Router *gin.Engine

func InitRouter() {
    Router = gin.Default()
    Router.POST("/login", login.Login)
    Router.POST("/register", register.Register)
    Router.GET("/view", view.View)
   
    return

}