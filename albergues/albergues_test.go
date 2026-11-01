package albergues

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Albergues(t, Nuevo) }
