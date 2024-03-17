package router

import (
	"GinRanking/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {

	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/info", controllers.GetUserInfo)
		user.POST("/list", controllers.GerList)

		user.PUT("/add", func(context *gin.Context) {
			context.String(http.StatusOK, "user add")
		})

		user.DELETE("/delete", func(context *gin.Context) {
			context.String(http.StatusOK, "user delete")
		})
	}

	return r
}
