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
	Writerid  string `json:"writer_id"`
	Content   string `json:"message"`
	Image_url string `json:"image_url"`
	Time      string `json:"write_time"`
}

// @Summary 发出留言
// @Description 根据userid检测是否为已注册用户进行的操作，随后根据传入的信息进行留言记录
// @Tags message
// @Accept json
// @Produce json
// @Param token header string true "user的认证令牌"
// @Param data body model.Note true "留言的对象以及留言内容"
// @Success 200 {object} model.Res "{"message":"留言成功!"}"
// @Failure 401 {object} model.Error "{"message":"身份认证错误，请先登录或注册！"}"
// @Failure 400 {object} model.Error "{"message":"Bad Request!"}"
// @Failure 404 {object} model.Error "{"message":"无留言内容，请输入留言内容"} or {"message":"留言失败！"} "
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

// @Summary 获取用户留言板信息
// @Description 传入hostid与pagesize，page获取数组数据
// @Tags message
// @Accept json
// @Produce json
// @Param token header string true "user的认证令牌"
// @Param hostid query string true "留言板主人的id"
// @Param page query string true "页码"
// @Param pagesize query string true "页码内容大小，即显示数量"
// @Success 200 {object} model.Getmessage "{"message":"操作成功","message_list":数组,返回一系列留言属性}"
// @Failure 401 {object} model.Error "{"message":"身份认证错误，请先登录或注册！"}"
// @Failure 400 {object} model.Error "{"message":"传入参数不全"}"
// @Failure 404 {object} model.Error "{"message":"未能成功找到相应留言内容"} or {"message":"未能成功获取留言用户头像信息！"}"
// @Router /message/ [get]
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
	if hostid == "" {
		log.Print("请至少输入hostid")
		c.JSON(400, gin.H{
			"message": "传入参数不全",
		})
		return
	}
	if err := model.DB.Self.Model(&model.Message{}).Where(&model.Message{HostId: hostid}).Offset(sum).Limit(pagesize).Find(&tmpnote).Error; err != nil {
		log.Println(err)
		log.Print("")
		c.JSON(404, gin.H{
			"message": "未能成功找到相应留言内容",
		})
		return
	}
	//i记录序号，j表示内容
	for _, j := range tmpnote {
		note.Content = j.Content
		note.Writerid = j.WriterId
		note.Time = j.WriteTime
		if err := model.DB.Self.Model(&model.User{}).Where(&model.User{User_id: j.WriterId}).First(&tmpuser).Error; err != nil {
			log.Print("未能成功获取留言用户头像信息")
			c.JSON(404, gin.H{
				"message": "未成功",
			})
			return
		}
		note.Image_url = tmpuser.Image_url
		tmpmessage = append(tmpmessage, note)
		//log.Print(tmpnote)
	}
	//log.Print(tmpmessage)
	c.JSON(200, gin.H{
		"message":      "操作成功",
		"message_list": tmpmessage,
	})
}
