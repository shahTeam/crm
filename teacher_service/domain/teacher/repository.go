package teacher

import "context"


type Repository interface {
	CreateTeacher(context.Context, Teacher) error
	GetTeacher( context.Context, By) (Teacher, error)
	UpdateTeacher( context.Context, Teacher) error
    DeleteTeacher( context.Context, By) error
    ListTeachers( context.Context, int32, int32) ([]Teacher, int, error)
}