// Package rutas — M4 Rutas y aislamiento.
// Equipo: ______________________
//
// Estructuras propias: grafo con lista de adyacencia, cola (BFS), pila (DFS). Para traducir nombre de zona a índice basta una búsqueda lineal: son pocas zonas.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./rutas/
package rutas

import (
	"alerta/contratos"
)

type Rutas struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Rutas {
	return &Rutas{}
}

func (g *Rutas) AgregarTramo(a, b string)    { panic("por implementar") }
func (g *Rutas) CortarTramo(a, b string)     { panic("por implementar") }
func (g *Rutas) Saltos(a, b string) int      { panic("por implementar") }
func (g *Rutas) Camino(a, b string) []string { panic("por implementar") }
func (g *Rutas) Aislados() [][]string        { panic("por implementar") }
