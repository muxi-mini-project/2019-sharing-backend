package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/background"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/collection_list"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/deletion"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/down_list"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/fans"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/following"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/following_list"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/image"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/login"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/register"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/signture"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/up_list"
	"github.com/muxi-mini-project/2020-sharing-backend/handler/user/view"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()

	Router.POST("/login", login.Login) //用户登录
	Router.POST("/register", register.Register)//用户注册
	Router.GET("/view", view.View)//查看用户信息
	Router.PUT("/background", background.Background)//用户背景
	Router.PUT("/image", image.Image)//用户头像
	Router.PUT("/signture", signture.Signture)//用户个性签名
	Router.POST("/following", following.Following)//follow别人
	Router.GET("/collection_list", collection_list.CollectionList)//收藏列表
	Router.GET("/up_list", up_list.UpList)//上传列表
	Router.GET("/down_list", down_list.DownList)//下载列表
	Router.GET("/fans", fans.Fans)//粉丝
	Router.GET("/following_list", following_list.FollowingList//)关注列表
	Router.DELETE("/deletion", deletion.Deletion)//取消关注

	Router.POST("/file/upload", handler.UploadFile)
	Router.GET("/file/fileinfo/:fileid", handler.GetFileInfo)
	Router.DELETE("/file/delete", handler.DeleteFile)
	Router.GET("/file/download", handler.DownloadFile)
	Router.POST("/file/collect", handler.Collect)
	Router.DELETE("/file/unfavorite", handler.Unfavourite)
	Router.POST("/file/like", handler.Like)
	Router.DELETE("/file/unlike", handler.Unlike)
	Router.GET("/file/searching/popular", handler.FileSearchingBydownloadnums)
	Router.GET("/file/searching/latest", handler.FileSearchingByuploadtime)
	Router.GET("/message/", handler.GetMessageInfoByhostid)
	Router.POST("/message/leave", handler.LeaveMessage)
	Router.POST("/file/score", handler.Score)
	Router.POST("/user/collect_list/create", handler.CreateNewCollectlist)
	Router.PUT("/user/collect_list", handler.ChangeCollectionlistName)

	return
}
