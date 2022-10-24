package service

import (
	"context"
	"teacher/domain/subject"
	"teacher/domain/teacher"
)

type Service struct {
	repo Repository
	subjectFactory subject.Factory
	teacherFactory teacher.Factory
}

func New(repo Repository, subjectFactory subject.Factory, teacherFactory teacher.Factory) Service {
	return Service{
		repo: repo,
		subjectFactory: subjectFactory,
		teacherFactory: teacherFactory,
	}
}

func (s Service) RegisterTeacher(ctx context.Context, t teacher.Teacher) (teacher.Teacher, error) {
	if err := s.repo.CreateTeacher(ctx, t); err != nil {
		return teacher.Teacher{}, err
	}
	return t, nil
}  

func (s Service) CreateSubject(ctx context.Context, sub subject.Subject) (subject.Subject, error) {
	if err := s.repo.CreateSubject(ctx, sub); err != nil {
		return subject.Subject{}, err
	}
	return sub, nil
}

func (s Service) UpdateTeacher(ctx context.Context, t teacher.Teacher) (teacher.Teacher, error) {
	if err := s.repo.UpdateTeacher(ctx, t); err != nil {
		return teacher.Teacher{}, err
	}
	return t, nil
} 

func (s Service) UpdateSubject(ctx context.Context, sub subject.Subject) (subject.Subject, error) {
	if err := s.repo.UpdateSubject(ctx, sub); err != nil {
		return subject.Subject{}, err
	}
	return sub, nil
}

func (s Service) GetTeacher(ctx context.Context, by teacher.By) (teacher.Teacher, error) {
	teachers, err := s.repo.GetTeacher(ctx, by)
	if err != nil {
		return teacher.Teacher{}, err
	}
	return teachers, nil
}

func (s Service) GetSubject(ctx context.Context, sub subject.By) (subject.Subject, error) {
	subjects, err := s.repo.GetSubject(ctx, sub) 
	if err != nil {
		return subject.Subject{}, err
	}
	return subjects, nil
}

func (s Service) DeleteTeacher(ctx context.Context, by teacher.By) error {
	if err := s.repo.DeleteTeacher(ctx, by); err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteSubject(ctx context.Context, sub subject.By) error {
	if err := s.repo.DeleteSubject(ctx, sub); err != nil {
		return err
	}
	return nil
}

func (s Service) ListTeacher(ctx context.Context, page, limit int32) ([]teacher.Teacher, int, error) {
	teachers, count,  err := s.repo.ListTeachers(ctx, page, limit)
	if err != nil {
		return []teacher.Teacher{}, 0, err
	}
    return teachers, count, nil
}

func (s Service) ListSubject(ctx context.Context, page, limit int32) ([]subject.Subject, int, error) {
	subjects, count, err := s.repo.ListSubjects(ctx, page, limit)
	if err != nil {
		return []subject.Subject{}, 0, err
	} 
	return subjects, count, nil
}


