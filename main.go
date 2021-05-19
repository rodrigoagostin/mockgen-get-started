package main

import (
	"fmt"
	"mockgen-get-started/handler"
	"mockgen-get-started/model"
	"mockgen-get-started/repository"
)

func main() {
	personHandler := handler.NewPersonHandler(repository.NewPersonRepository())
	person := personHandler.CreatePerson(model.Person{Name: "Person Name", Age: 44})
	fmt.Println(person)
}
