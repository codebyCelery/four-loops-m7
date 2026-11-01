package salud

import (
	"testing"

	"alerta/contratos"
	"alerta/contratotest"
)

func TestContrato(t *testing.T) {
	contratotest.Salud(t, func(r contratos.Rutas) contratos.Salud { return Nuevo(r) })
}
