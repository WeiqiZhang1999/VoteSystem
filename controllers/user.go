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

func (u UserController) GetList(c *gin.Context) {
	//logger.Write("log information", "user")

	num1 := 1
	num2 := 0
	num3 := num1 / num2

	ReturnError(c, 4004, num3)
}
