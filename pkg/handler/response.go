package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	errUserExist                  = errors.New("user exist in db")
	errUserIncorrectLoginPassword = errors.New("incorrect user's login or password")
	errEmptyHeader                = errors.New("empty auth header")
	errInvalidAuthHeader          = errors.New("invalid auth header")
	errEmptyToken                 = errors.New("token is empty")
	errCantLogout                 = errors.New("can't sign out")
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Message: message,
	})
}
