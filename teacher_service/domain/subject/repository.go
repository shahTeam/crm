package subject

import "context"

type Repository interface {
	CreateSubject(context.Context, Subject) error
	GetSubject(context.Context, By) (Subject, error)
	UpdateSubject( context.Context, Subject) error
	DeleteSubject(context.Context, By) error
	ListSubjects(context.Context, int32, int32) ([]Subject, int, error)
}
