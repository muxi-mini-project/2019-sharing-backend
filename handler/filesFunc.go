package handler

import (
	"fmt"
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

type tmpscore struct {
	Fileid int `json:"file_id"`
	Score  int `json:"score"`
}

type collecttmp struct {
	FileId        int `json:"file_id"`
	CollectlistId int `json:"collectlist_id"`
}

func UploadFile(c *gin.Context) {
	var tmpfile model.File
	var tmpuser model.User
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
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
	//因为是建立记录的同时需要建立上传记录，所以这里可以一并处理，利用
	if err := model.DB.Self.Model(&model.File{}).Order("file_id desc").Last(&tmpfile).Error; err != nil {
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
		"likes_num":   strconv.Itoa(tmpfile.LikeNum),
		"grade":       strconv.FormatFloat(tmpfile.Grade, 'f', -1, 32),
		"collect_num": strconv.Itoa(tmpfile.CollcetNum),
		"down_num":    strconv.Itoa(tmpfile.DownloadNum),
		"upload_time": tmprecord.Uploadtime,
	})
}

func DeleteFile(c *gin.Context) {
	var a model.Tmpfileid
	var tmprecord model.File_uploader
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	key, _ := model.Token_info(token)
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	//log.Println(a.FileId)
	//tmpfileid用于将只存储单个属性的结构体内的数据转化为int格式
	//tmpfileid,_ := strconv.Atoi(a.FileId)
	if err := model.DB.Self.Model(&model.File_uploader{}).Where(&model.File_uploader{FileId: a.FileId}).First(&tmprecord).Error; key != tmprecord.UploaderId {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误！",
		})
		return
	}
	if err := model.DB.Self.Where(&model.File{FileId: a.FileId}).Delete(&model.File{}).Error; err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "未找到或删除出现错误！",
		})
		return
	}
	/*if err := model.DB.Self.Model(&model.File{}).Delete(&model.File{FileId: a.FileId}).Error; err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "未找到或删除出现错误！",
		})
		return
	}*/
	c.JSON(200, gin.H{
		"message": "删除成功！",
	})
}

func DownloadFile(c *gin.Context) {
	var a model.Tmpfileid
	var tmpfile model.File
	var tmpuser model.User
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
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
	//tmpfileid用于将只存储单个属性的结构体内的数据转化为int格式
	//tmpfileid,_ := strconv.Atoi(a.FileId)
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: a.FileId}).First(&tmpfile).Error; err != nil {
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
	var a collecttmp
	var tmpuser model.User
	var tmpfile model.File
	//得到token，并解码为string的学号
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
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
	if a.CollectlistId == 0 {
		log.Print("请输入collectlist_id")
		c.JSON(401,gin.H{
			"message":"参数不全，请输入collectlist_id",
		})
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: a.FileId}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "文件未找到!",
		})
		return
	}
	if err := model.CreateNewCollectRecord(tmpfile.FileId, key, a.CollectlistId); !err {
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
	var a collecttmp
	var tmpfile model.File
	//var tmpuser model.User
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
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
	key, _ := model.Token_info(token)
	/*if err := model.DB.Self.Model(&model.File_collecter{}).Where(&model.File_collecter{CollecterId: key, FileId: a.FileId}).First(&tmpuser).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}*/
	if err := model.DB.Self.Where(&model.File_collecter{FileId: a.FileId, CollecterId: key,CollectlistId:a.CollectlistId}).Delete(&model.File_collecter{}).Error; err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"message": "未找到或取消收藏失败！",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: a.FileId}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(404,gin.H{
			"message": "找不到对应文件",
		})
		return
	}
	tmpfile.CollcetNum --
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: a.FileId}).Update("collect_num", tmpfile.CollcetNum).Error; err != nil {
		log.Println(err)
		log.Print("收藏统计失败")
		c.JSON(403,gin.H{
			"message":"收藏统计失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "取消收藏成功！",
	})
}

func Like(c *gin.Context) {
	var a model.Tmpfileid
	var tmpuser model.User
	var tmplike model.Likes
	//利用token解码出的userid来检验进行该操作的是否为已注册用户
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
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
	//利用userid和fileid在upploader的中间表里进行查询，对于找到的数据进行判断，若结构体内对应值存在说明记录存在，则判定已点赞
	if err := model.DB.Self.Model(&model.Likes{}).Where(&model.Likes{FileId: a.FileId, UserId: key}).First(&tmplike); tmplike.FileId != 0 {
		log.Println(err)
		log.Print("该用户已点过赞")
		c.JSON(401, gin.H{
			"message": "该用户已点过赞",
		})
		return
	}
	//key为userid
	if err := model.Like(a.FileId, key); !err {
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
	var a model.Tmpfileid
	//var tmprecord model.Likes
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	key, _ := model.Token_info(token)
	/*if err := model.DB.Self.Model(&model.Likes{}).Where(&model.Likes{UserId: key, FileId: a.FileId}).First(&tmprecord).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "非本人操作！",
		})
		return
	}*/
	if err := c.BindJSON(&a); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.Unlike(a.FileId, key); !err {
		log.Print("取消点赞无法记录")
		c.JSON(404, gin.H{
			"message": "取消点赞行为无法被记录！",
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
	sum := page * pagesize
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{Format: tmp.Format, College: tmp.College, Type: tmp.Type, Subject: tmp.Subject}).Count(&count).Error; err != nil {
		log.Println(err)
		log.Print("获取总数失败")
		return
	}
	i := count - sum
	if i < 0 {
		i = -1
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{Format: tmp.Format, College: tmp.College, Type: tmp.Type, Subject: tmp.Subject}).Order("file_id desc").Offset(i).Limit(pagesize).Find(&files).Error; err != nil {
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
	if err := model.DB.Self.Model(&model.File{}).Order("download_num desc").Where(&model.File{Format: tmp.Format, College: tmp.College, Type: tmp.Type, Subject: tmp.Subject}).Offset(sum).Limit(pagesize).Find(&files).Error; err != nil {
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

func Score(c *gin.Context) {
	var tmpscore tmpscore
	//var tmp model.Score
	var tmpfile model.File
	if err := c.BindJSON(&tmpscore); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	key, _ := model.Token_info(token)
	/*if err := model.DB.Self.Model(&model.Score{}).Where(&model.Score{Userid: key, Fileid: tmpscore.Fileid}).First(&tmp).Error; err != nil {
		log.Println(err)
		log.Print("该用户已评分")
		c.JSON(401, gin.H{
			"message": "该用户已评分",
		})
		return
	}*/
	if err := model.CreateScoreRecord(key, tmpscore.Fileid, tmpscore.Score); !err {
		log.Print("评分失败")
		c.JSON(401, gin.H{
			"message": "评分失败",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: tmpscore.Fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		log.Print("获取文件信息失败")
		c.JSON(404, gin.H{
			"message": "未找到相应文件",
		})
		return
	}
	s := tmpfile.Scored + 1
	tmpfile.Grade = (tmpfile.Grade*model.InttoFloat(tmpfile.Scored) + model.InttoFloat(tmpscore.Score)) / model.InttoFloat(s)
	fmt.Println(tmpfile.Grade)
	tmpfile.Scored++
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: tmpscore.Fileid}).Update(&model.File{Grade: tmpfile.Grade, Scored: tmpfile.Scored}).Error; err != nil {
		log.Println(err)
		log.Print("评分统计失败")
		c.JSON(404, gin.H{
			"message": "评分统计失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "评分成功",
	})
}
