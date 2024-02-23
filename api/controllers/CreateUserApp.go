package controllers

import (
	"SimonBK_Login_APP/domain/models"
	"SimonBK_Login_APP/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Crea un nuevo usuario de aplicación
// @Description Añade un nuevo usuario de aplicación a la base de datos
// @Tags appUser
// @Accept json
// @Produce json
// @Param vehiculo body swagger.AplicationUser true "AplicationUser"
// @Success 200 {string} string "ID o mensaje de éxito del AplicationUser creado"
// @Failure 400 {object} map[string]string"
// @Failure 500 {object} map[string]string
// @Router /userApp/create/ [post]
func CreateUserAppController(c *gin.Context) {
	// Obtener FkCompany y FkCustomer del contexto de Gin
	FkCompany, _ := c.Get("FkCompany")
	FkCustomer, _ := c.Get("FkCustomer")

	// Validar y decodificar el cuerpo de la solicitud
	var userApp models.AplicationUser
	if err := c.ShouldBindJSON(&userApp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar a CreateUserApp
	password, err := services.CreateUserApp(userApp.UserNameApp, FkCompany.(*uint), FkCustomer.(*uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"password": password})
}
