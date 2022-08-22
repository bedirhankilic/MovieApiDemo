package responsemodels

type TokenResponse struct {
	GenericResponse
	Data string `json:"Data"`
}
