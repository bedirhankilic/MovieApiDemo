package models

type Movie struct {
	Id          int    `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	IsDeleted   bool   `json:"isdeleted"`
}
