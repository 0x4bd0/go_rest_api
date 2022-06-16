package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Stack string

const (
	BackEnd Stack = "BackEnd"
	FrontEnd        = "FrontEnd" 
	Mobile        = "Mobile"
	FullStack        = "FullStack"
)

type Person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Stack  Stack `json:"stack"`
}

var persons = []Person{
	{Name: "John", Stack: BackEnd, Id : "1"},
	{Name: "Doe", Stack: FrontEnd, Id : "2"},
	{Name: "Smith", Stack: FullStack, Id : "3"},
	{Name: "Mark", Stack: Mobile, Id : "4"},
	{Name: "Matt", Stack: Mobile, Id : "5"},
	{Name: "Rick", Stack: BackEnd, Id : "6"},
	{Name: "Morty", Stack: FrontEnd, Id : "7"},
}
var person Person

func GetPersons(c *gin.Context) { 
	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched Data sucessfully",
		"data " : persons,
		"status" : http.StatusOK,
	})
}

func CreatePerson(c *gin.Context) {

	var person Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error": err.Error(),
			"status" : http.StatusBadRequest,
	})
	return
}

	persons = append(persons, person)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created successfully",
		"data " : person,
		"status" : http.StatusCreated,
	})
}

func GetPersonById(id string)(*Person, error) {

	for _, person := range persons {
		if(person.Id == id) {
			return &person, nil
		}
	}

	return nil, errors.New("Person not found")
}

func GetPerson(c *gin.Context) {
	id := c.Param("id")

	person, err := GetPersonById(id)
	
	if(err != nil) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"status" : http.StatusNotFound,
	})
	return
}

c.JSON(http.StatusOK, gin.H{
	"message": "Fetched Data sucessfully",
	"data " : person,
})

}


func main() {

	router := gin.Default()

	router.GET("/persons", GetPersons)
	router.GET("/persons/:id", GetPerson)
	router.POST("/persons", CreatePerson)

	router.Run("localhost:8088")

}