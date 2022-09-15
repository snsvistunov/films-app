package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createFilm(c *gin.Context) {

}

func (h *Handler) getFilmByID(c *gin.Context) {

}

func (h *Handler) getWholeFilmsList(c *gin.Context) {

}

func (h *Handler) addFilmToFavourites(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id by context": id,
	})
}

func (h *Handler) addFilmToWishlist(c *gin.Context) {

}

func (h *Handler) exportFilmsListToCsv(c *gin.Context) {

}
