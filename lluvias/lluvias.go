// Package lluvias — M9 Lluvias.
// Equipo: ______________________
//
// Estructuras propias: deque (ventana de 6 h), heap, tabla hash.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./lluvias/
package lluvias

import (
	"time"

	"alerta/contratos"
)

type Lluvias struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Lluvias {
	return &Lluvias{}
}

func (s *Lluvias) Registrar(estacionID, zonaID string, hora time.Time, mm float64) {
	panic("por implementar")
}
func (s *Lluvias) Acumulado(estacionID string) float64 { panic("por implementar") }
func (s *Lluvias) EnRiesgo(k int) []string             { panic("por implementar") }
