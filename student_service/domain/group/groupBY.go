package group

import "github.com/google/uuid"

type GroupBy interface {
	isGroupBy()
}

type ByID struct {
	ID uuid.UUID
}

func (b ByID) isGroupBy() {}

type ByName struct{
	Name string
}

func (b ByName) isGroupBy() {}

