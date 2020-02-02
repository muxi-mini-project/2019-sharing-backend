package handler

import(
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"log"
	"strconv"
)

type note struct {
	Hostid string `json:"host_id"`
	Content  string `json:"message"`
}

func LeaveMessage(c *gin.Context){
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
	if err := model.CreateNewMessage(key,tmpnote.Hostid,tmpnote.Content); !err {
		log.Print("无法留言")
		c.JSON(404, gin.H{
			"message" : "留言失败！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "留言成功！",
	})
}

func GetMessageInfoByhostid(c *gin.Context) {
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
	hostid, _ := strconv.Atoi(c.Param("hostid"))
	
}