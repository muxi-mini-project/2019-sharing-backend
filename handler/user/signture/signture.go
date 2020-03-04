package signture

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
)

// @Summary signture
// @Description 修改签名
// @Tags user
// @Accept json
// @Produce json
// @Param User.Signture body model.User.Signture true "签名"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"messsge":"successfully"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录 or {"error_code":"00002", "message":"wrong mysql."} "
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/signture/ [get]

func Signture(c *gin.Context) {

	Token := c.Request.Header.Get("Token")
	// if  err = nil	{
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

	var data model.User
	data.User_id, _ = model.Token_info(Token)

	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c)
		return
	}

	//model.Signture_modify(data.User_id, data.Signture)
	if err := model.Signture_modify(data.User_id, data.Signture); err != nil {
		//Println("222")
		c.JSON(401, gin.H{
			"message": "wrong mysql",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Modify Image Successful!",
	})
	return
}
