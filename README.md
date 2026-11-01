# Alerta Niño — AED 16951

Cada equipo implementa su módulo en su carpeta (reportes/, prioridad/, sensores/, rutas/,
albergues/, ayuda/, brigadas/, almacenes/, lluvias/, salud/, maquinaria/, reencuentro/, alertas/). No modifiquen tipos/, contratos/, stubs/ ni contratotest/.

    go test ./rutas/            # su módulo contra el contrato (cambie la carpeta)
    go run ./cmd/generar        # genera el escenario en datos/ para medir
    go run ./cmd/generar -escala=50   # escenario grande para el benchmark

Mientras M4 y M8 no estén listos, M5, M6, M7, M10 y M11 se prueban con stubs.Rutas{} y stubs.Almacenes{}.
Dentro de su módulo no se permite usar map, container/* ni sort.
