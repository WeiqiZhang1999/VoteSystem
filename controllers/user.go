package controllers

import (
	"GinRanking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 0, name, user, 1)

}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("name", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "Save Error")
		return
	}
	ReturnSuccess(c, 0, "Save Successfully", id, 1)
}

func (u UserController) GetList(c *gin.Context) {
	//logger.Write("log information", "user")

	num1 := 1
	num2 := 2
	num3 := num1 / num2

	ReturnError(c, 4004, num3)
}
