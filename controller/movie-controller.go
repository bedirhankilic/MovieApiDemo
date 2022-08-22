package controller

import (
	"net/http"
	"strconv"

	businesslayer "github.com/bedirhankilic/movieapicase/business-layer"
	"github.com/bedirhankilic/movieapicase/constants"
	requestmodels "github.com/bedirhankilic/movieapicase/models/request-models"
	responsemodels "github.com/bedirhankilic/movieapicase/models/response-models"
	"github.com/gin-gonic/gin"
)

// GetAllMovies godoc
// @Summary      Get all movies list
// @Description  This EP supports pagination (GetAllMovies?page=2&limit=3)
// @Tags         movie
// @Accept       json
// @Produce      json
// @Param page query int false "page number"
// @Param limit query int false "limit count"
// @Success      200  {object}  responsemodels.MoviesListItem{}
// @Failure      400  {object}  responsemodels.GenericResponse{}
// @Router       /GetAllMovies/ [get]
func GetAllMovies(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 5
	}

	movieService := businesslayer.NewMovieService()

	response, errCode := movieService.GetAllMovies(limit, page)
	res := responsemodels.MoviesListItem{
		GenericResponse: responsemodels.GenericResponse{ResponseCode: errCode},
		Data:            response,
	}
	if errCode != constants.OK {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// AddMovie godoc
// @Summary      Add Movie EP
// @Description  This method addes a new movie record
// @Tags         movie
// @Accept       json
// @Produce      json
// @Param addMovie body requestmodels.AddMovie{} true "requestmodels.AddMovie Add Movie Item"
// @Success      200  {object}  responsemodels.GenericResponse{}
// @Failure      400  {object}  responsemodels.GenericResponse{}
// @Failure      401  {object}  responsemodels.GenericResponse{}
// @Security     ApiKeyAuth
// @Router       /AddMovie/ [put]
func AddMovie(ctx *gin.Context) {
	var req requestmodels.AddMovie

	if ctx.ShouldBindJSON(&req) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: constants.WRONG_REQ_FORMAT,
		})
		return
	}

	movieService := businesslayer.NewMovieService()
	errCode := movieService.AddMovie(req)
	if errCode != constants.OK {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: errCode,
		})
		return
	}
	ctx.JSON(http.StatusOK, responsemodels.GenericResponse{
		ResponseCode: errCode,
	})
}

// DeleteMovie godoc
// @Summary      Delete movie
// @Description  This method deletes a movie
// @Tags         movie
// @Accept       json
// @Produce      json
// @Param movieId path int true "movie id must be integer"
// @Success      200  {object}  responsemodels.GenericResponse{}
// @Failure      400  {object}  responsemodels.GenericResponse{}
// @Failure      401  {object}  responsemodels.GenericResponse{}
// @Security     ApiKeyAuth
// @Router       /DeleteMovie/{movieId} [post]
func DeleteMovie(ctx *gin.Context) {
	movieId := ctx.Param("movieId")
	movieIdInt, errParse := strconv.Atoi(movieId)
	if errParse != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responsemodels.GenericResponse{
			ResponseCode: constants.WRONG_REQ_FORMAT,
		})
		return
	}

	movieService := businesslayer.NewMovieService()
	errCode := movieService.DeleteMovie(movieIdInt)
	if errCode != constants.OK {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: errCode,
		})
		return
	}
	ctx.JSON(http.StatusOK, responsemodels.GenericResponse{
		ResponseCode: errCode,
	})
}

// UpdateMovie godoc
// @Summary      Update Movie EP
// @Description  This method update a movie record
// @Tags         movie
// @Accept       json
// @Produce      json
// @Param updateMovie body requestmodels.UpdateMovie{} true "requestmodels.UpdateMovie Updae Movie Item"
// @Success      200  {object}  responsemodels.GenericResponse{}
// @Failure      400  {object}  responsemodels.GenericResponse{}
// @Failure      401  {object}  responsemodels.GenericResponse{}
// @Security     ApiKeyAuth
// @Router       /UpdateMovie/ [put]
func UpdateMovie(ctx *gin.Context) {
	var req requestmodels.UpdateMovie
	if ctx.ShouldBindJSON(&req) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: constants.WRONG_REQ_FORMAT,
		})
		return
	}

	movieService := businesslayer.NewMovieService()
	errCode := movieService.UpdateMovie(req)
	if errCode != constants.OK {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responsemodels.GenericResponse{
			ResponseCode: errCode,
		})
		return
	}
	ctx.JSON(http.StatusOK, responsemodels.GenericResponse{
		ResponseCode: errCode,
	})
}
