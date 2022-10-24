package service

import (
	"teacher/domain/subject"
	"teacher/domain/teacher"
)

//Repository defines methods that should be present in our storage provider
type Repository interface {
	teacher.Repository
	subject.Repository
}