package controllers

import (
	views "SimonBK_Login_APP/api/views/Inpunts"
	"SimonBK_Login_APP/domain/services"
	"SimonBK_Login_APP/infra/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Actualizar un userApp
// @Description Actualiza un usuario de la aplicación
// @Tags  appUser
// @Accept json
// @Produce json
// @Param id path int true "ID del usuario de la aplicación"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /userApp/{id} [put]
func UpdateAppUserController(c *gin.Context) {
	var UserApp views.AplicationUserInputs

	// 1. Llamamos las variables del contexto y los parametros de entrada
	userAppID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}
	userAppIDPtr := uint(userAppID)
	fkCompany := c.MustGet("FkCompany").(*uint)
	userId := c.MustGet("UserId").(*uint)

	// 2 . Aplicamos los datos en el struct
	UserApp = views.AplicationUserInputs{
		UserAppID: &userAppIDPtr,
		FkCompany: fkCompany,
		UserID:    userId,
	}

	// 3 . Llamamos el servicio
	if _, err := services.UpdateAppUser(db.DBConn, UserApp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
