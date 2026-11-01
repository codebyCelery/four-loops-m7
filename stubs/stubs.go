// Package stubs: versiones falsas de las dependencias, para que M5, M6 y M7
// puedan desarrollar y probar sin esperar a M4 y M8. No resuelven nada de verdad.
package stubs

import "alerta/tipos"

// Rutas falsa: todo está a 3 tramos y siempre hay camino directo.
type Rutas struct{}

func (Rutas) AgregarTramo(a, b string) {}
func (Rutas) CortarTramo(a, b string)  {}
func (Rutas) Saltos(a, b string) int {
	if a == b {
		return 0
	}
	return 3
}
func (Rutas) Camino(a, b string) []string { return []string{a, b} }
func (Rutas) Aislados() [][]string        { return nil }

// Almacenes falso: el stock nunca se acaba.
type Almacenes struct{}

func (Almacenes) Ingresar(l tipos.Lote) {}
func (Almacenes) Retirar(p string, c int) ([]tipos.Lote, error) {
	return []tipos.Lote{{ID: "stub", Producto: p, Cantidad: c}}, nil
}
func (Almacenes) Stock(p string) int        { return 1 << 30 }
func (Almacenes) BajoMinimo(m int) []string { return nil }
