package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-delivery/app/controllers"
	"go-delivery/docs"
)

var userController = new(controllers.UserController)

func SetupRouter() *gin.Engine {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "go delivery"
	docs.SwaggerInfo.Description = "This is a sample go-delivery server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
