package main

import (
	"github.com/gin-gonic/gin"
)

type Stack string

const (
	BackEnd Stack = "BackEnd"
	FrontEnd        = "FrontEnd" 
	Mobile        = "Mobile"
	FullStack        = "FullStack"
)

type Team struct { 
	Id int `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	Name string `json:"name"`
	Stack  Stack `json:"stack"`
	Team	int `json:"team"`
}

func main() {

	router := gin.Default()
	router.run("localhost:8088")

}