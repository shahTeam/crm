package repository

import (
	"student/domain/group"
	"student/domain/student"

)



func toRepositoryStudents(s student.Student) Student {
	return Student{
		ID: s.ID(),
		FirstName: s.FirstName(),
		LastName: s.LastName(),
		Email: s.Email(),
		PhoneNumber: s.PhoneNumber(),
		Level: s.Level(),
		Password: s.Password(),
		GroupID: s.GroupID(),
	}
}

func toRepositoryGroup(g group.Group) Group {
	return Group{
		ID: g.ID(),
		Name: g.Name(),
		MainTeacherID: g.MainTeacherID(),
	}
}

func toDomanStudents(repoStudents []Student) ([]student.Student, error) {
	students := make([]student.Student, 0, len(repoStudents))
	for _, repoStudent := range repoStudents {
		t, err := student.UnmarshalStudent(student.UnmarshalStudentArgs(repoStudent))
		if err != nil {
			return []student.Student{}, err
		}
		students = append(students, t)
	}
	return students, nil
}

func toDomanGroups(repoGroups []Group) ([]group.Group, error) {
	groups := make([]group.Group, 0, len(repoGroups))
	for _, reporepoGroup := range repoGroups {
		g, err := group.UnmarshalGroup(group.UnmarshalGroupArgs(reporepoGroup))
		if err != nil {
			return []group.Group{}, err
		}
		groups = append(groups, g)
	}
	return groups, nil
}