package background

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
)

func Background(c *gin.Context) {

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
	//Println(data.User_id)
	if err := c.BindJSON(&data); err != nil {
		SendBadRequest(c)
		return
	}

	if err := model.Background_modify(data.User_id, data.Background_url); err != nil {
		Println("222")
		c.JSON(403, gin.H{
			"message": "wrong Mysql",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Modify Background Successful!",
	})
	return
}
