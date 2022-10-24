package service

import (
	"context"
	"student/domain/group"
	"student/domain/student"

	"github.com/google/uuid"
)

func New(repo Repository, studentFactory student.Factory, groupFactory group.Factory) Service {
	return Service{
		repo: repo,
		studentFactory: studentFactory,
		groupFactory: groupFactory,
	}
}

//Service binds all the core logic together provideing necessary methods for APIs
type Service struct {
	repo           Repository
	studentFactory student.Factory
	groupFactory   group.Factory
}

func (s Service) ListGroup(ctx context.Context, pages, limit int32) ([]group.Group, int, error) {
	return s.repo.ListGroup(ctx, pages, limit)
}

func (s Service) DeleteGroup(ctx context.Context, by group.GroupBy) error {
	return s.repo.DeleteGroup(ctx, by)
}

func (s Service) UpdateGroup(ctx context.Context, g group.Group) error {
	return s.repo.UpdateGroup(ctx, g)
}

func (s Service) ListStudent(ctx context.Context, page, limit int32) ([]student.Student, int, error) {
	return s.repo.ListStudent(ctx, page, limit)
}

func (s Service) DeleteStudent(ctx context.Context, by student.StudentBy) error {
	return s.repo.DeleteStudent(ctx, by)
}

func (s Service) UpdateStudent(ctx context.Context, st student.Student) error {
	return s.repo.UpdateStudent(ctx, st)
}
func (s Service) GetGroup(ctx context.Context, by group.GroupBy) (group.Group, error) {
	return s.repo.GetGroup(ctx, by)
} 

func (s Service) GetStudent(ctx context.Context, by student.StudentBy) (student.Student, error) {
	return s.repo.GetStudent(ctx, by)
}

func (s Service) RegisterStudent(ctx context.Context, st student.Student) error {
	return s.repo.CreateStudent(ctx, st)
}

func (s Service) CreateGroup(ctx context.Context, g group.Group) error {
	return s.repo.CreateGroup(ctx, g)
}

func (s Service) GetStudentsByGroup(ctx context.Context, groupID uuid.UUID) ([]student.Student, error) {
	return s.repo.GetStudentsByGroup(ctx, groupID)
}
