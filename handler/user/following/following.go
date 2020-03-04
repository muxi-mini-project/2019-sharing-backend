package following

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
)

// @Summary 关注
// @Description Following
// @Tags user
// @Accept json
// @Produce json
// @Param Following_fans.Following_id body model.Following_fans.Following_id true "关注对象id"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"} 成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"wrong token"} 身份认证失败 重新登录 or {"error_code":"00002", "message":"wrong mysql."}"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/following/ [post]

func Following(c *gin.Context) {

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

	var data model.Following_fans
	data.Fans_id, _ = model.Token_info(Token)
	Println(data.Fans_id)
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c)
		return
	}

	//model.CreateFollowing(data.Fans_id, data.Following_id)
	if err := model.CreateFollowing(data.Fans_id, data.Following_id); err != nil {
		//Println("222")
		c.JSON(401, gin.H{
			"message": "wrong mysql",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "following Successful!",
	})
	return
}
