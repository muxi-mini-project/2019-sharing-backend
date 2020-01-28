package router

import(
	"github.com/MitsuhaOma/goproject/winter1/2020-sharing-backend/handler"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	Router.POST("/file/upload", handler.UploadFile)
	
}