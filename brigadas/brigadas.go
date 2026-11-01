// Package brigadas — M7 Brigadas de rescate.
// Equipo: ______________________
//
// Estructuras propias: cola (emergencias), pila (deshacer), tabla hash (brigada -> estado). Usa Rutas.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./brigadas/
package brigadas

import (
	"alerta/contratos"
	"alerta/tipos"
)

type Brigadas struct {
	// TODO: los campos de sus estructuras
}

func Nuevo(r contratos.Rutas) contratos.Brigadas {
	return &Brigadas{}
}

func (s *Brigadas) AgregarBrigada(id, baseZonaID string) { panic("por implementar") }
func (s *Brigadas) Emergencia(r tipos.Reporte)           { panic("por implementar") }
func (s *Brigadas) AsignarSiguiente() (brigadaID, reporteID string, err error) {
	panic("por implementar")
}
func (s *Brigadas) Deshacer() error          { panic("por implementar") }
func (s *Brigadas) Liberar(brigadaID string) { panic("por implementar") }
