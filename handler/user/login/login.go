package login

import (
	"github.com/gin-gonic/gin"
	"github.com/jepril/sharing/model"
)

type LoginPayload struct { // 用于接收payload的结构体
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

func Login(c *gin.Context) { // 用于登录路由的处理函数
	var data LoginPayload // 声明payload变量，因为BindJSON方法需要接收一个指针进行操作
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!", // 接收过程中的错误视为Bad Request
		})
		return
	}

	if !model.CheckUserByUser_id(data.User_id) {
		c.JSON(401, gin.H{
			"message": "User does not Existed!",
		})
		return
	}

	if !model.ConfirmUser(data.User_id, data.Password) { // 检查失败的情况
		c.JSON(401, gin.H{
			"message": "Authentication Failed.",
		})
		return
	} else {
		token := model.CreateToken(data.User_id)
		c.JSON(200, gin.H{
			"message": "Authentiaction Success.",
			"token":   token,
		})
	}

	return
}
