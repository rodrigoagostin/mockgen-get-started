package repository

import "mockgen-example/model"

type PersonRepository interface {
	CreatePerson(person model.Person) model.Person
}

type personRepository struct{}

func NewPersonRepository() PersonRepository {
	return &personRepository{}
}

func (r *personRepository) CreatePerson(person model.Person) model.Person {
	return person
}
