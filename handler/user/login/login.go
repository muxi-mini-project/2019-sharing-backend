package login

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
)

type LoginPayload struct { // 用于接收payload的结构体
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

// @Summary 登录
// @Description Login
// @Tags user
// @Accept json
// @Produce json
// @Param LoginPayload body handler.LoginPayload true "学号和密码"
// @Success 200 {object} model.Res "{"message":"Authentiaction Success.", "token": string}"
// @Failure 401 {object} error.Error "{"error_code":"20001", "message":"Password or account wrong."} 登录失败, {"error_code":"10001", "message":"User does not Existed!"} 用户不存在"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Bad Request!"} 接收失败"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /login/ [post]

func Login(c *gin.Context) { // 用于登录路由的处理函数
	var data LoginPayload // 声明payload变量，因为BindJSON方法需要接收一个指针进行操作
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
			"err":     err, // 接收过程中的错误视为Bad Request
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
			"message": "Password or account wrong.",
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
