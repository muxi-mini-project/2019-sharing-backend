package handler
// @Success 200 {object} model.Res "{"message":"删除成功！"}"
import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"log"
	"strconv"
)

// @Tags upload
// @Summary 上传文件,返回url，文件储存在oss上，同时得到一个可访问的地址，实质上实现了上传与下载
// @Description
// @Param token header string true "token“
// @Param fileid path string true "token“
// @Param file formData file true "资料文件"
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Success 200 {object} model.Res "{"message":"删除成功！"}"
// @Failure 401 {object} model.Error "{"message":"查无此项或查询失败!"}"
// @Failure 400 {object} model.Error "{"message":"Bad Request!"}"
// @Failure 404 {object} model.Error "{"message":"生成地址失败“} or{"message":"上传文件未生成一个可供匹配的地址，已删除，上传不成功"}"
// @Router /file/uploadbyOss/:fileid/ [post]
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