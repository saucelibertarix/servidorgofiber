package main

/*
Crear modelo director
Crear servicio director
    Create
    Get
    Update
    Delete
Crear controller director
Crear rutas director
Crear middleware director (delete solo admin)
 */

import (
    "github.com/saucelibertarix/servidorgofiber/app"
)

func main() {
    a := app.App{}
    a.Initialize(":3000")
}
