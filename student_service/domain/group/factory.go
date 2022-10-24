package group

import (
	"github.com/google/uuid"
	"github.com/shahTeam/crmconnect/id"
)

//Newfactory initializes a new Factory
func NewFactory(idGenerater id.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerater,
	}
}

//Factory is a struct that creates new instance of Group. Read more about factory pattern
type Factory struct {
	idGenerator id.IGenerator
}

//NewGroup creates a new instance of group with given data, checking its validaty
func (f Factory) NewGroup(name string, mainTeacherID uuid.UUID) (Group, error) {
	g := Group{
		id: f.idGenerator.GenerateUUID(),
		name: name,
		mainTeacherID: mainTeacherID,
	}
	if err := g.validate(); err != nil {
		return Group{}, err
	}
	return g, nil
}

