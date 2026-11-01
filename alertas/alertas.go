// Package alertas — M13 Alertas.
// Equipo: ______________________
//
// Estructuras propias: heap, tabla hash, lista circular.
// No se permite usar map, container/* ni sort en este paquete.
// Implementen todos los métodos y corran: go test ./alertas/
package alertas

import (
	"alerta/contratos"
)

type Alertas struct {
	// TODO: los campos de sus estructuras
}

func Nuevo() contratos.Alertas {
	return &Alertas{}
}

func (s *Alertas) Emitir(zonaID, mensaje string, prioridad int) bool { panic("por implementar") }
func (s *Alertas) Enviar() (string, string, error)                   { panic("por implementar") }
func (s *Alertas) Fallo(zonaID, mensaje string)                      { panic("por implementar") }
func (s *Alertas) Reintentar() (string, string, error)               { panic("por implementar") }
func (s *Alertas) Confirmar(zonaID, mensaje string)                  { panic("por implementar") }
