package image
import (
	"github.com/gin-gonic/gin"
	."fmt"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
)


func Image(c *gin.Context){
	
	Token := c.Request.Header.Get("Token")
	// if  err = nil	{
    //     handler.SendBadRequest(c)
    //     return
    // }
	Println(Token)
	_ , error := model.Token_info(Token)
	if !error{
		c.JSON(401,gin.H{
			"message":"wrong token",
		})
		return
	}

	var data model.User
	data.User_id , _ = model.Token_info(Token)
	
    if err := c.BindJSON(&data); err != nil {
        handler.SendBadRequest(c)
        return
	}
	
	model.Image_modify(data.User_id,data.Image_url)

	c.JSON(200, gin.H{
		"message": "Modify Image Successful!",
	})
	return
}