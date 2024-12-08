package models

type Items struct {
	IdItem uint   `json:"idItem"`
	Nombre string `json:"nombre"`
	Precio string `json:"precio"`
	Stock  uint   `json:"stock"`
}
