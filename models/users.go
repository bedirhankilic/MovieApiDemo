package models

type User struct {
	Id       int    `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty" gorm:"unique"`
	Username string `json:"username,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
}
