package models

type Contactos struct {
	Id              int
	NombreCompleto  string
	NumeroDocumento string
	Direccion       string
	Estado          bool
	IdCargo         Parametros
}
