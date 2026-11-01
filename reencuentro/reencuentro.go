// Package reencuentro — M12 Reencuentro.
// Equipo: ______________________
//
// Estructuras propias: tabla hash, BST, lista enlazada.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./reencuentro/
package reencuentro

import (
	"alerta/contratos"
	"alerta/tipos"
)

type Reencuentro struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Reencuentro {
	return &Reencuentro{}
}

func (s *Reencuentro) Registrar(p tipos.Persona)                  { panic("por implementar") }
func (s *Reencuentro) Buscar(id string) (tipos.Persona, bool)     { panic("por implementar") }
func (s *Reencuentro) Familia(familiaID string) []tipos.Persona   { panic("por implementar") }
func (s *Reencuentro) PorApellido(prefijo string) []tipos.Persona { panic("por implementar") }
