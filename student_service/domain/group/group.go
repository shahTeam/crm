package group

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrInvalidGroupData = errors.New("invalid group data")
)

//Group represents domain object that holds required info for a group
//All core business logic relevant to groups should be done through this struct
type Group struct {
	id            uuid.UUID
	name          string
	mainTeacherID uuid.UUID
}

func (g Group) ID() uuid.UUID {
	return g.id
}

func (g Group) Name() string {
	return g.name
}

func (g Group) MainTeacherID() uuid.UUID {
	return g.mainTeacherID
}

func (g Group) validate() error {
	if g.name == "" {
		return fmt.Errorf("%w: name is empty", ErrInvalidGroupData)
	}
	return nil
}

type UnmarshalGroupArgs struct {
	ID            uuid.UUID
	Name          string
	MainTeacherID uuid.UUID
}

func UnmarshalGroup(args UnmarshalGroupArgs) (Group, error) {
	g := Group{
		id: args.ID,
		name: args.Name,
		mainTeacherID: args.MainTeacherID,
	}

	if err:= g.validate(); err != nil {
		return Group{}, err
	} 
	return g, nil
}

type Limit struct {
	page, limit int32
}