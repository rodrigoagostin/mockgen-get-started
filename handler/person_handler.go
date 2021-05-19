package handler

import (
	"mockgen-get-started/model"
	"mockgen-get-started/repository"
)

type PersonHandler interface {
	CreatePerson(p model.Person) model.Person
}

type personHandler struct {
	repo repository.PersonRepository
}

func NewPersonHandler(repo repository.PersonRepository) PersonHandler {
	return &personHandler{
		repo,
	}
}

func (h *personHandler) CreatePerson(p model.Person) model.Person {
	createdPerson := h.repo.CreatePerson(p)

	return createdPerson
}
