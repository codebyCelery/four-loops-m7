package lluvias

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Lluvias(t, Nuevo) }
