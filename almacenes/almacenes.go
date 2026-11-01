// Package almacenes — M8 Almacenes de ayuda.
// Equipo: ______________________
//
// Estructuras propias: heap (lotes por vencimiento), tabla hash (producto -> stock), BST (productos por stock).
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./almacenes/
package almacenes

import (
	"alerta/contratos"
	"alerta/tipos"
)

type Almacenes struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Almacenes {
	return &Almacenes{}
}

func (s *Almacenes) Ingresar(l tipos.Lote) { panic("por implementar") }
func (s *Almacenes) Retirar(producto string, cantidad int) ([]tipos.Lote, error) {
	panic("por implementar")
}
func (s *Almacenes) Stock(producto string) int      { panic("por implementar") }
func (s *Almacenes) BajoMinimo(minimo int) []string { panic("por implementar") }
