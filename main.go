package main

import (
    "github.com/saucelibertarix/servidorgofiber/app"
)

func main() {
    a := app.App{}
    a.Initialize(":3000")
}
