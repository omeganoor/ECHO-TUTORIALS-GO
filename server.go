package main

import (
	srv "github.com/ECHO-TUTORIALS-GO/controller"
	rtg "github.com/ECHO-TUTORIALS-GO/router"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	r := rtg.RouteParams{
		e,
		srv.Employee{},
	}
	rtg.InitRouter(r)
	e.Logger.Fatal(e.Start(":1323"))
}
