package idgen

import "github.com/google/uuid"

type IGenerator interface {
	GeneratorUUID() uuid.UUID
}

type Generator struct{}

func (g Generator) GeneratorUUID() uuid.UUID {
	return uuid.New()
}