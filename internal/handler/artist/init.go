package artist

import (
	artistUsecase "postgres/internal/usecase/artist"

	"github.com/gin-gonic/gin"
)

type ArtistHandler interface {
	Get(context *gin.Context)
	Create(context *gin.Context)
	GetAllArtist(context *gin.Context)
	BatchCreate(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type artistHandler struct {
	artistUsecase artistUsecase.ArtistUsecase
}

// The function is to initialize the artist handler
func NewArtistHandler(artistUsecase artistUsecase.ArtistUsecase) ArtistHandler {
	return &artistHandler{
		artistUsecase: artistUsecase,
	}
}
