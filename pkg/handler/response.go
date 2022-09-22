package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	errUserExists                 = errors.New("user exist in db")
	errUserIncorrectLoginPassword = errors.New("incorrect user's login or password")
	errEmptyHeader                = errors.New("empty auth header")
	errInvalidAuthHeader          = errors.New("invalid auth header")
	errEmptyToken                 = errors.New("token is empty")
	errCantLogout                 = errors.New("can't sign out")
	errForbidden                  = errors.New("permission denied")
	errUnknown                    = errors.New("unknown error")
	errUserdIDNotFound            = errors.New("user id not found")
	errUserdIDType                = errors.New("user id has invalid type")
	errFilmExists                 = errors.New("film exist in db")
	errDirectorNotExists          = errors.New("director don't exist")
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
