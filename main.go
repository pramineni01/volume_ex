// Problem definition:
// https://docs.google.com/document/d/1NbxiClee0xhqxPmfyHGkxe8JkZNQ2zg2EgEySyJK4jY/edit

package main

import (
	"github.com/labstack/echo/v4"

	ft "github.com/pramineni01/volume_ex/flighttracker"
)

const address = ":8080"

func main() {
	e := echo.New()
	ft.RegisterHandlers(e, ft.Server{})

	e.Logger.Fatal(e.Start(address))
}
