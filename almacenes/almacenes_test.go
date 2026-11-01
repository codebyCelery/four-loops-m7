package almacenes

import (
	"testing"

	"alerta/contratotest"
)

func TestContrato(t *testing.T) { contratotest.Almacenes(t, Nuevo) }
