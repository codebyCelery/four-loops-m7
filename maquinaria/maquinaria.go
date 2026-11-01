// Package maquinaria — M11 Maquinaria.
// Equipo: ______________________
//
// Estructuras propias: heap, tabla hash (consulta el grafo de M4).
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./maquinaria/
package maquinaria

import (
	"alerta/contratos"
)

type Maquinaria struct {
	rutas contratos.Rutas
	// TODO: los campos de sus estructuras
}

func Nuevo(r contratos.Rutas) contratos.Maquinaria {
	return &Maquinaria{rutas: r}
}

func (s *Maquinaria) AgregarZona(zonaID string, poblacion int) { panic("por implementar") }
func (s *Maquinaria) AgregarMaquina(id, baseZonaID string)     { panic("por implementar") }
func (s *Maquinaria) TramoCortado(a, b string)                 { panic("por implementar") }
func (s *Maquinaria) Reabrir() (string, string, string, error) { panic("por implementar") }
func (s *Maquinaria) Terminar(maquinaID string)                { panic("por implementar") }
