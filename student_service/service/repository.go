package service

import (
	"student/domain/group"
	"student/domain/student"
)

//Repository defines methods that should be present in our storage provider
type Repository interface{
	student.Repository
	group.Repository
}