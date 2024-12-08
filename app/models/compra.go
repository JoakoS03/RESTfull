package models

type Compra struct {
	IdCompra    uint    `json:"idCompra"`
	IdItem      uint    `json:"idItem"`
	IdUser      uint    `json:"idUser"`
	PrecioTotal float64 `json:"precioTotal"`
	Cant        uint    `json:"cant"`
}
