package register

import (
	"github.com/gin-gonic/gin"
	"github.com/jepril/sharing/model"
)

func Register(c *gin.Context)  {
	var data model.User
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Requset!",
		})
		return
	}
	
	if model.CheckUserByUser_id(data.User_id) {
        c.JSON(401, gin.H{
            "message": "User Already Existed!",
        })
        return
    }

	if model.ConfirmUser(data.User_id, data.Password) {
		model.CreateUser(data.User_id, data.User_name, data.Password)
	}else{
		c.JSON(401, gin.H{
			"message": "wrong id or password",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Create Student Successful!",
	})
	return
}