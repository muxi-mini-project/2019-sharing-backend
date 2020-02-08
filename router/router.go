package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jepril/sharing/handler"
	"github.com/jepril/sharing/handler/user/background"
	"github.com/jepril/sharing/handler/user/following"
	"github.com/jepril/sharing/handler/user/image"
	"github.com/jepril/sharing/handler/user/login"
	"github.com/jepril/sharing/handler/user/register"
	"github.com/jepril/sharing/handler/user/signture"
	"github.com/jepril/sharing/handler/user/view"
	"github.com/jepril/sharing/handler/user/collection_list"
	"github.com/jepril/sharing/handler/user/deletion"
	"github.com/jepril/sharing/handler/user/down_list"
	"github.com/jepril/sharing/handler/user/fans"
	"github.com/jepril/sharing/handler/user/following_list"
	"github.com/jepril/sharing/handler/user/up_list"
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
	Router.GET("/collection_list", collection_list.CollectionList)
	Router.GET("/up_list", up_list.UpList)
	Router.GET("/down_list", down_list.DownList)
	Router.GET("/fans", fans.Fans)
	Router.GET("/following_list", following_list.FollowingList)
	Router.DELETE("/deletion", deletion.Deletion)

	return
}
