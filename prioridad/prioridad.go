// Package prioridad — M2 Prioridad de zonas.
// Equipo: ______________________
//
// Estructuras propias: heap, tabla hash (zona -> posición en el heap).
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./prioridad/
package prioridad

import (
	"alerta/contratos"
)

type Prioridad struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Prioridad {
	return &Prioridad{}
}

func (p *Prioridad) Actualizar(zonaID string, riesgo int)            { panic("por implementar") }
func (p *Prioridad) Siguiente() (zonaID string, riesgo int, ok bool) { panic("por implementar") }
