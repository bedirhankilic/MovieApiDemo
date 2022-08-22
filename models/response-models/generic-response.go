package responsemodels

import "github.com/bedirhankilic/movieapicase/constants"

type GenericResponse struct {
	ResponseCode constants.ErrorCode `json:"StatusCode"`
}
