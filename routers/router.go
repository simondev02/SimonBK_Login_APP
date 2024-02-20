package routers

import (
	"SimonBK_Login_APP/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// Rutas para el modelo Usuario
	auth := r.Group("/appUser")
	{
		auth.POST("/", controllers.HandleUser)
	}
	return r
}
