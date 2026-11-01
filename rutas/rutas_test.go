package rutas

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Rutas(t, Nuevo) }
