package deletion

import (
        "github.com/jepril/sharing/handler"
        "github.com/jepril/sharing/model"
		"github.com/gin-gonic/gin"
		."fmt"	
)


func Deletion(c *gin.Context){
	
	Token := c.Request.Header.Get("Token")
	// if  err = nil	{
    //     handler.SendBadRequest(c)
    //     return
    // }
	Println(Token)
	_ , error := model.Token_info(Token)
	if  !error{
		c.JSON(401,gin.H{
			"message":"Wrong Token",
		})
		return
	}

	var data model.Following_fans
	data.Fans_id , _ = model.Token_info(Token)
	//Println(data.Fans_id)
    if err := c.BindJSON(&data); err != nil {
        handler.SendBadRequest(c)
        return
	}
	
	if model.CheckFollowingByFans_id(data.Following_id,data.Fans_id){
		if err :=model.DeleteFollowing(data.Fans_id,data.Following_id);err !=nil {
				c.JSON(400, gin.H{
					"message": err,
				})
		}else{
		c.JSON(200, gin.H{
			"message": "Deleting Successful!",
		})
		}
	}else{
		c.JSON(401, gin.H{
			"message": "Following doesn't Existed!",
		})
	}
	
	return
}