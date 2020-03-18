package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"log"
	"strconv"
)

func UploadByoss(c *gin.Context){
	var tmpfile model.File
	fileid, _ := strconv.Atoi(c.Param("fileid"))
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message": "Bad Request!",
		})
		return
	}
	dataLen := header.Size

	url, err := model.Uploadfile(header.Filename, uint32(fileid), file, dataLen)
   log.Print(fileid)
	if err != nil {
		c.JSON(404,gin.H{
			"message": "生成地址失败",
		})
		return
	}
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: fileid}).First(&tmpfile).Error; err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "查无此项或查询失败!",
		})
	}

	tmpfile.FileUrl=url
	if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId: tmpfile.FileId}).Update("file_url", tmpfile.FileUrl).Error; err != nil {
		log.Println(err)
		log.Print("更新地址失败")
		if err := model.DB.Self.Model(&model.File{}).Where(&model.File{FileId:tmpfile.FileId}).Delete(&model.File{}).Error; err != nil {
			log.Println(err)
			log.Print("删除无下载地址的文件失败")
		}
		log.Print("文件记录删除成功！")
		c.JSON(404, gin.H{
			"message":"上传文件未生成一个可供匹配的地址，已删除，上传不成功",
		})
	}
    log.Print(tmpfile.FileUrl)
	c.JSON(200,gin.H{
		"message": "操作成功",
	})
}