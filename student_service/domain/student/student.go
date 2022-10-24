package student

import (
	"errors"
	"fmt"

	"github.com/shahTeam/crmconnect/validate"
	"github.com/google/uuid"
)

var (
	//ErrInvalidStudentData means that data passed for constructing Student structure was bad
	ErrInvalidStudentData = errors.New("invalid student data")
)

// Student represents domain object that holds required info for a student
// All core business logic relevant to students should be done through this struct

type Student struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	level       int32
	password    string
	groupID     uuid.UUID
}

func (s Student) ID() uuid.UUID {
	return s.id
}

func (s Student) FirstName() string {
	return s.firstName
}

func (s Student) LastName() string {
	return s.lastName
}

func (s Student) Email() string {
	return s.email
}

func (s Student) PhoneNumber() string {
	return s.phoneNumber
}

func (s Student) Level() int32 {
	return s.level
}

func (s Student) Password() string {
	return s.password
}

func (s Student) GroupID() uuid.UUID {
	return s.groupID
}

func (s Student) validate() error {
	if s.firstName == "" {
		return fmt.Errorf("%w: empty first name", ErrInvalidStudentData)
	}
	if s.lastName == "" {
		return fmt.Errorf("%w: empty last name", ErrInvalidStudentData)
	}
	if s.password == "" {
		return fmt.Errorf("%w: empty password", ErrInvalidStudentData)
	}
	if err := validate.PhoneNumber(s.phoneNumber); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidStudentData, err)
	}
	if s.level <= 0 || s.level >= 5 {
		return fmt.Errorf("%w: invalid student level data", ErrInvalidStudentData)
	}
	return nil
}

type UnmarshalStudentArgs struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Level       int32
	Password    string
	GroupID     uuid.UUID
}

func UnmarshalStudent(args UnmarshalStudentArgs) (Student, error) {
	s := Student {
		id: args.ID,
		firstName: args.FirstName,
		lastName: args.LastName,
		email: args.Email,
		phoneNumber: args.PhoneNumber,
		level: args.Level,
		password: args.Password,
		groupID: args.GroupID,
	}

	if err := s.validate(); err != nil {
		return Student{}, err
	}
	return s, nil
}

type Limit struct{
	Page, Limit int32
}
