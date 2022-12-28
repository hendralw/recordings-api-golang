package song

import (
	"net/http"

	"postgres/internal/entity"
	"postgres/internal/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// It will call the function Get in song usecase
func (handler *songHandler) Get(context *gin.Context) {
	// Get id from request param
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	song, err := handler.songUsecase.Get(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

// It will call the function Create in song usecase
func (handler *songHandler) Create(context *gin.Context) {
	var requestBody entity.Song

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	song, err := handler.songUsecase.Create(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

// It will call the function GetAllSong in song usecase
func (handler *songHandler) GetAllSong(context *gin.Context) {
	// Get id from request param
	// albumID := context.Query("album_id")
	// artistID := context.Query("artist_id")

	paramID := uuid.Nil
	isAlbumID := false
	isArtistID := false

	if len(context.Query("album_id")) > 0 {
		paramID = uuid.MustParse(context.Query("album_id"))
		isAlbumID = true
	}

	if len(context.Query("artist_id")) > 0 {
		paramID = uuid.MustParse(context.Query("artist_id"))
		isArtistID = true
	}

	if isAlbumID {
		// Get all songs from usecase
		songs, err := handler.songUsecase.GetAllSongByAlbumID(context, paramID)
		if err != nil {
			res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		res := helper.BuildResponse(true, "OK!", songs)
		context.JSON(http.StatusOK, res)
		return
	}

	if isArtistID {

		// Get all songs from usecase
		songs, err := handler.songUsecase.GetAllSongByArtistID(context, paramID)
		if err != nil {
			res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		res := helper.BuildResponse(true, "OK!", songs)
		context.JSON(http.StatusOK, res)
		return
	}

	// Get all songs from usecase
	songs, err := handler.songUsecase.GetAllSong(context)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", songs)
	context.JSON(http.StatusOK, res)
}

// It will call the function BatchCreate in song usecase
func (handler *songHandler) BatchCreate(context *gin.Context) {
	var requestBody []entity.Song

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	songs, err := handler.songUsecase.BatchCreate(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", songs)
	context.JSON(http.StatusOK, res)
}

// It will call the function Update in song usecase
func (handler *songHandler) Update(context *gin.Context) {
	var requestBody entity.Song

	// Get id from request param
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get request body from user
	err = context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Set id from params
	requestBody.SongID = id

	// Call the usecase
	song, err := handler.songUsecase.Update(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

// It will call the function Delete in song usecase
func (handler *songHandler) Delete(context *gin.Context) {
	// Get id from request param
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.songUsecase.Delete(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
