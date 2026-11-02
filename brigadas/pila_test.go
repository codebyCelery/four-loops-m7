package brigadas

import "testing"

func TestPilaApilarYDesapilar(t *testing.T) {
	p := NuevaPila()

	if p.Len() != 0 {
		t.Errorf("Se esperaba longitud 0, se obtuvo %d", p.Len())
	}

	p.Apilar("A")
	p.Apilar("B")

	if p.Len() != 2 {
		t.Errorf("Se esperaba longitud 2, se obtuvo %d", p.Len())
	}

	tope, err := p.Tope()
	if err != nil || tope != "B" {
		t.Errorf("Tope incorrecto: se esperaba B, se obtuvo %v", tope)
	}

	val, err := p.Desapilar()
	if err != nil || val != "B" {
		t.Errorf("Desapilar incorrecto: se esperaba B, se obtuvo %v", val)
	}

	if p.Len() != 1 {
		t.Errorf("Se esperaba longitud 1, se obtuvo %d", p.Len())
	}
}

func TestPilaVaciaErrores(t *testing.T) {
	p := NuevaPila()

	_, err := p.Desapilar()
	if err == nil {
		t.Error("Se esperaba un error al desapilar una pila vacía")
	}

	_, err = p.Tope()
	if err == nil {
		t.Error("Se esperaba un error al consultar el tope de una pila vacía")
	}
}
