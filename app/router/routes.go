package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	_ = r.Group("/api")
	{
		_ = r.Group("/vi")
		{
			users := r.Group("/users")
			{
				users.GET()
				users.GET()
				users.POST()
				users.PUT()
				users.DELETE()
			}
		}
	}

	return r
}
