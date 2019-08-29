package entity

type Employee struct {
	Name     string `bson:"name"`
	Position string `bson:"position"`
	Salary   int    `bson:"salary"`
}
