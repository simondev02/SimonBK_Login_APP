package routers

import (
	"SimonBK_Login_APP/api/controllers"
	"SimonBK_Login_APP/middleware" // Asegúrate de importar tu middleware

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
	auth := r.Group("/login")
	{
		auth.POST("/userApp/", controllers.LoginUserApp)
	}

	// Rutas para Usuario aplicación
	auth = r.Group("/userApp")
	{
		// Aplicar el middleware de validación de token solo a este grupo
		auth.Use(middleware.ValidateTokenMiddleware())
		auth.POST("/create/", controllers.CreateUserAppController)
		auth.GET("/", controllers.GetAllAppUserController)
		auth.PUT("/", controllers.UpdateAppUserController)
	}

	return r
}

//
