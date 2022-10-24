package repository

import (
	"teacher/domain/subject"
	"teacher/domain/teacher"
)

func toRepositoryTeacher(t teacher.Teacher) Teacher {
	return Teacher{
		ID: t.ID(),
		FirstName: t.FirstName(),
		LastName: t.LastName(),
		Email: t.Email(),
		PhoneNumber: t.PhoneNumber(),
		Password: t.Password(),
		SubjectID: t.SubjectID(),
	}
}

func toRepositorySubject(s subject.Subject) Subject {
	return Subject{
		ID: s.ID(),
		Name: s.Name(),
		Description: s.Description(),
	}
}

func toDomainTeachers(repoTeachers []Teacher) ([]teacher.Teacher, error) {
	teachers := make([]teacher.Teacher, 0, len(repoTeachers))
	for _, rerepoTeacher := range repoTeachers {
		t, err := teacher.UnmarshalTeacher(teacher.UnmarshalTeacherArgs(rerepoTeacher))
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}
	return teachers, nil
}

func toDomainSubjects(repoSubjects []Subject) ([]subject.Subject, error) {
	subjects := make([]subject.Subject, 0, len(repoSubjects))
	for _, reporepoSubject := range repoSubjects {
		s, err := subject.UnmarshalSubject(subject.UnmarshalSubjectArgs(reporepoSubject))
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}