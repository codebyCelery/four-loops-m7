// Package reportes — M1 Reportes.
// Equipo: ______________________
//
// Estructuras propias: cola, tabla hash, BST.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./reportes/
package reportes

import (
	"time"

	"alerta/contratos"
	"alerta/tipos"
)

type Reportes struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Reportes {
	return &Reportes{}
}

func (s *Reportes) Registrar(r tipos.Reporte) bool                    { panic("por implementar") }
func (s *Reportes) Siguiente() (tipos.Reporte, bool)                  { panic("por implementar") }
func (s *Reportes) Buscar(id string) (tipos.Reporte, bool)            { panic("por implementar") }
func (s *Reportes) EntreHoras(desde, hasta time.Time) []tipos.Reporte { panic("por implementar") }
