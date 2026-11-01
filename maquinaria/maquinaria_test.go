package maquinaria

import (
	"testing"

	"alerta/contratos"
	"alerta/contratotest"
)

func TestContrato(t *testing.T) {
	contratotest.Maquinaria(t, func(r contratos.Rutas) contratos.Maquinaria { return Nuevo(r) })
}
