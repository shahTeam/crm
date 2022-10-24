package student

import "github.com/google/uuid"



type StudentBy interface {
	isStudentBy()
}

type ByID struct{
	ID uuid.UUID
}

func(b ByID) isStudentBy() {}

type ByEamil struct {
	Email string
}

func(b ByEamil) isStudentBy() {}

type ByPhoneNumber struct {
	ByPhoneNumber string
}

func(b ByPhoneNumber) isStudentBy() {}
