package view

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"

	//"encoding/json"
	. "fmt"
)

// @Summary 显示用户信息
// @Description 显示用户信息,可以是自己的也可以是别人的
// @Tags user
// @Accept json
// @Produce json
// @Param User.User_id body model.User.User_id true "被查看人的ID"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"messsge":"successfully"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录 or {"error_code":"00002", "message":"wrong mysql."}"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/viewing/ [get]

func View(c *gin.Context) {
	//var Token	string

	Token := c.Request.Header.Get("Token")
	// if  Token = nil	{
	//     handler.SendBadRequest(c)
	//     return
	// }
	Println(Token)
	_, error := model.Token_info(Token)
	if !error {
		c.JSON(401, gin.H{
			"message": "wrong token",
		})
		return
	}
	//user_id , _ := model.Token_info(token)

	var data model.User
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c)
		return
	}

	rows, err := model.Viewing(data.User_id)

	if err != nil {
		//Println("222")
		c.JSON(401, gin.H{
			"message": "wrong mysql",
			"err":     err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": rows,
	})
	return
}
