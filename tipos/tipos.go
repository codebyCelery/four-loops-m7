// Package tipos: datos que comparten todos los módulos. Lo mantiene el profesor.
package tipos

import "time"

type Zona struct {
	ID        string
	Nombre    string
	Poblacion int
}

type Reporte struct {
	ID, ZonaID, Tipo string // Tipo: HUAICO, INUNDACION, DESBORDE
	Hora             time.Time
}

type Lectura struct {
	EstacionID, ZonaID string
	Hora               time.Time
	Nivel              float64 // metros sobre el nivel normal
}

type Familia struct {
	ID, ZonaID               string
	Integrantes, Vulnerables int
}

type Albergue struct {
	ID, ZonaID string
	Capacidad  int // personas
}

type Lote struct {
	ID, Producto string
	Cantidad     int
	Vence        time.Time
}

type Herido struct {
	ID, ZonaID string
	Gravedad   int // 1 leve … 5 crítico
	Hora       time.Time
}

type Persona struct {
	ID, Nombre, Apellido  string
	FamiliaID, AlbergueID string
}
