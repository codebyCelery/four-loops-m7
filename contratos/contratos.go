// Package contratos: interfaces de los ocho módulos. Lo mantiene el profesor;
// ningún equipo lo modifica.
package contratos

import (
	"errors"
	"time"

	"alerta/tipos"
)

var (
	ErrVacio        = errors.New("no hay nada pendiente")
	ErrSinRecursos  = errors.New("no hay recursos libres")
	ErrSinCapacidad = errors.New("no hay albergue con capacidad alcanzable")
	ErrSinRuta      = errors.New("no hay ruta a la zona")
	ErrSinStock     = errors.New("stock insuficiente")
)

// M1 — cola, tabla hash, BST
type Reportes interface {
	Registrar(r tipos.Reporte) bool // false si es duplicado
	Siguiente() (tipos.Reporte, bool)
	Buscar(id string) (tipos.Reporte, bool)
	EntreHoras(desde, hasta time.Time) []tipos.Reporte // ordenados por hora, extremos incluidos
}

// M2 — heap, tabla hash
type Prioridad interface {
	Actualizar(zonaID string, riesgo int)            // inserta o cambia el riesgo
	Siguiente() (zonaID string, riesgo int, ok bool) // mayor riesgo; empate: menor ID
}

// M3 — BST, heap
type Sensores interface {
	Registrar(l tipos.Lectura)
	Lecturas(estacionID string, desde, hasta time.Time) []tipos.Lectura
	EnAlerta(k int) []string // k zonas con mayor nivel actual; empate: menor ID
}

// M4 — grafo, BFS (cola), DFS (pila)
type Rutas interface {
	AgregarTramo(a, b string)
	CortarTramo(a, b string)
	Saltos(a, b string) int // -1 si no hay camino
	Camino(a, b string) []string
	Aislados() [][]string // componentes conexas, cada una ordenada
}

// M5 — heap, tabla hash
type Albergues interface {
	AgregarAlbergue(a tipos.Albergue)
	Solicitar(f tipos.Familia)
	AsignarSiguiente() (familiaID, albergueID string, err error) // ErrSinRuta si no hay albergue alcanzable
	DondeEsta(familiaID string) (string, bool)
}

// M6 — heap, cola, tabla hash
type Ayuda interface {
	AgregarCamion(id, baseZonaID string)
	RegistrarNecesidad(zonaID string, cantidad int)
	Despachar() (camionID string, ruta []string, err error)
}

// M7 — cola, pila, tabla hash
type Brigadas interface {
	AgregarBrigada(id, baseZonaID string)
	Emergencia(r tipos.Reporte)
	AsignarSiguiente() (brigadaID, reporteID string, err error) // ErrSinRuta si nadie puede llegar
	Deshacer() error                                            // revierte la última asignación
	Liberar(brigadaID string)
}

// M8 — heap, tabla hash, BST
type Almacenes interface {
	Ingresar(l tipos.Lote)
	Retirar(producto string, cantidad int) ([]tipos.Lote, error) // vence primero, sale primero
	Stock(producto string) int
	BajoMinimo(minimo int) []string // de menor a mayor stock
}

// M9 — deque (ventana de 6 h), heap, tabla hash
type Lluvias interface {
	Registrar(estacionID, zonaID string, hora time.Time, mm float64)
	Acumulado(estacionID string) float64 // mm de las últimas 6 h, contadas desde su lectura más reciente (sume en décimas enteras)
	EnRiesgo(k int) []string             // k zonas con mayor acumulado; empate: menor ID
}

// M10 — heap, cola, tabla hash
type Salud interface {
	AgregarPuesto(id, zonaID string, camas int)
	Ingresar(h tipos.Herido)
	DerivarSiguiente() (heridoID, puestoID string, err error) // ErrSinRuta: sale; ErrSinCapacidad: espera
	Alta(puestoID string) (heridoID string, err error)        // el que lleva más tiempo en el puesto
}

// M11 — heap, tabla hash (consulta el grafo de M4)
type Maquinaria interface {
	AgregarZona(zonaID string, poblacion int)
	AgregarMaquina(id, baseZonaID string)
	TramoCortado(a, b string)
	Reabrir() (maquinaID, a, b string, err error) // el tramo que reconecta más población; ErrSinRecursos si no hay máquina que llegue
	Terminar(maquinaID string)
}

// M12 — tabla hash, BST, lista enlazada
type Reencuentro interface {
	Registrar(p tipos.Persona)
	Buscar(id string) (tipos.Persona, bool)
	Familia(familiaID string) []tipos.Persona   // en orden de registro
	PorApellido(prefijo string) []tipos.Persona // por apellido, nombre e ID
}

// M13 — heap, tabla hash, lista circular
type Alertas interface {
	Emitir(zonaID, mensaje string, prioridad int) bool // false si ya se emitió
	Enviar() (zonaID, mensaje string, err error)       // mayor prioridad; empate: orden de emisión
	Fallo(zonaID, mensaje string)                      // pasa a la ronda de reintentos
	Reintentar() (zonaID, mensaje string, err error)   // el siguiente de la ronda, en círculo
	Confirmar(zonaID, mensaje string)                  // sale de la ronda
}
