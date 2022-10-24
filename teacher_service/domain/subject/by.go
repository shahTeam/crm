package subject

import "github.com/google/uuid"

type By interface {
	isSubjectBy()
}

type ByID struct {
	ID uuid.UUID
}

func (b ByID) isSubjectBy() {}

type ByName struct {
	Name string
}

func (b ByName) isSubjectBy() {}
