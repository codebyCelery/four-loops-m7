package reencuentro

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Reencuentro(t, Nuevo) }
