package brigadas

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Brigadas(t, Nuevo) }
