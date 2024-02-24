package controllers

import (
	"SimonBK_Login_APP/api/views"
	"SimonBK_Login_APP/domain/services"
	"SimonBK_Login_APP/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Iniciar sesión
// @Description Autentica a un usuario y devuelve un token de acceso y los detalles del usuario.
// @Tags appUser
// @Accept  json
// @Produce  json
// @Param   login  body  swagger.LoginForm  true  "Credenciales del usuario"
// @Success 200 {object} views.Response "Respuesta exitosa con token y detalles del usuario"
// @Failure 400 "Error: Datos inválidos o problema con el formato del email"
// @Failure 401 "Error: Credenciales incorrectas o intento de inicio de sesión fallido"
// @Failure 500 "Error interno del servidor"
// @Router /login/userApp/ [post]
func LoginUserApp(c *gin.Context) {

	var input views.LoginForm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := services.LoginUserApp(db.DBConn, input.UserNameApp, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !loginResponse.Success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password does not match"})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}
