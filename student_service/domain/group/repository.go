package group

import "context"

type Repository interface{
	CreateGroup(context.Context, Group) error
	GetGroup(context.Context, GroupBy) (Group, error)
	UpdateGroup(context.Context, Group) error
	DeleteGroup(context.Context, GroupBy) error
	ListGroup(context.Context, int32, int32) ([]Group, int, error)
}