package student

import (
	"fmt"
	"student/id"

	"github.com/google/uuid"
	//"github.com/shahTeam/crmconnect/id"
)

//Factory constructs new student domain objects
func NewFactory(idGenerator id.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

type Factory struct {
	idGenerator id.IGenerator
}



//NewStudent is a constructor that checks if the provided data for student is valid or not 
// New student objects can only be created through this constructor which ensures everything is valid

func (f Factory) NewStudent(firstName, lastName, email, phoneNumber string, level int32, password string, groupID uuid.UUID) (Student, error) {
	s := Student{
		id: f.idGenerator.GenerateUUID(),
		firstName: firstName,
		lastName: lastName,
		email: email,
		phoneNumber: phoneNumber,
		level: level,
		password: password,
		groupID: groupID,
	}
	fmt.Println(s.id)
	if err := s.validate(); err != nil {
		return Student{}, err
	}
	return s, nil
}