package main

/*
Proteger rutas de crear, actualizar y eliminar para solo clientes autenticados con role admin
Enlazar director y peliculas 1-n
Crear modelo oscar con peliculas 1-n
Añadir a peliculas la posibilidad de que sean puntuadas (Solo pueden puntuar usuarios registrado)
*/

import (
    "github.com/saucelibertarix/servidorgofiber/app"
)

func main() {
    a := app.App{}
    a.Initialize(":3000")
}
