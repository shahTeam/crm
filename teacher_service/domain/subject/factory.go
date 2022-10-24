package subject

import (
	"teacher/pkg/idgen"
	//"github.com/google/uuid"
)

type Factory struct {
	idGenerator idgen.IGenerator
}

func NewFactory(idGenerator idgen.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

func (f Factory) NewSubject(name, description string) (Subject, error) {
	s := Subject{
		id:          f.idGenerator.GeneratorUUID(),
		name:        name,
		description: description,
	}

	if err := s.validate(); err != nil {
		return Subject{}, err
	}
	return s, nil
}
