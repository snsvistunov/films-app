package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	tokenType           = "Bearer"
	userCtx             = "userID"
	admin               = "administrator"
)

func (h *Handler) userIdentity(c *gin.Context) {
	token, err := h.checkAuthHeader(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userID, err := h.services.GetUserID(token)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)

}

func (h *Handler) adminIdentity(c *gin.Context) {
	val, exist := c.Get(userCtx)
	if !exist {
		NewErrorResponse(c, http.StatusInternalServerError, errUnknown.Error())
		return
	}
	userID := val.(string)
	isAdmin, err := h.services.IsAdmin(userID, admin)
	if err != nil {
		NewErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	if !isAdmin {
		NewErrorResponse(c, http.StatusForbidden, errForbidden.Error())
		return
	}
}

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

func getUserID(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return "", errUserdIDNotFound
	}

	userID, ok := id.(string)
	if !ok {
		return "", errUserdIDType
	}

	return userID, nil
}
