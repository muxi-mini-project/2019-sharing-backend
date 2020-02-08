package following

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jepril/sharing/handler"
	"github.com/jepril/sharing/model"
)

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

	model.CreateFollowing(data.Fans_id, data.Following_id)

	c.JSON(200, gin.H{
		"message": "following Successful!",
	})
	return
}
