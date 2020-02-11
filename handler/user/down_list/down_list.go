package down_list

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"

	//"encoding/json"
	. "fmt"
)

func DownList(c *gin.Context) {
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

	var data model.User
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c)
		return
	}

	if !model.CheckUserByUser_id(data.User_id) {
		c.JSON(401, gin.H{
			"message": "User doesn't Existed!",
		})
		return
	}

	fid, err := model.GetDownFileid(data.User_id)
	rows, err := model.List(fid)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":      "successfully",
		"collect_list": rows,
	})
	return
}
