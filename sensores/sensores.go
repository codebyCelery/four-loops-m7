// Package sensores — M3 Sensores.
// Equipo: ______________________
//
// Estructuras propias: BST, heap.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./sensores/
package sensores

import (
	"time"

	"alerta/contratos"
	"alerta/tipos"
)

type Sensores struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Sensores {
	return &Sensores{}
}

func (s *Sensores) Registrar(l tipos.Lectura) { panic("por implementar") }
func (s *Sensores) Lecturas(estacionID string, desde, hasta time.Time) []tipos.Lectura {
	panic("por implementar")
}
func (s *Sensores) EnAlerta(k int) []string { panic("por implementar") }
