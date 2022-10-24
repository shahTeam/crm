package teacher

import (
	"teacher/pkg/idgen"

	"github.com/google/uuid"
)

type Factory struct {
	idGenerator idgen.IGenerator 
}

func NewFactory(idGenerator idgen.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

func (f Factory) NewFactory(
	firstName, lastName, email, phoneNumber, password string, 
	subjectID uuid.UUID,
) (Teacher, error) {
	t := Teacher {
		id: f.idGenerator.GeneratorUUID(),
		firstName: firstName,
		lastName: lastName,
		email: email,
		phoneNumber: phoneNumber,
		password: password,
		subjectID: subjectID,
	}

	if err := t.validate(); err != nil {
           return Teacher{}, err
	}
	return t, nil
}