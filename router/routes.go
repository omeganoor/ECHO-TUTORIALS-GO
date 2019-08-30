package router

import (
	"github.com/ECHO-TUTORIALS-GO/controller"
	"github.com/labstack/echo"
)

// RouteParams is route parameter
type RouteParams struct {
	*echo.Echo
	controller.Employee
}

// Routes of API
func InitRouter(p RouteParams) {
	group := p.Group("/tix-echo-go")
	route := group.Group("/employee")
	route.GET("/save", p.SaveEmployee)
	route.GET("/", p.GetAll)
	route.GET("/users/:id", p.GetUser)
	route.GET("/show", p.Show)
}
