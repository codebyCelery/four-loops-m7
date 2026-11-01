package ayuda

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Ayuda(t, Nuevo) }
