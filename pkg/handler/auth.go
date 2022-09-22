package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/snsvistunov/films-app/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	exist, err := h.services.Authorization.CheckUserExists(input.Login)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if exist {
		NewErrorResponse(c, http.StatusConflict, errUserExists.Error())
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
	var input models.UserSignIn

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusUnauthorized, errUserIncorrectLoginPassword.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *Handler) signOut(c *gin.Context) {
	var token string

	token, err := h.checkAuthHeader(c)
	if err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err := h.services.DeleteToken(token); err != nil {
		logrus.Error(err)
		NewErrorResponse(c, http.StatusUnauthorized, errCantLogout.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"signout": "successfully"})
}
