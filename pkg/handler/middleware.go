package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	tokenType           = "Bearer"
)

func (h *Handler) checkAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		return "", errEmptyHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != tokenType {
		return "", errInvalidAuthHeader
	}

	if len(headerParts[1]) == 0 {
		return "", errEmptyToken
	}
	token := headerParts[1]
	return token, nil
}
