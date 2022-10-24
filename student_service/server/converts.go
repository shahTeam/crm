package server

import (
	"student/domain/group"
	"student/domain/student"
	//"student/id"

	"github.com/shahTeam/crmprotos/studentpb"
)

func toProtoGroups(grs []group.Group) *studentpb.GroupList {
	groups := make([]*studentpb.Group, 0, len(grs))
	for _, item := range grs {
         groups = append(groups, toProtoGroup(item))
	}
	return &studentpb.GroupList{
		Groups: groups,
	}
}

func toProtoStudents(ss []student.Student) *studentpb.StudentList{
	students := make([]*studentpb.StudentResponse, 0, len(ss))
	for _, item := range ss {
		students = append(students, toProtoStudent(item))
	}
	return &studentpb.StudentList{
		Students: students,
	}
}

func toProtoGroup(g group.Group) *studentpb.Group {
	return &studentpb.Group{
		Id: g.ID().String(),
		Name: g.Name(),
		MainTeacherId: g.MainTeacherID().String(),
	}
}

func toProtoStudent(s student.Student) *studentpb.StudentResponse {
	return &studentpb.StudentResponse{
		Id: s.ID().String(),
		FirstName: s.FirstName(),
		LastName: s.LastName(),
		Email: s.Email(),
		PhoneNumber: s.PhoneNumber(),
		Level: s.Level(),
		Password: s.Password(),
		GroupId: s.GroupID().String(),
	}
}