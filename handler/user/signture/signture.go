package signture

import (
        "github.com/jepril/sharing/handler"
        "github.com/jepril/sharing/model"
		"github.com/gin-gonic/gin"
		."fmt"	
)


func Signture(c *gin.Context){
	
	Token := c.Request.Header.Get("Token")
	// if  err = nil	{
    //     handler.SendBadRequest(c)
    //     return
    // }
	Println(Token)
	_ , error := model.Token_info(Token)
	if error{
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

	model.Signture_modify(data.User_id,data.Signture)

	c.JSON(200, gin.H{
		"message": "Modify Image Successful!",
	})
	return
}