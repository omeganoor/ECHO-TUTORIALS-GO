package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ECHO-TUTORIALS-GO/entity"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func main() {
	e := echo.New()
	e.GET("/", getAll)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.GET("/save", save)
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

func save(c echo.Context) error {
	var colName = "employee"
	coll, err := connect(colName)
	if err != nil {
		log.Fatal(err.Error())
	}
	emp := entity.Employee{
		Name:     "kims",
		Position: "SBE",
		Salary:   100000,
	}
	fmt.Println(emp)
	var result entity.Employee
	filter := bson.D{{"name", "Kimsha"}}

	coll.InsertOne(ctx, emp)
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return c.JSON(http.StatusOK, result)
}

func connect(colName string) (*mongo.Collection, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to MongoDB closed.")

	return client.Database("Echo_Tutorials").Collection(colName), nil
}
