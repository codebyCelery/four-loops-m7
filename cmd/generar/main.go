// Genera el escenario simulado "Niño Costero" en la carpeta datos/.
// Las zonas, poblaciones y tramos son una simplificación de la región Piura
// para fines del curso; no son datos oficiales.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var t0 = time.Date(2027, 2, 10, 0, 0, 0, 0, time.UTC)

type zona struct {
	id, nombre string
	pob        int
	pico, alto float64 // hora del pico de lluvia y nivel máximo (m)
}

var zonas = []zona{
	{"PIU", "Piura", 480000, 18, 2.2}, {"CAS", "Castilla", 160000, 20, 2.0},
	{"VEI", "Veintiséis de Octubre", 170000, 0, 0}, {"CAT", "Catacaos", 75000, 22, 3.4},
	{"CUR", "Cura Mori", 20000, 22, 3.8}, {"ARE", "La Arena", 40000, 26, 3.1},
	{"UNI", "La Unión", 40000, 28, 2.6}, {"SEC", "Sechura", 45000, 0, 0},
	{"PAI", "Paita", 100000, 0, 0}, {"SUL", "Sullana", 180000, 14, 2.4},
	{"BEL", "Bellavista", 40000, 0, 0}, {"MAR", "Marcavelica", 30000, 15, 2.8},
	{"TAM", "Tambogrande", 120000, 12, 2.5}, {"CHU", "Chulucanas", 80000, 30, 3.0},
	{"MOR", "Morropón", 15000, 31, 3.6}, {"TAL", "Talara", 90000, 0, 0},
	{"MAN", "Máncora", 13000, 0, 0}, {"LAS", "Las Lomas", 27000, 10, 2.2},
}

var tramos = [][2]string{
	{"PIU", "CAS"}, {"PIU", "VEI"}, {"VEI", "CAS"}, {"PIU", "CAT"}, {"CAT", "CUR"},
	{"CAT", "ARE"}, {"CUR", "ARE"}, {"ARE", "UNI"}, {"UNI", "SEC"}, {"PIU", "PAI"},
	{"PAI", "SUL"}, {"PIU", "SUL"}, {"SUL", "BEL"}, {"SUL", "MAR"}, {"SUL", "TAL"},
	{"TAL", "MAN"}, {"CAS", "TAM"}, {"TAM", "LAS"}, {"CAS", "CHU"}, {"CHU", "MOR"},
}

func main() {
	dir := flag.String("datos", "datos", "carpeta de salida")
	escala := flag.Int("escala", 1, "multiplica el volumen de eventos (benchmark)")
	semilla := flag.Int64("semilla", 2017, "semilla aleatoria")
	flag.Parse()
	rnd := rand.New(rand.NewSource(*semilla))
	rnd2 := rand.New(rand.NewSource(*semilla + 1)) // M9–M13: no altera los eventos de M1–M8
	os.MkdirAll(*dir, 0o755)
	esc := func(nombre string, filas [][]string) {
		f, err := os.Create(filepath.Join(*dir, nombre))
		if err != nil {
			panic(err)
		}
		w := csv.NewWriter(f)
		w.WriteAll(filas)
		f.Close()
	}
	hora := func(min float64) string {
		return t0.Add(time.Duration(min * float64(time.Minute))).Format(time.RFC3339)
	}

	fz := [][]string{{"id", "nombre", "poblacion"}}
	for _, z := range zonas {
		fz = append(fz, []string{z.id, z.nombre, fmt.Sprint(z.pob)})
	}
	esc("zonas.csv", fz)
	ft := [][]string{{"a", "b"}}
	for _, t := range tramos {
		ft = append(ft, []string{t[0], t[1]})
	}
	esc("tramos.csv", ft)
	esc("albergues.csv", [][]string{{"id", "zona", "capacidad"},
		{"ALB-PIU1", "PIU", "900"}, {"ALB-PIU2", "PIU", "600"}, {"ALB-CAS", "CAS", "500"},
		{"ALB-SUL", "SUL", "700"}, {"ALB-PAI", "PAI", "600"}, {"ALB-TAL", "TAL", "400"},
		{"ALB-CHU", "CHU", "300"}, {"ALB-SEC", "SEC", "400"}, {"ALB-TAM", "TAM", "300"}})
	esc("recursos.csv", [][]string{{"tipo", "id", "base"},
		{"camion", "CAM-01", "PIU"}, {"camion", "CAM-02", "PIU"}, {"camion", "CAM-03", "SUL"},
		{"camion", "CAM-04", "PAI"}, {"camion", "CAM-05", "CAS"}, {"camion", "CAM-06", "SUL"},
		{"brigada", "BRI-01", "PIU"}, {"brigada", "BRI-02", "CAS"}, {"brigada", "BRI-03", "SUL"},
		{"brigada", "BRI-04", "CAT"}, {"brigada", "BRI-05", "CHU"}, {"brigada", "BRI-06", "TAL"},
		{"maquina", "MAQ-01", "PIU"}, {"maquina", "MAQ-02", "SUL"}, {"maquina", "MAQ-03", "CAS"}})
	esc("puestos.csv", [][]string{{"id", "zona", "camas"},
		{"PST-PIU", "PIU", "40"}, {"PST-CAS", "CAS", "15"}, {"PST-SUL", "SUL", "20"}, {"PST-CAT", "CAT", "8"},
		{"PST-CHU", "CHU", "8"}, {"PST-PAI", "PAI", "10"}, {"PST-TAL", "TAL", "6"}})

	type ev struct {
		min  float64
		fila []string
	}
	var evs []ev
	add := func(min float64, campos ...string) {
		evs = append(evs, ev{min, append([]string{hora(min)}, campos...)})
	}
	// Almacenes: stock inicial y reabastecimiento a las 24 h.
	for i, l := range []struct {
		p    string
		c, d int
	}{{"agua", 9000, 20}, {"agua", 6000, 5}, {"agua", 5000, 60}, {"arroz", 3000, 180}, {"frazadas", 800, 900}, {"medicinas", 400, 30}} {
		add(0, "lote", fmt.Sprintf("L%02d", i+1), l.p, fmt.Sprint(l.c), fmt.Sprint(l.d))
	}
	add(24*60, "lote", "L07", "agua", "8000", "15")
	// Cortes de vía.
	add(21*60, "corte", "CAT", "CUR")
	add(23*60, "corte", "CUR", "ARE")
	add(31*60, "corte", "CHU", "MOR")
	add(36*60, "corte", "CAS", "CHU")

	apellidos := []string{"Chunga", "Sernaqué", "Yarlequé", "Ipanaqué", "Zapata", "Ruiz", "Castillo", "Silva",
		"Ramos", "Purizaca", "Querevalú", "Chiroque", "Facundo", "Panta", "Vílchez", "Saavedra", "Juárez", "Morales", "Cruz", "Távara"}
	nombres := []string{"Rosa", "José", "María", "Luis", "Ana", "Carlos", "Juana", "Pedro", "Lucía", "Jorge", "Carmen", "Miguel"}
	nRep, nFam, nHer, nPer := 0, 0, 0, 0
	tipos := []string{"INUNDACION", "HUAICO", "DESBORDE"}
	for _, z := range zonas {
		if z.alto == 0 {
			continue
		}
		for m := 0.0; m <= 48*60; m += 30 / float64(*escala) {
			h := m / 60
			nivel := z.alto*math.Exp(-math.Pow(h-z.pico, 2)/18) + rnd.Float64()*0.2
			add(m, "lectura", "E-"+z.id, z.id, fmt.Sprintf("%.2f", nivel))
			// La lluvia llega unas 3 h antes que la crecida del río.
			mm := 6*z.alto*math.Exp(-math.Pow(h-(z.pico-3), 2)/8) + rnd2.Float64()*1.5
			add(m, "lluvia", "P-"+z.id, z.id, fmt.Sprintf("%.1f", mm))
			if nivel > 1.2 && rnd.Float64() < nivel/4 {
				t := tipos[rnd.Intn(3)]
				for k := 0; k < 1+rnd.Intn(4); k++ { // ráfaga: varios vecinos reportan lo mismo
					nRep++
					add(m+float64(k)*3+rnd.Float64()*2, "reporte", fmt.Sprintf("R%05d", nRep), z.id, t)
				}
				if t != "INUNDACION" && rnd2.Float64() < 0.6 {
					for k := 0; k < 1+rnd2.Intn(3); k++ {
						nHer++
						add(m+rnd2.Float64()*20, "herido", fmt.Sprintf("H%05d", nHer), z.id, fmt.Sprint(1+rnd2.Intn(5)))
					}
				}
			}
			if nivel > 1.8 {
				for k := 0; k < 1+int(nivel)*z.pob/60000; k++ {
					nFam++
					mf, integ := m+rnd.Float64()*30, 2+rnd.Intn(5)
					add(mf, "familia", fmt.Sprintf("F%05d", nFam), z.id, fmt.Sprint(integ), fmt.Sprint(rnd.Intn(3)))
					ap := apellidos[rnd2.Intn(len(apellidos))]
					for k := 0; k < integ; k++ {
						nPer++
						add(mf, "persona", fmt.Sprintf("P%06d", nPer), fmt.Sprintf("F%05d", nFam), ap, nombres[rnd2.Intn(len(nombres))])
					}
				}
			}
		}
	}
	// Familiares que buscan a los suyos en los albergues (M12).
	for k := 1; k <= 30; k++ {
		ap := []rune(apellidos[rnd2.Intn(len(apellidos))])
		add(6*60+rnd2.Float64()*42*60, "busqueda", fmt.Sprintf("Q%03d", k), string(ap[:min(len(ap), 3+rnd2.Intn(3))]))
	}
	sort.SliceStable(evs, func(i, j int) bool { return evs[i].min < evs[j].min })
	fe := [][]string{{"hora", "tipo", "c1", "c2", "c3", "c4"}}
	for _, e := range evs {
		f := append(e.fila, "", "", "", "")[:6]
		fe = append(fe, f)
	}
	esc("eventos.csv", fe)
	fmt.Printf("escenario generado en %s: %d eventos (%d reportes, %d familias, %d personas, %d heridos)\n", *dir, len(evs), nRep, nFam, nPer, nHer)
}
