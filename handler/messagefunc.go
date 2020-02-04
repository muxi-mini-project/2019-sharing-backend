package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"log"
	"strconv"
)

type note struct {
	Hostid  string `json:"host_id"`
	Content string `json:"message"`
}

type messagelist struct {
	writerid  string
	content   string
	image_url string
	time      string
}

func LeaveMessage(c *gin.Context) {
	var tmpnote note
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
	if err := c.BindJSON(&tmpnote); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.CreateNewMessage(key, tmpnote.Hostid, tmpnote.Content); !err {
		log.Print("无法留言")
		c.JSON(404, gin.H{
			"message": "留言失败！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "留言成功！",
	})
}

func GetMessageInfoByhostid(c *gin.Context) {
	var tmpuser model.User
	var tmpnote []model.Message
	var tmpmessage []messagelist
	var note messagelist
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
	//defaultValue为默认值，在defaultquery没有传入值时使用defaultvalue的默认值
	hostid := c.DefaultQuery("hostid", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "20"))
	sum := (page - 1) * pagesize
	if err := model.DB.Self.Model(&model.Message{}).Where(&model.Message{HostId: hostid}).Offset(sum).Limit(pagesize).Find(&tmpnote); err != nil {
		log.Println(err)
		log.Print("")
		c.JSON(400, gin.H{
			"message": "传入参数不全",
		})
		return
	}
	//i记录序号，j表示内容
	for _, j := range tmpnote {
		note.content = j.Content
		note.writerid = j.WriterId
		note.time = j.WriteTime
		model.DB.Self.Model(&model.User{}).Where(&model.User{User_id: j.WriterId}).First(&tmpuser)
		note.image_url = tmpuser.Image_url
		tmpmessage = append(tmpmessage, note)
	}
	c.JSON(200, gin.H{
		"message":      "操作成功",
		"message_list": tmpmessage,
	})
}
