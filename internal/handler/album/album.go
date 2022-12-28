package album

import (
	"net/http"

	"postgres/internal/entity"
	"postgres/internal/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// It will call the function Get in album usecase
func (handler *albumHandler) Get(context *gin.Context) {
	// Get id from request param
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	album, err := handler.albumUsecase.Get(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", album)
	context.JSON(http.StatusOK, res)
}

// It will call the function Create in album usecase
func (handler *albumHandler) Create(context *gin.Context) {
	var requestBody entity.Album

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	album, err := handler.albumUsecase.Create(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", album)
	context.JSON(http.StatusOK, res)
}

// It will call the function GetAllAlbum in album usecase
func (handler *albumHandler) GetAllAlbum(context *gin.Context) {
	// Get id from request param
	id, err := uuid.Parse(context.Query("artist_id"))

	if id.String() != "" && id.String() != uuid.Nil.String() {
		if err != nil {
			res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		albums, err := handler.albumUsecase.GetAllAlbumByArtistID(context, id)
		if err != nil {
			res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		res := helper.BuildResponse(true, "OK!", albums)
		context.JSON(http.StatusOK, res)
		return
	}

	// Get all albums from usecase
	albums, err := handler.albumUsecase.GetAllAlbum(context)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", albums)
	context.JSON(http.StatusOK, res)
}

// It will call the function BatchCreate in album usecase
func (handler *albumHandler) BatchCreate(context *gin.Context) {
	var requestBody []entity.Album

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	albums, err := handler.albumUsecase.BatchCreate(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", albums)
	context.JSON(http.StatusOK, res)
}

// It will call the function Update in album usecase
func (handler *albumHandler) Update(context *gin.Context) {
	var requestBody entity.Album

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
	requestBody.AlbumID = id

	// Call the usecase
	album, err := handler.albumUsecase.Update(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", album)
	context.JSON(http.StatusOK, res)
}

// It will call the function Delete in album usecase
func (handler *albumHandler) Delete(context *gin.Context) {
	// Get id from request param
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.albumUsecase.Delete(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
