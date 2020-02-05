package following_list

import (
        "github.com/jepril/sharing/handler"
        "github.com/jepril/sharing/model"
		"github.com/gin-gonic/gin"
		//"encoding/json"
		."fmt"	
)


func FollowingList(c *gin.Context){

	Token := c.Request.Header.Get("Token")

	Println(Token)
	_ , error := model.Token_info(Token)
	if !error{
		c.JSON(401,gin.H{
			"message":"wrong token",
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

	fid,err := model.GetFollowingid(data.User_id)
	rows,err := model.FollowingList(fid)
	num ,err := model.FansNum(data.User_id)
	
	if err!=nil{
		c.JSON(400, gin.H{
            "message": err,
        })
        return
	}

	c.JSON(200,gin.H{
		"message":"successfully",
		"collect_list": rows,
		"total": num,
	})
	return
}