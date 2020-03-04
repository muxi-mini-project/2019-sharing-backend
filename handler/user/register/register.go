package register

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
)

// @Summary 注册
// @Description Register
// @Tags user
// @Accept json
// @Produce json
// @Param User.User_id body model.User.User_id true "学号"
// @Param User.Password body model.User.Password true "密码"
// @Param User.User_name body model.User.User_name true "昵称"
// @Success 200 {object} model.Res "{"msg":"Create Student Successful!"} 成功"
// @Failure 401 {object} error.Error "{"error_code":"20001", "message":"Password or account wrong."} 登录失败, {"error_code":"00002", "message":"wrong mysql."}"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Bad Request!"} 接收失败"or{"error_code":"10001", "message":"User does not Existed!"} 用户不存在"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /login/ [post]

func Register(c *gin.Context) {
	var data model.User
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Requset!",
		})
		return
	}

	if model.CheckUserByUser_id(data.User_id) {
		c.JSON(400, gin.H{
			"message": "User Already Existed!",
		})
		return
	}

	if model.ConfirmUser(data.User_id, data.Password) {
		err :=model.CreateUser(data.User_id, data.User_name, data.Password)
		if err !=nil{
			c.JSON(401, gin.H{
				"message": "wrong mysql",
			})
			return
		}
	} else {
		c.JSON(401, gin.H{
			"message": "wrong id or password",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Create Student Successful!",
	})
	return
}
