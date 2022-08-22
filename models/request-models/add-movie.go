package requestmodels

type AddMovie struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
