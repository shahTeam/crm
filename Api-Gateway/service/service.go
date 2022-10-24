package service

import (
	"context"


	"github.com/CRM/Api-Gateway/request"
	"github.com/CRM/Api-Gateway/response"
)

func New(teacherService TeacherServiceClient) Service {
       return Service{
		Teacher: teacherService,
	   }
}

type Service struct{
	Teacher TeacherServiceClient
}

type TeacherServiceClient interface {
	RegisterTeacher(context.Context, request.RegisterTeacherRequest) (response.Teacher, error)
	GetTeacher(context.Context, request.GetTeacherRequest) (response.Teacher, error)
	DeleteTeacher(context.Context, request.DeleteTeacherRequest) error
	ListTeacher(context.Context, int32, int32) ([]response.Teacher, error)

	CreateSubject(context.Context, request.CreateSubjectRequest) (response.Subject, error)
	GetSubject(context.Context, request.GetSubjectRequest) (response.Subject, error)
	DeleteSubject(context.Context, request.DeleteSubjectRequest) error
	ListSubject(context.Context, int32, int32) ([]response.Subject, error)
}