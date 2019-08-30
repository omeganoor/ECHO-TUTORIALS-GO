package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ECHO-TUTORIALS-GO/entity"

	"github.com/ECHO-TUTORIALS-GO/module"
	"github.com/labstack/echo"
)

var ctx = context.Background()

type Employee struct {
}

func (e Employee) SaveEmployee(c echo.Context) error {
	var colName = "employee"
	coll, err := module.GetConnection(colName)
	if err != nil {
		log.Fatal(err.Error())
	}
	emp := entity.Employee{
		Name:     "kims",
		Position: "SBE",
		Salary:   100000,
	}
	fmt.Println(emp)
	result, _ := coll.InsertOne(ctx, emp)

	return c.JSON(http.StatusOK, result)
}

func (e Employee) GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "Hi, Budddd!!!")
}

// e.GET("/users/:id", getUser)
func (e Employee) GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func (e Employee) Show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
