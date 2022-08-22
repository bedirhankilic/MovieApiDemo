package businesslayer

import (
	"math"

	constants "github.com/bedirhankilic/movieapicase/constants"
	dataaccesslayer "github.com/bedirhankilic/movieapicase/data-access-layer"
	"github.com/bedirhankilic/movieapicase/models"
	requestmodels "github.com/bedirhankilic/movieapicase/models/request-models"
	responsemodels "github.com/bedirhankilic/movieapicase/models/response-models"
	"gorm.io/gorm"
)

type IMovieService interface {
	GetAllMovies(limit int, page int) (responsemodels.PaginationModel, constants.ErrorCode)
	GetMovieWithId(movieId int) (models.Movie, constants.ErrorCode)
	DeleteMovie(movieId int) constants.ErrorCode
	AddMovie(item requestmodels.AddMovie) constants.ErrorCode
	UpdateMovie(item requestmodels.UpdateMovie) constants.ErrorCode
}

type MovieService struct {
	Db      *gorm.DB
	DbError error
}

func NewMovieService() *MovieService {
	db, err := dataaccesslayer.DbContext()
	return &MovieService{
		Db:      db,
		DbError: err,
	}
}

func (service MovieService) GetAllMovies(limit int, page int) (responsemodels.PaginationModel, constants.ErrorCode) {
	res := responsemodels.PaginationModel{
		Page:  page,
		Limit: limit,
	}
	if service.DbError != nil {
		return res, constants.DB_ERROR
	}

	pageCalc := page - 1

	if pageCalc < 0 {
		pageCalc = 0
	}
	offset := pageCalc * limit

	var list []models.Movie

	var totalRows int64 = 0
	countErr := service.Db.Model(&models.Movie{}).Where("is_deleted=false").Count(&totalRows)
	if countErr.Error != nil {
		return res, constants.DB_ERROR
	}

	if totalRows == 0 {
		return res, constants.OK
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	if pageCalc > totalPages {
		return res, constants.NOT_FOUND
	}

	result := service.Db.Where("is_deleted=false").Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return res, constants.UNKOWN_ERROR
	}

	res.Rows = list
	res.TotalPages = totalPages

	return res, constants.OK
}

func (service MovieService) AddMovie(item requestmodels.AddMovie) constants.ErrorCode {
	if service.DbError != nil {
		return constants.DB_ERROR
	}

	if item.Description == "" || item.Name == "" || item.Type == "" {
		return constants.UNKOWN_ERROR
	}
	movie := models.Movie{
		Name:        item.Name,
		Description: item.Description,
		Type:        item.Type,
		IsDeleted:   false,
	}
	result := service.Db.Create(&movie)
	if result.Error != nil {
		return constants.UNKOWN_ERROR
	}
	return constants.OK
}

func (service MovieService) DeleteMovie(movieId int) constants.ErrorCode {
	if service.DbError != nil {
		return constants.DB_ERROR
	}
	var movie models.Movie
	result := service.Db.Where("id = ? and is_deleted=false", movieId).First(&movie)
	if result.RowsAffected == 0 {
		return constants.NOT_FOUND
	}
	movie.IsDeleted = true
	saveResult := service.Db.Save(&movie)
	if saveResult.RowsAffected == 0 {
		return constants.DB_ERROR
	}

	return constants.OK
}

func (service MovieService) UpdateMovie(item requestmodels.UpdateMovie) constants.ErrorCode {
	if service.DbError != nil {
		return constants.DB_ERROR
	}
	var movie models.Movie
	result := service.Db.Where("id = ? and is_deleted=false", item.Id).First(&movie)
	if result.RowsAffected == 0 {
		return constants.NOT_FOUND
	}

	if item.Description == "" || item.Name == "" || item.Type == "" {
		return constants.UNKOWN_ERROR
	}

	movie.Description = item.Description
	movie.Name = item.Name
	movie.Type = item.Type

	saveResult := service.Db.Save(&movie)
	if saveResult.RowsAffected == 0 {
		return constants.DB_ERROR
	}
	return constants.OK
}
