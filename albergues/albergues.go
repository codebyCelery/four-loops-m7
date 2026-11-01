// Package albergues — M5 Albergues.
// Equipo: ______________________
//
// Estructuras propias: heap (lista de espera), tabla hash (familia -> albergue). Usa Rutas.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./albergues/
package albergues

import (
	"alerta/contratos"
	"alerta/tipos"
)

type Albergues struct {
	// TODO: los campos de sus estructuras
}

func Nuevo(r contratos.Rutas) contratos.Albergues {
	return &Albergues{}
}

func (s *Albergues) AgregarAlbergue(a tipos.Albergue) { panic("por implementar") }
func (s *Albergues) Solicitar(f tipos.Familia)        { panic("por implementar") }
func (s *Albergues) AsignarSiguiente() (familiaID, albergueID string, err error) {
	panic("por implementar")
}
func (s *Albergues) DondeEsta(familiaID string) (string, bool) { panic("por implementar") }
