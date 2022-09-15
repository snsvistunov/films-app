package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/films-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
		auth.POST("/signout", h.signOut)
	}

	films := router.Group("/films", h.userIdentity)
	{
		films.POST("/favourites", h.addFilmToFavourites)
		films.POST("/wishlist", h.addFilmToWishlist)
		films.GET("/list", h.getWholeFilmsList)
		films.GET("/export", h.exportFilmsListToCsv)
		films.POST("/film", h.createFilm)
		films.GET("/film", h.getFilmByID)
	}

	return router
}
