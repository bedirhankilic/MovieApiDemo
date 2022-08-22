package responsemodels

type MoviesListItem struct {
	GenericResponse
	Data PaginationModel `json:"Data"`
}
