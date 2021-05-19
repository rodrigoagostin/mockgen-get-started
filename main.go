package main

import (
	"fmt"
	"mockgen-example/handler"
	"mockgen-example/model"
	"mockgen-example/repository"
)

func main() {
	personHandler := handler.NewPersonHandler(repository.NewPersonRepository())
	person := personHandler.CreatePerson(model.Person{"Person Name", 44})
	fmt.Println(person)
}
