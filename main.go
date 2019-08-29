package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", getAll)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

func getAll(c echo.Context) error {
	return c.String(http.StatusOK, "Hi, Budddd!!!")
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
