package router

import (
	"github.com/gin-gonic/gin"
	"go-delivery/app/controllers"
)

var userController = new(controllers.UserController)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	_ = r.Group("/api")
	{
		_ = r.Group("/vi")
		{
			users := r.Group("/users")
			{
				users.GET("", userController.List)
				users.GET(":id", userController.Detail)
				users.POST("", userController.Create)
				users.PUT(":id", userController.Update)
				users.DELETE(":id", userController.Delete)
			}
		}
	}

	return r
}
