package student

import (
	"context"

	"github.com/google/uuid"
)



type Repository interface {
	CreateStudent(context.Context, Student) error
	GetStudent(context.Context, StudentBy) (Student, error)
	UpdateStudent(context.Context, Student) error
	DeleteStudent(context.Context, StudentBy) error
	ListStudent(context.Context, int32, int32) ([]Student, int, error)
	GetStudentsByGroup(context.Context, uuid.UUID) ([]Student, error)
}