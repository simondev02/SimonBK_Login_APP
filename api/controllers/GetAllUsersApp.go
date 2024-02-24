package controllers

import (
	"SimonBK_Login_APP/domain/services"
	"SimonBK_Login_APP/infra/db"
	"log"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// GetAllDispositivosHandler - Controlador para obtener todos los ususarios de palicacion
// @Summary Obtiene todos Ususarios de palicacion
// @Obtiene todos los Usuarios de aplicacion
// @Tags appUser
// @Accept json
// @Produce json
// @Param page query int false "Número de página para la paginación" default(1)
// @Security ApiKeyAuth
// @Success 200 {array} map[string]string
// @Failure 500 {object} map[string]string
// @Router /userApp/ [get]
func GetAllAppUserController(c *gin.Context) {
	response, err := services.GetAllAppUser(db.DBConn)
	if err != nil {
		log.Print("[GetAllAppUserController]", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, response)
}
