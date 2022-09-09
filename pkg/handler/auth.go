package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/films-app/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	exist, err := h.services.Authorization.CheckUserExist(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if exist {
		NewErrorResponse(c, http.StatusConflict, errUserExist.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {

}

func (h *Handler) signOut(c *gin.Context) {

}
