package models

import (
	"time"

	"gorm.io/gorm"
)

type Reserva struct {
	gorm.Model
	Nombre    string
	Fecha     time.Time
	ClienteID int
	Cliente   Cliente
}
type Cliente struct {
	gorm.Model
	Nombre string
	PaisID int
	Pais   Pais
}

type Pais struct {
	gorm.Model
	Nombre string
	Codigo string
}
