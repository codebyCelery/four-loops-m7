package brigadas

import (
	"alerta/tipos"
	"fmt"
	"testing"
)

func TestDequeVacio(t *testing.T) {
	d := NuevoDeque()

	if d.Len() != 0 {
		t.Error("El deque debería estar vacío")
	}

	_, ok := d.QuitarFrente()
	if ok {
		t.Error("No debería quitar elementos")
	}

	_, ok = d.Frente()
	if ok {
		t.Error("No debería encontrar elementos")
	}
}

func TestDequeFIFO(t *testing.T) {
	d := NuevoDeque()

	reportes := []tipos.Reporte{
		{ID: "A", ZonaID: "Z1", Tipo: "HUAICO"},
		{ID: "B", ZonaID: "Z2", Tipo: "INUNDACION"},
		{ID: "C", ZonaID: "Z3", Tipo: "DESBORDE"},
	}

	for _, r := range reportes {
		d.InsertarFinal(r)
	}

	for _, esperado := range reportes {
		obtenido, ok := d.QuitarFrente()

		if !ok {
			t.Fatal("No se pudo quitar el reporte")
		}

		if obtenido.ID != esperado.ID {
			t.Errorf("Esperaba %s, pero salió %s",
				esperado.ID, obtenido.ID)
		}
	}
}

func TestDequeInsertarFrente(t *testing.T) {
	d := NuevoDeque()

	reporteA := tipos.Reporte{ID: "A"}
	reporteB := tipos.Reporte{ID: "B"}
	reporteC := tipos.Reporte{ID: "C"}

	d.InsertarFinal(reporteA)
	d.InsertarFinal(reporteB)
	d.InsertarFrente(reporteC)

	esperados := []string{"C", "A", "B"}

	for _, esperado := range esperados {
		obtenido, ok := d.QuitarFrente()

		if !ok {
			t.Fatal("No se pudo quitar el reporte")
		}

		if obtenido.ID != esperado {
			t.Errorf("Esperaba %s, pero salió %s",
				esperado, obtenido.ID)
		}
	}
}

func TestDequeVueltaDelIndice(t *testing.T) {
	d := NuevoDeque()

	for i := 0; i < 100; i++ {
		reporte := tipos.Reporte{
			ID: fmt.Sprintf("R%d", i),
		}

		d.InsertarFinal(reporte)

		obtenido, ok := d.QuitarFrente()

		if !ok {
			t.Fatal("No se pudo quitar el reporte")
		}

		if obtenido.ID != reporte.ID {
			t.Errorf("Esperaba %s, pero salió %s",
				reporte.ID, obtenido.ID)
		}
	}

	if d.Len() != 0 {
		t.Error("El deque debería terminar vacío")
	}
}

func TestDequeCrece(t *testing.T) {
	d := NuevoDeque()

	for i := 0; i < 100; i++ {
		reporte := tipos.Reporte{
			ID: fmt.Sprintf("R%d", i),
		}

		d.InsertarFinal(reporte)
	}

	if d.Len() != 100 {
		t.Errorf("Esperaba 100 elementos, pero hay %d", d.Len())
	}

	for i := 0; i < 100; i++ {
		obtenido, ok := d.QuitarFrente()

		if !ok {
			t.Fatal("No se pudo quitar el reporte")
		}

		esperado := fmt.Sprintf("R%d", i)

		if obtenido.ID != esperado {
			t.Errorf("Esperaba %s, pero salió %s",
				esperado, obtenido.ID)
		}
	}

	if d.Len() != 0 {
		t.Error("El deque debería terminar vacío")
	}
}

func TestDequeMezcla(t *testing.T) {
	d := NuevoDeque()

	reporteA := tipos.Reporte{ID: "A"}
	reporteB := tipos.Reporte{ID: "B"}
	reporteC := tipos.Reporte{ID: "C"}
	reporteD := tipos.Reporte{ID: "D"}

	d.InsertarFinal(reporteA)
	d.InsertarFrente(reporteB)
	d.InsertarFinal(reporteC)
	d.InsertarFrente(reporteD)

	esperados := []string{"D", "B", "A", "C"}

	for _, esperado := range esperados {
		obtenido, ok := d.QuitarFrente()

		if !ok {
			t.Fatal("No se pudo quitar el reporte")
		}

		if obtenido.ID != esperado {
			t.Errorf("Esperaba %s, pero salió %s",
				esperado, obtenido.ID)
		}
	}

	if d.Len() != 0 {
		t.Error("El deque debería terminar vacío")
	}
}

func TestDequeFrenteNoQuita(t *testing.T) {
	d := NuevoDeque()

	reporteA := tipos.Reporte{ID: "A"}
	reporteB := tipos.Reporte{ID: "B"}

	d.InsertarFinal(reporteA)
	d.InsertarFinal(reporteB)

	if d.Len() != 2 {
		t.Errorf("Esperaba 2 elementos, pero hay %d", d.Len())
	}

	obtenido, ok := d.Frente()

	if !ok {
		t.Fatal("No se pudo consultar el frente")
	}

	if obtenido.ID != "A" {
		t.Errorf("Esperaba A, pero salió %s", obtenido.ID)
	}

	if d.Len() != 2 {
		t.Error("Frente no debería eliminar elementos")
	}

	primero, ok := d.QuitarFrente()

	if !ok || primero.ID != "A" {
		t.Error("El primer reporte debería seguir siendo A")
	}
}
