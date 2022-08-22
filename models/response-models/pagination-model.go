package responsemodels

import "github.com/bedirhankilic/movieapicase/models"

type PaginationModel struct {
	Rows       []models.Movie `json:"Rows"`
	Page       int            `json:"Page"`
	Limit      int            `json:"Limit"`
	TotalPages int            `json:"TotalPages"`
}
