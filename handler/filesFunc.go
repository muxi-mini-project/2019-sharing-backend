package handler

import(
	"github.com/gin-gonic/gin"
	"github.com/MitsuhaOma/goproject/winter1/2020-sharing-backend/model"
	"log"
)

type middleItem struct {
	Token  string
	Fileid int
}

func UploadFile(c *gin.Context) {
     var tmpfile model.File
     if err := c.BindJSON(&tmpfile); err != nil {
     	log.Println(err)
     	c.JSON(400, gin.H{
     		"message" : "Bad Request!",
		})
		 return
	 }
	 if err := model.CreateNewfile(tmpfile); !err {
	 	log.Println("未建立成功")
	 	c.JSON(401, gin.H{
	 		"message" : "上传失败",
		})
		 return
	 }
	 c.JSON(200, gin.H{
	 	"message" : "上传成功",
	 })
}

func GetFileInfo(c *gin.Context) {
	var tmpitem middleItem
	var tmpfile model.File
	if err := c.BindJSON(&tmpitem); err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message" : "Bad Request!",
		})
		return
	}
	if err := model.Db.Self.Model(&model.File{}).Where(&model.File{FileId:tmpitem.Fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(401,gin.H{
			"message" : "查无此项或查询失败!",
		})
	}
	c.JSON(200, gin.H{
		"message" : "信息获取成功",
		"file_name" : tmpfile.FileName ,
		"file_url" : tmpfile.FileUrl,
		
	})



}