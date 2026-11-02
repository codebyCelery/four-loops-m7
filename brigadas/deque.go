package brigadas

import "alerta/tipos"

const capacidadInicial = 4

type Deque struct {
	datos  []tipos.Reporte
	frente int
	cuenta int
}

func NuevoDeque() *Deque {
	return &Deque{datos: make([]tipos.Reporte, capacidadInicial)}
}

func (d *Deque) Len() int { return d.cuenta }

// crecer duplica el arreglo y copia los elementos EN ORDEN desde el frente.
func (d *Deque) crecer() {
	nuevo := make([]tipos.Reporte, len(d.datos)*2)
	for i := 0; i < d.cuenta; i++ {
		nuevo[i] = d.datos[(d.frente+i)%len(d.datos)]
	}
	d.datos = nuevo
	d.frente = 0
}

func (d *Deque) InsertarFinal(r tipos.Reporte) {
	if d.cuenta == len(d.datos) {
		d.crecer()
	}
	pos := (d.frente + d.cuenta) % len(d.datos)
	d.datos[pos] = r
	d.cuenta++
}

func (d *Deque) InsertarFrente(r tipos.Reporte) {
	if d.cuenta == len(d.datos) {
		d.crecer()
	}
	d.frente = (d.frente - 1 + len(d.datos)) % len(d.datos)
	d.datos[d.frente] = r
	d.cuenta++
}

func (d *Deque) QuitarFrente() (tipos.Reporte, bool) {
	if d.cuenta == 0 {
		return tipos.Reporte{}, false
	}
	r := d.datos[d.frente]
	d.datos[d.frente] = tipos.Reporte{}
	d.frente = (d.frente + 1) % len(d.datos)
	d.cuenta--
	return r, true
}

func (d *Deque) Frente() (tipos.Reporte, bool) {
	if d.cuenta == 0 {
		return tipos.Reporte{}, false
	}
	return d.datos[d.frente], true
}
