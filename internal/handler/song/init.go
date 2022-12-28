package song

import (
	songUsecase "postgres/internal/usecase/song"

	"github.com/gin-gonic/gin"
)

type SongHandler interface {
	Get(context *gin.Context)
	Create(context *gin.Context)
	GetAllSong(context *gin.Context)
	BatchCreate(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type songHandler struct {
	songUsecase songUsecase.SongUsecase
}

// The function is to initialize the song handler
func NewSongHandler(songUsecase songUsecase.SongUsecase) SongHandler {
	return &songHandler{
		songUsecase: songUsecase,
	}
}
