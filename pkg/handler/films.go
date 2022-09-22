package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/films-app/models"
)

func (h *Handler) createFilm(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input models.Film

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	existedFilmName, err := h.services.Film.CheckFilmExists(input.Name)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if existedFilmName {
		NewErrorResponse(c, http.StatusConflict, errFilmExists.Error())
		return
	}

	existedDirector, err := h.services.Film.CheckDirectorExists(string(input.DirectorId))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !existedDirector {
		NewErrorResponse(c, http.StatusBadRequest, errDirectorNotExists.Error())
		return
	}

	filmID, err := h.services.Film.Create(userID, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"film_id": filmID})
}

func (h *Handler) getFilmByID(c *gin.Context) {

}

func (h *Handler) getWholeFilmsList(c *gin.Context) {

}

func (h *Handler) addFilmToFavourites(c *gin.Context) {

}

func (h *Handler) addFilmToWishlist(c *gin.Context) {

}

func (h *Handler) exportFilmsListToCsv(c *gin.Context) {

}
