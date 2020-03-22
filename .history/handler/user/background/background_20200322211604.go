package background

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
)

// @Summary background
// @Description 修改背景
// @Tags user
// @Accept json
// @Produce json
// @Param User.Background_url body model.User.Background_url true "背景url"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"messsge":"successfully"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录 or {"error_code":"00002", "message":"wrong mysql."} "
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/background/ [put]

func Background(c *gin.Context) {

	Token := c.Request.Header.Get("Token")
	// if  err = nil	{
	//     handler.SendBadRequest(c)
	//     return
	// }
	//Println(Token)
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
		handler.SendBadRequest(c)
		return
	}

	//var tmpfile model.File
	fileid, _ := strconv.Atoi(c.Param("fileid"))
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message": "Bad Request!",
		})
		return
	}
	dataLen := header.Size

	data.Background_url, err := model.Uploadfile(header.Filename, uint32(fileid), file, dataLen)
	log.Print(fileid)

	if err != nil {
		c.JSON(404,gin.H{
			"message": "生成地址失败",
		})
		return
	}

	if err := model.Background_modify(data.User_id, data.Background_url); err != nil {
		//Println("222")
		
		c.JSON(401, gin.H{
			"message": "wrong mysql",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Modify Background Successful!",
	})
	return
}
