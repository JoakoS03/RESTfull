package models

type User struct {
	IdUser uint   `json:"idUser"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}
