// Package ayuda — M6 Ayuda humanitaria.
// Equipo: ______________________
//
// Estructuras propias: heap (necesidad por zona), cola (camiones), tabla hash. Usa Rutas y Almacenes.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./ayuda/
package ayuda

import (
	"alerta/contratos"
)

type Ayuda struct {
	// TODO: los campos de sus estructuras
}

func Nuevo(r contratos.Rutas, a contratos.Almacenes) contratos.Ayuda {
	return &Ayuda{}
}

func (s *Ayuda) AgregarCamion(id, baseZonaID string)                    { panic("por implementar") }
func (s *Ayuda) RegistrarNecesidad(zonaID string, cantidad int)         { panic("por implementar") }
func (s *Ayuda) Despachar() (camionID string, ruta []string, err error) { panic("por implementar") }
