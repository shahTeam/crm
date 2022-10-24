package service

import (
	"context"
	"schedule/domain/schedules"
	"time"

	"github.com/google/uuid"
)
type Repository interface {
	CreateSchedule(context.Context, schedules.Schedule) error
	GetSchedule(context.Context, uuid.UUID) (schedules.Schedule, error)
	UpdateSchedule(context.Context, schedules.Schedule) error
	DeleteSchedule(context.Context, uuid.UUID) error 
	GetFullScheduleForGroup(context.Context, uuid.UUID) ([]schedules.Schedule, error)
	GetScheduleForGroup(context.Context, uuid.UUID, time.Weekday) ([]schedules.Schedule, error) 
	GetFullScheduleForTeacher(context.Context, uuid.UUID) ([]schedules.Schedule, error)
	GetScheduleForTeacher(context.Context, uuid.UUID, time.Weekday) ([]schedules.Schedule, error)
}
