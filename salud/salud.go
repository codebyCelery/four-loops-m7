// Package salud — M10 Salud.
// Equipo: ______________________
//
// Estructuras propias: heap, cola, tabla hash.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./salud/
package salud

import (
	"alerta/contratos"
	"alerta/tipos"
)

type Salud struct {
	rutas contratos.Rutas
	// TODO: los campos de sus estructuras
}

func Nuevo(r contratos.Rutas) contratos.Salud {
	return &Salud{rutas: r}
}

func (s *Salud) AgregarPuesto(id, zonaID string, camas int) { panic("por implementar") }
func (s *Salud) Ingresar(h tipos.Herido)                    { panic("por implementar") }
func (s *Salud) DerivarSiguiente() (string, string, error)  { panic("por implementar") }
func (s *Salud) Alta(puestoID string) (string, error)       { panic("por implementar") }
