package main

import (
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
	Name string `json:"name"`
	Stack  Stack `json:"stack"`
}

var persons = []Person{
	{Name: "John", Stack: BackEnd},
	{Name: "Doe", Stack: FrontEnd},
	{Name: "Smith", Stack: FullStack},
	{Name: "Mark", Stack: Mobile},
	{Name: "Matt", Stack: Mobile},
	{Name: "Rick", Stack: BackEnd},
	{Name: "Morty", Stack: FrontEnd},
}

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

func main() {

	router := gin.Default()

	router.GET("/persons", GetPersons)
	router.POST("/persons", CreatePerson)

	router.Run("localhost:8088")

}