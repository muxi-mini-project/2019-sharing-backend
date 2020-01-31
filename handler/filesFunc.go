package handler

import(
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"log"
	"strconv"
)

type middleItem struct {
	Fileid int
}

func UploadFile(c *gin.Context) {
     token := c.Request.Header.Get("token")
     key,_ := untoken(token)

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
	var tmpfile model.File
	var tmprecord model.File_uploader
	fileid, _ := strconv.Atoi(c.Param("fileid"))
	if err := model.Db.Self.Model(&model.File{}).Where(&model.File{FileId:fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(401,gin.H{
			"message" : "查无此项或查询失败!",
		})
	}
	if err := model.Db.Self.Model(&model.File_uploader{}).Where(&model.File_uploader{Fileid:fileid}).First(&tmprecord).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message" : "该文件的上传时间查询出现问题！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message" : "信息获取成功",
		"file_name" : tmpfile.FileName ,
		"file_url" : tmpfile.FileUrl,
		"format" : tmpfile.Format ,
		"content" : tmpfile.Content ,
		"subject": tmpfile.College ,
		"likes_num": strconv.Itoa(tmpfile.Likes) ,
		"grade": strconv.FormatFloat(tmpfile.Grade,'f',-1,32) ,
		"collect_num": strconv.Itoa(tmpfile.CollcetNumber),
		"down_num": strconv.Itoa(tmpfile.DownloadNumber) ,
	})
}

func DeleteFile(c *gin.Context) {
	var a int
	var tmprecord model.File_uploader
	token := c.Request.Header.Get("token")
	key,_ := untoken(token)
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message" : "Bad Request!",
		})
		return
	}

	if err := model.Db.Self.Model(&model.File_uploader{}).Where(&model.File_uploader{Fileid:a}).First(&tmprecord).Error; key == tmprecord.Uploaderid{
		log.Println(err)
		c.JSON(401, gin.H{
			"message" : "身份认证错误！" ,
		})
		return
	}

	if err := model.Db.Self.Model(&model.File{}).Delete(&model.File{FileId:a}).Error; err != nil {
		log.Println(err)
		c.JSON(404,gin.H{
			"message" : "未找到或删除出现错误！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message" : "删除成功！" ,
	})
}

func Download(c *gin.Context) {
     var a int
     var tmpfile model.File

	 if err := c.BindJSON(&a); err != nil {
		 log.Println(err)
		 c.JSON(400, gin.H{
			 "message" : "Bad Request!",
		 })
		 return
	 }
	 if err := model.Db.Self.Model(&model.File{}).Where(&model.File{FileId:a}).First(&tmpfile).Error; err != nil {
	 	log.Println(err)
	 	c.JSON(404, gin.H{
	 		"message" : "文件未找到!",
		})
		 return
	 }
	 c.JSON(200, gin.H{
	 	"message" : "下载成功",
	 	"file_url" : tmpfile.FileUrl ,
	 })
}

func Collect(c *gin.Context) {
	var a int
//得到token，并解码为string的学号
	token := c.Request.Header.Get("token")
	key,_ := untoken(token)

	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message" : "Bad Request!",
		})
		return
	}



}
















