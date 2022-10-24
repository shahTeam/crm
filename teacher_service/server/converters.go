package server

import (
	"teacher/domain/subject"
	"teacher/domain/teacher"

	"github.com/shahTeam/crmprotos/teacherpb"
)

func toProtoTeacher(t teacher.Teacher) *teacherpb.Teacher {
	return &teacherpb.Teacher{
		Id: t.ID().String(),
		FirstName: t.FirstName(),
		LastName: t.LastName(),
		Email: t.Email(),
		PhoneNumber: t.PhoneNumber(),
		SubjectId: t.SubjectID().String(),
	}
}

func toProtoSubject(sub subject.Subject) *teacherpb.Subject {
	return &teacherpb.Subject{
		Id: sub.ID().String(),
		Name: sub.Name(),
		Description: sub.Description(),
	}
}