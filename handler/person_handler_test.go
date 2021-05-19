package handler_test

import (
	"mockgen-example/handler"
	"mockgen-example/mocks"
	"mockgen-example/model"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreatePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	person := model.Person{
		Name: "Person Name",
		Age:  20,
	}

	mockRepo := mocks.NewMockPersonRepository(ctrl)
	mockRepo.EXPECT().CreatePerson(person).Return(person)

	personHandler := handler.NewPersonHandler(mockRepo)
	createdPerson := personHandler.CreatePerson(person)

	gomock.Eq(createdPerson).Matches(person)
}
