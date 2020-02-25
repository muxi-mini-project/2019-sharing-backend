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

// @Summary 发出留言
// @Description 根据userid检测是否为已注册用户进行的操作，随后根据传入的信息进行留言记录
// @Tags message
// @Accept json
// @Produce json
// @Param token header string true "user的认证令牌"
// @Param data body string true "请求id"
// @Success 200 {object} model.Res "{"msg":"success"}/{"msg":"需求已经被删除了!"}/{"msg":"已经处理过了!"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /message/leave/ [post]
func LeaveMessage(c *gin.Context) {
	var tmpnote note
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
	if err := c.BindJSON(&tmpnote); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if tmpnote.Content == "" {
		log.Print("请输入留言内容")
		c.JSON(404, gin.H{
			"message": "无留言内容，请输入留言内容",
		})
		return
	}
	if err := model.CreateNewMessage(key, tmpnote.Hostid, tmpnote.Content); !err {
		log.Print("无法留言")
		c.JSON(404, gin.H{
			"message": "留言失败！",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "留言成功！",
		})
	}
}

func GetMessageInfoByhostid(c *gin.Context) {
	var tmpuser model.User
	var tmpnote []model.Message
	var tmpmessage []messagelist
	var note messagelist
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
	//defaultValue为默认值，在defaultquery没有传入值时使用defaultvalue的默认值
	hostid := c.DefaultQuery("hostid", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "20"))
	sum := (page - 1) * pagesize
	if err := model.DB.Self.Model(&model.Message{}).Where(&model.Message{HostId: hostid}).Offset(sum).Limit(pagesize).Find(&tmpnote).Error; err != nil {
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
		if err := model.DB.Self.Model(&model.User{}).Where(&model.User{User_id: j.WriterId}).First(&tmpuser).Error; err != nil {
			log.Print("未能成功获取留言用户头像信息")
			c.JSON(404, gin.H{
				"message": "未成功",
			})
			return
		}
		note.image_url = tmpuser.Image_url
		tmpmessage = append(tmpmessage, note)
		log.Print(note)
	}
	c.JSON(200, gin.H{
		"message":      "操作成功",
		"message_list": tmpmessage,
	})
}
