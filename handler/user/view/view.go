package view

import (
        "github.com/jepril/sharing/handler"
        "github.com/jepril/sharing/model"
		"github.com/gin-gonic/gin"
		//"encoding/json"
		."fmt"	
)


func View(c *gin.Context){
	//var Token	string
	
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
	//user_id , _ := model.Token_info(token)

	var data model.User
    if err := c.BindJSON(&data); err != nil {
        handler.SendBadRequest(c)
        return
	}

	rows := model.Viewing(data.User_id)

	c.JSON(200,gin.H{
		"message":rows,
	})
	return
}