package router

import (
	"GinRanking/controllers"
	"GinRanking/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)

		user.PUT("/add", func(context *gin.Context) {
			context.String(http.StatusOK, "user add")
		})

		user.DELETE("/delete", func(context *gin.Context) {
			context.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
