package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"log"
	"strconv"
)

type tmp struct {
	College string `json:"college"`
	Subject string `json:"subject"`
	Format  string `json:"format"`
	Type    string `json:"type"`
}

func UploadFile(c *gin.Context) {
	var tmpfile model.File
	var tmpuser model.User
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	key, _ := model.Token_info(token)
	if err := model.DB.Self.Model(&model.User{}).Where(&model.User{User_id: key}).First(&tmpuser).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	//读取发送来的json格式的数据并储存，检验成功与否
	if err := c.BindJSON(&tmpfile); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	//建立新的记录，检验成功与否
	if err := model.CreateNewfile(tmpfile); !err {
		log.Print("建立数据失败")
		c.JSON(404, gin.H{
			"message": "建立数据失败",
		})
		return
	}

	if err := model.CreateNewUploadRecord(tmpfile.FileId, key); !err {
		log.Print("上传无法记录")
		c.JSON(404, gin.H{
			"message": "上传行为无法被记录！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "上传成功",
	})
}

func GetFileInfo(c *gin.Context) {
	var tmpfile model.File
	var tmprecord model.File_uploader
	fileid, _ := strconv.Atoi(c.Param("fileid"))
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "查无此项或查询失败!",
		})
	}
	if err := model.DB.Self.Model(&model.File_uploader{}).Where(&model.File_uploader{FileId: fileid}).First(&tmprecord).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "该文件的上传时间查询出现问题！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":     "信息获取成功",
		"file_name":   tmpfile.FileName,
		"file_url":    tmpfile.FileUrl,
		"format":      tmpfile.Format,
		"content":     tmpfile.Content,
		"subject":     tmpfile.College,
		"likes_num":   strconv.Itoa(tmpfile.Likes),
		"grade":       strconv.FormatFloat(tmpfile.Grade, 'f', -1, 32),
		"collect_num": strconv.Itoa(tmpfile.CollcetNumber),
		"down_num":    strconv.Itoa(tmpfile.DownloadNumber),
	})
}

func DeleteFile(c *gin.Context) {
	var a int
	var tmprecord model.File_uploader
	token := c.Request.Header.Get("token")
	key, _ := model.Token_info(token)
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File_uploader{}).Where(&model.File_uploader{FileId: a}).First(&tmprecord).Error; key == tmprecord.UploaderId {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误！",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File{}).Delete(&model.File{FileId: a}).Error; err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "未找到或删除出现错误！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功！",
	})
}

func DownloadFile(c *gin.Context) {
	var a int
	var tmpfile model.File
	var tmpuser model.User
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	key, _ := model.Token_info(token)
	if err := model.DB.Self.Model(&model.User{}).Where(&model.User{User_id: key}).First(&tmpuser).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: a}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "文件未找到!",
		})
		return
	}
	if err := model.CreateNewDownloadRecord(tmpfile.FileId, key); !err {
		log.Print("下载无法记录")
		c.JSON(404, gin.H{
			"message": "下载行为无法被记录！",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "下载成功",
		"file_url": tmpfile.FileUrl,
	})
}

func Collect(c *gin.Context) {
	var a int
	var tmpuser model.User
	var tmpfile model.File
	//得到token，并解码为string的学号
	token := c.Request.Header.Get("token")
	key, _ := model.Token_info(token)
	if err := model.DB.Self.Model(&model.User{}).Where(&model.User{User_id: key}).First(&tmpuser).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: a}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "文件未找到!",
		})
		return
	}
	if err := model.CreateNewCollectRecord(tmpfile.FileId, key); !err {
		log.Print("收藏无法记录")
		c.JSON(404, gin.H{
			"message": "收藏行为无法被记录！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "收藏成功！",
	})
}

func Unfavourite(c *gin.Context) {
	var a int
	var tmpuser model.User
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	key, _ := model.Token_info(token)
	if err := model.DB.Self.Model(&model.File_collecter{}).Where(&model.File_collecter{CollecterId: key, FileId: a}).First(&tmpuser).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File_collecter{}).Delete(&model.File_collecter{FileId: a, CollecterId: key}); err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "未找到或取消收藏失败！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "取消收藏成功！",
	})
}

func Like(c *gin.Context) {
	var a int
	var tmpuser model.User
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	key, _ := model.Token_info(token)
	if err := model.DB.Self.Model(&model.User{}).Where(&model.User{User_id: key}).First(&tmpuser).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.Like(a, key); !err {
		log.Print("点赞无法记录")
		c.JSON(404, gin.H{
			"message": "点赞行为无法被记录！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "点赞成功！",
	})
}

func Unlike(c *gin.Context) {
	var a int
	var tmprecord model.Likes
	token := c.Request.Header.Get("token")
	key, _ := model.Token_info(token)
	if err := model.DB.Self.Model(&model.Likes{}).Where(&model.Likes{UserId: key, FileId: a}).First(&tmprecord).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Model(&model.Likes{}).Delete(&model.Likes{FileId: a, UserId: key}); err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "未找到或取消点赞失败！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "取消点赞成功！",
	})
}

func FileSearchingByuploadtime(c *gin.Context) {
	var tmp tmp
	var files []model.File
	var count int
	if err := c.BindJSON(&tmp); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "20"))
	sum := (page - 1) * pagesize
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{Format: tmp.Format, College: tmp.College, Type: tmp.Type, Subject: tmp.Subject}).Count(&count); err != nil {
		log.Println(err)
		log.Print("获取总数失败")
		return
	}
	i := count - sum
	if i < 0 {
		i = -1
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{Format: tmp.Format, College: tmp.College, Type: tmp.Type, Subject: tmp.Subject}).Offset(i).Limit(pagesize).Find(&files); err != nil {
		log.Println(err)
		log.Print("获取数据失败")
		c.JSON(400, gin.H{
			"message": "数据获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "获取成功",
		"file":    files,
	})
}

func FileSearchingBydownloadnums(c *gin.Context) {
	var tmp tmp
	var files []model.File
	var count int
	if err := c.BindJSON(&tmp); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{Format: tmp.Format, College: tmp.College, Type: tmp.Type, Subject: tmp.Subject}).Count(&count); err != nil {
		log.Println(err)
		log.Print("获取总数失败")
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "20"))
	sum2 := (page - 1) * pagesize

	if err := model.DB.Self.Model(&model.File{}).Order("download_num desc").Where(&model.File{Format: tmp.Format, College: tmp.College, Type: tmp.Type, Subject: tmp.Subject}).Offset(sum2).Limit(pagesize).Find(&files); err != nil {
		log.Println(err)
		log.Print("获取数据失败")
		c.JSON(400, gin.H{
			"message": "数据获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "获取成功",
		"file":    files,
	})
}
