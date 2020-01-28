package view

import (
        "github.com/jepril/sharing/handler"
        "github.com/jepril/sharing/model"
		"github.com/gin-gonic/gin"
		//"encoding/json"	
)


func View(c *gin.Context){
	var token	string
	if err := c.BindJSON(&token); err != nil {
        handler.SendBadRequest(c)
        return
    }

	_ , error := model.Token_info(token)
	if !error{
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