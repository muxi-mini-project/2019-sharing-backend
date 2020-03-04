package deletion

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
)

// @Summary deletion
// @Description 取消关注
// @Tags user
// @Accept json
// @Produce json
// @Param Following_fans.following_id body model.Following_fans.following_id true "关注者ID"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"messsge":"successfully"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录 or {"error_code":"00002", "message":"wrong mysql."} or {"error_code":"00003", "message":"Following doesn't Existed!"}"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/deletion/ [delete]

func Deletion(c *gin.Context) {

	Token := c.Request.Header.Get("Token")
	// if  err = nil	{
	//     handler.SendBadRequest(c)
	//     return
	// }
	Println(Token)
	_, error := model.Token_info(Token)
	if !error {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	var data model.Following_fans
	data.Fans_id, _ = model.Token_info(Token)
	//Println(data.Fans_id)
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c)
		return
	}

	if model.CheckFollowingByFans_id(data.Following_id, data.Fans_id) {
		if err := model.DeleteFollowing(data.Fans_id, data.Following_id); err != nil {
			c.JSON(401, gin.H{
				"message": "wrong mysql",
				"err":err,
		} else {
			c.JSON(200, gin.H{
				"message": "Deleting Successful!",
			})
		}
	} else {
		c.JSON(401, gin.H{
			"message": "Following doesn't Existed!",
		})
	}

	return
}
