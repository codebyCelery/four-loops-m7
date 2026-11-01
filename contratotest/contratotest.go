// Package contratotest: la misma batería de pruebas para todos los equipos.
// Si su módulo la pasa, encaja en la integración. En cada revisión el
// profesor corre además pruebas que ustedes no conocen.
package contratotest

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"alerta/contratos"
	"alerta/tipos"
)

var t0 = time.Date(2027, 2, 10, 0, 0, 0, 0, time.UTC)

func h(min int) time.Time { return t0.Add(time.Duration(min) * time.Minute) }

func sinPanico(t *testing.T) {
	t.Helper()
	if r := recover(); r != nil {
		t.Fatalf("el módulo entró en pánico: %v", r)
	}
}

func igual(t *testing.T, que string, got, want any) {
	t.Helper()
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("%s: obtuvo %v, se esperaba %v", que, got, want)
	}
}

// rutasFijas: grafo pequeño conocido, para probar M5, M6 y M7.
//
//	A - B - C - D      E (aislada)
type rutasFijas struct{}

var orden = map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}

func (rutasFijas) AgregarTramo(a, b string) {}
func (rutasFijas) CortarTramo(a, b string)  {}
func (rutasFijas) Saltos(a, b string) int {
	x, ok1 := orden[a]
	y, ok2 := orden[b]
	if !ok1 || !ok2 {
		if a == b {
			return 0
		}
		return -1
	}
	if x > y {
		return x - y
	}
	return y - x
}
func (r rutasFijas) Camino(a, b string) []string {
	if r.Saltos(a, b) < 0 {
		return nil
	}
	nombres := []string{"A", "B", "C", "D"}
	x, y := orden[a], orden[b]
	var c []string
	for i := x; ; {
		c = append(c, nombres[i])
		if i == y {
			return c
		}
		if x < y {
			i++
		} else {
			i--
		}
	}
}
func (rutasFijas) Aislados() [][]string { return [][]string{{"A", "B", "C", "D"}, {"E"}} }

func Reportes(t *testing.T, nuevo func() contratos.Reportes) {
	defer sinPanico(t)
	s := nuevo()
	igual(t, "primer reporte", s.Registrar(tipos.Reporte{ID: "r1", ZonaID: "A", Tipo: "HUAICO", Hora: h(0)}), true)
	igual(t, "mismo evento a los 10 min es duplicado", s.Registrar(tipos.Reporte{ID: "r2", ZonaID: "A", Tipo: "HUAICO", Hora: h(10)}), false)
	igual(t, "otro tipo en la misma zona no es duplicado", s.Registrar(tipos.Reporte{ID: "r3", ZonaID: "A", Tipo: "INUNDACION", Hora: h(5)}), true)
	igual(t, "a los 40 min ya no es duplicado", s.Registrar(tipos.Reporte{ID: "r4", ZonaID: "A", Tipo: "HUAICO", Hora: h(40)}), true)
	_, ok := s.Buscar("r2")
	igual(t, "un duplicado no se guarda", ok, false)
	r, _ := s.Siguiente()
	igual(t, "la cola respeta el orden de llegada", r.ID, "r1")
	var ids []string
	for _, x := range s.EntreHoras(h(0), h(40)) {
		ids = append(ids, x.ID)
	}
	igual(t, "EntreHoras ordena por hora e incluye extremos", ids, []string{"r1", "r3", "r4"})
}

func Prioridad(t *testing.T, nuevo func() contratos.Prioridad) {
	defer sinPanico(t)
	p := nuevo()
	p.Actualizar("A", 5)
	p.Actualizar("B", 9)
	p.Actualizar("C", 7)
	p.Actualizar("B", 1) // bajó el riesgo de B
	var got []string
	for {
		z, _, ok := p.Siguiente()
		if !ok {
			break
		}
		got = append(got, z)
	}
	igual(t, "orden por riesgo con actualización", got, []string{"C", "A", "B"})
}

func Sensores(t *testing.T, nuevo func() contratos.Sensores) {
	defer sinPanico(t)
	s := nuevo()
	s.Registrar(tipos.Lectura{EstacionID: "e1", ZonaID: "A", Hora: h(0), Nivel: 3.0})
	s.Registrar(tipos.Lectura{EstacionID: "e1", ZonaID: "A", Hora: h(30), Nivel: 1.0})
	s.Registrar(tipos.Lectura{EstacionID: "e2", ZonaID: "B", Hora: h(10), Nivel: 2.0})
	s.Registrar(tipos.Lectura{EstacionID: "e3", ZonaID: "C", Hora: h(20), Nivel: 2.5})
	igual(t, "EnAlerta usa el nivel más reciente", s.EnAlerta(2), []string{"C", "B"})
	igual(t, "Lecturas por rango", len(s.Lecturas("e1", h(0), h(30))), 2)
}

func Rutas(t *testing.T, nuevo func() contratos.Rutas) {
	defer sinPanico(t)
	g := nuevo()
	for _, e := range [][2]string{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"A", "D"}, {"E", "F"}} {
		g.AgregarTramo(e[0], e[1])
	}
	igual(t, "saltos A-C", g.Saltos("A", "C"), 2)
	igual(t, "sin camino", g.Saltos("A", "E"), -1)
	g.CortarTramo("A", "D")
	igual(t, "saltos A-D tras cortar", g.Saltos("A", "D"), 3)
	igual(t, "camino A-D", g.Camino("A", "D"), []string{"A", "B", "C", "D"})
	g.CortarTramo("B", "C")
	igual(t, "componentes", len(g.Aislados()), 3)
}

func Albergues(t *testing.T, nuevo func(contratos.Rutas) contratos.Albergues) {
	defer sinPanico(t)
	s := nuevo(rutasFijas{})
	s.AgregarAlbergue(tipos.Albergue{ID: "lejos", ZonaID: "D", Capacidad: 10})
	s.AgregarAlbergue(tipos.Albergue{ID: "cerca", ZonaID: "B", Capacidad: 4})
	s.Solicitar(tipos.Familia{ID: "f1", ZonaID: "A", Integrantes: 3, Vulnerables: 0})
	s.Solicitar(tipos.Familia{ID: "f2", ZonaID: "A", Integrantes: 4, Vulnerables: 2})
	f, a, _ := s.AsignarSiguiente()
	igual(t, "primero la familia más vulnerable al albergue más cercano", f+"->"+a, "f2->cerca")
	f, a, _ = s.AsignarSiguiente()
	igual(t, "si no cabe, el siguiente más cercano", f+"->"+a, "f1->lejos")
	d, _ := s.DondeEsta("f2")
	igual(t, "DondeEsta", d, "cerca")
	_, _, err := s.AsignarSiguiente()
	igual(t, "sin espera", errors.Is(err, contratos.ErrVacio), true)
	s.Solicitar(tipos.Familia{ID: "f3", ZonaID: "E", Integrantes: 2, Vulnerables: 3})
	s.Solicitar(tipos.Familia{ID: "f4", ZonaID: "A", Integrantes: 9, Vulnerables: 0})
	f, _, err = s.AsignarSiguiente()
	igual(t, "familia sin ruta sale de la espera", f+fmt.Sprint(errors.Is(err, contratos.ErrSinRuta)), "f3true")
	f, _, err = s.AsignarSiguiente()
	igual(t, "sin cupo, la familia espera", f+fmt.Sprint(errors.Is(err, contratos.ErrSinCapacidad)), "f4true")
}

type almacenInfinito struct{}

func (almacenInfinito) Ingresar(tipos.Lote) {}
func (almacenInfinito) Retirar(p string, c int) ([]tipos.Lote, error) {
	return []tipos.Lote{{Producto: p, Cantidad: c}}, nil
}
func (almacenInfinito) Stock(string) int        { return 1 << 30 }
func (almacenInfinito) BajoMinimo(int) []string { return nil }

func Ayuda(t *testing.T, nuevo func(contratos.Rutas, contratos.Almacenes) contratos.Ayuda) {
	defer sinPanico(t)
	s := nuevo(rutasFijas{}, almacenInfinito{})
	s.AgregarCamion("c1", "A")
	s.AgregarCamion("c2", "A")
	s.RegistrarNecesidad("C", 100)
	s.RegistrarNecesidad("D", 300)
	c, ruta, _ := s.Despachar()
	igual(t, "primero la zona con más necesidad", c+fmt.Sprint(ruta), "c1[A B C D]")
	c, _, _ = s.Despachar()
	igual(t, "los camiones salen en orden de cola", c, "c2")
	_, _, err := s.Despachar()
	igual(t, "sin necesidad pendiente", errors.Is(err, contratos.ErrVacio), true)
}

func Brigadas(t *testing.T, nuevo func(contratos.Rutas) contratos.Brigadas) {
	defer sinPanico(t)
	s := nuevo(rutasFijas{})
	s.AgregarBrigada("b1", "A")
	s.AgregarBrigada("b2", "D")
	s.Emergencia(tipos.Reporte{ID: "r1", ZonaID: "C"})
	s.Emergencia(tipos.Reporte{ID: "r2", ZonaID: "B"})
	b, r, _ := s.AsignarSiguiente()
	igual(t, "la emergencia más antigua a la brigada más cercana", b+"->"+r, "b2->r1")
	igual(t, "deshacer", s.Deshacer(), nil)
	b, r, _ = s.AsignarSiguiente()
	igual(t, "tras deshacer, la emergencia vuelve al frente", b+"->"+r, "b2->r1")
	b, r, _ = s.AsignarSiguiente()
	igual(t, "la siguiente toma la brigada que queda", b+"->"+r, "b1->r2")
	s.Emergencia(tipos.Reporte{ID: "r3", ZonaID: "E"})
	s.Emergencia(tipos.Reporte{ID: "r4", ZonaID: "A"})
	_, r, err := s.AsignarSiguiente()
	igual(t, "zona sin ruta sale de la cola", r+fmt.Sprint(errors.Is(err, contratos.ErrSinRuta)), "r3true")
	_, r, err = s.AsignarSiguiente()
	igual(t, "sin brigadas libres, la emergencia espera", r+fmt.Sprint(errors.Is(err, contratos.ErrSinRecursos)), "r4true")
	s.Liberar("b1")
	b, r, _ = s.AsignarSiguiente()
	igual(t, "tras liberar", b+"->"+r, "b1->r4")
}

func Almacenes(t *testing.T, nuevo func() contratos.Almacenes) {
	defer sinPanico(t)
	s := nuevo()
	s.Ingresar(tipos.Lote{ID: "L1", Producto: "agua", Cantidad: 100, Vence: h(3000)})
	s.Ingresar(tipos.Lote{ID: "L2", Producto: "agua", Cantidad: 50, Vence: h(1000)})
	s.Ingresar(tipos.Lote{ID: "L3", Producto: "arroz", Cantidad: 20, Vence: h(5000)})
	ls, _ := s.Retirar("agua", 70)
	igual(t, "vence primero, sale primero", ls[0].ID+fmt.Sprint(ls[0].Cantidad)+ls[1].ID+fmt.Sprint(ls[1].Cantidad), "L250L120")
	igual(t, "stock", s.Stock("agua"), 80)
	_, err := s.Retirar("agua", 500)
	igual(t, "sin stock", errors.Is(err, contratos.ErrSinStock), true)
	igual(t, "un retiro fallido no cambia el stock", s.Stock("agua"), 80)
	igual(t, "bajo mínimo", s.BajoMinimo(50), []string{"arroz"})
}

func Lluvias(t *testing.T, nuevo func() contratos.Lluvias) {
	defer sinPanico(t)
	s := nuevo()
	for i, mm := range []float64{10, 20, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 30} { // cada 30 min
		s.Registrar("eA", "A", h(30*i), mm)
	}
	s.Registrar("eB", "B", h(0), 40)
	s.Registrar("eC", "C", h(0), 40)
	igual(t, "la ventana de 6 h descarta la primera lectura", s.Acumulado("eA"), 55.0)
	igual(t, "estación sin lecturas", s.Acumulado("eX"), 0.0)
	igual(t, "EnRiesgo con empate por menor ID", s.EnRiesgo(2), []string{"A", "B"})
}

func Salud(t *testing.T, nuevo func(contratos.Rutas) contratos.Salud) {
	defer sinPanico(t)
	s := nuevo(rutasFijas{})
	s.AgregarPuesto("p-lejos", "D", 5)
	s.AgregarPuesto("p-cerca", "B", 1)
	s.Ingresar(tipos.Herido{ID: "h1", ZonaID: "A", Gravedad: 2})
	s.Ingresar(tipos.Herido{ID: "h2", ZonaID: "A", Gravedad: 5})
	s.Ingresar(tipos.Herido{ID: "h3", ZonaID: "A", Gravedad: 5})
	hid, p, _ := s.DerivarSiguiente()
	igual(t, "el más grave al puesto más cercano", hid+"->"+p, "h2->p-cerca")
	hid, p, _ = s.DerivarSiguiente()
	igual(t, "empate de gravedad: el que ingresó antes; puesto lleno, el siguiente", hid+"->"+p, "h3->p-lejos")
	a, _ := s.Alta("p-cerca")
	igual(t, "alta del que lleva más tiempo", a, "h2")
	hid, p, _ = s.DerivarSiguiente()
	igual(t, "la cama liberada se vuelve a usar", hid+"->"+p, "h1->p-cerca")
	s.Ingresar(tipos.Herido{ID: "h4", ZonaID: "E", Gravedad: 4})
	hid, _, err := s.DerivarSiguiente()
	igual(t, "herido sin ruta sale de la espera", hid+fmt.Sprint(errors.Is(err, contratos.ErrSinRuta)), "h4true")
	_, err = s.Alta("p-vacio")
	igual(t, "alta sin pacientes", errors.Is(err, contratos.ErrVacio), true)
}

// rutasCortadas: A-B-C conectados; D y E sueltos.
type rutasCortadas struct{ rutasFijas }

func (rutasCortadas) Saltos(a, b string) int {
	o := map[string]int{"A": 0, "B": 1, "C": 2}
	x, ok1 := o[a]
	y, ok2 := o[b]
	if a == b {
		return 0
	}
	if !ok1 || !ok2 {
		return -1
	}
	if x > y {
		return x - y
	}
	return y - x
}
func (rutasCortadas) Aislados() [][]string { return [][]string{{"A", "B", "C"}, {"D"}, {"E"}} }

func Maquinaria(t *testing.T, nuevo func(contratos.Rutas) contratos.Maquinaria) {
	defer sinPanico(t)
	s := nuevo(rutasCortadas{})
	for z, p := range map[string]int{"A": 100, "B": 50, "C": 20, "D": 30, "E": 80} {
		s.AgregarZona(z, p)
	}
	s.AgregarMaquina("m2", "A")
	s.AgregarMaquina("m1", "B")
	s.TramoCortado("A", "B") // redundante: ya están conectados
	s.TramoCortado("C", "D") // reconecta 30
	s.TramoCortado("C", "E") // reconecta 80
	m, a, b, _ := s.Reabrir()
	igual(t, "primero el que reconecta más población, con la máquina más cercana", m+" "+a+"-"+b, "m1 C-E")
	m, a, b, _ = s.Reabrir()
	igual(t, "luego el siguiente", m+" "+a+"-"+b, "m2 C-D")
	_, _, _, err := s.Reabrir()
	igual(t, "sin máquinas libres", errors.Is(err, contratos.ErrSinRecursos), true)
	s.Terminar("m1")
	m, a, b, _ = s.Reabrir()
	igual(t, "el redundante al final", m+" "+a+"-"+b, "m1 A-B")
	_, _, _, err = s.Reabrir()
	igual(t, "sin tramos pendientes", errors.Is(err, contratos.ErrVacio), true)
}

func Reencuentro(t *testing.T, nuevo func() contratos.Reencuentro) {
	defer sinPanico(t)
	s := nuevo()
	s.Registrar(tipos.Persona{ID: "p3", Nombre: "Rosa", Apellido: "Chunga", FamiliaID: "f1", AlbergueID: "a1"})
	s.Registrar(tipos.Persona{ID: "p1", Nombre: "Luis", Apellido: "Chiroque", FamiliaID: "f2", AlbergueID: "a2"})
	s.Registrar(tipos.Persona{ID: "p2", Nombre: "Ana", Apellido: "Chunga", FamiliaID: "f1", AlbergueID: "a1"})
	s.Registrar(tipos.Persona{ID: "p4", Nombre: "Ana", Apellido: "Chunga", FamiliaID: "f3", AlbergueID: "a3"})
	var ids []string
	for _, p := range s.PorApellido("Ch") {
		ids = append(ids, p.ID)
	}
	igual(t, "prefijo ordenado por apellido, nombre e ID", ids, []string{"p1", "p2", "p4", "p3"})
	igual(t, "prefijo sin resultados", len(s.PorApellido("Zz")), 0)
	ids = nil
	for _, p := range s.Familia("f1") {
		ids = append(ids, p.ID)
	}
	igual(t, "familia en orden de registro", ids, []string{"p3", "p2"})
	p, _ := s.Buscar("p4")
	igual(t, "buscar por documento", p.AlbergueID, "a3")
}

func Alertas(t *testing.T, nuevo func() contratos.Alertas) {
	defer sinPanico(t)
	s := nuevo()
	igual(t, "emitir", s.Emitir("A", "evacuar", 3), true)
	igual(t, "no se repite el mismo aviso", s.Emitir("A", "evacuar", 3), false)
	s.Emitir("B", "lluvia", 2)
	s.Emitir("C", "aislada", 4)
	s.Emitir("D", "evacuar", 3)
	var got []string
	for {
		z, _, err := s.Enviar()
		if err != nil {
			break
		}
		got = append(got, z)
	}
	igual(t, "prioridad y orden de emisión", got, []string{"C", "A", "D", "B"})
	s.Fallo("A", "evacuar")
	s.Fallo("B", "lluvia")
	s.Fallo("C", "aislada")
	got = nil
	for i := 0; i < 4; i++ {
		z, _, _ := s.Reintentar()
		got = append(got, z)
	}
	igual(t, "la ronda vuelve al inicio", got, []string{"A", "B", "C", "A"})
	s.Confirmar("B", "lluvia")
	z, _, _ := s.Reintentar()
	igual(t, "confirmado sale de la ronda", z, "C")
	s.Confirmar("A", "evacuar")
	s.Confirmar("C", "aislada")
	_, _, err := s.Reintentar()
	igual(t, "ronda vacía", errors.Is(err, contratos.ErrVacio), true)
}
