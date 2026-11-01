package alertas

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Alertas(t, Nuevo) }
