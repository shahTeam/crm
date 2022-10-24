package handler

import (
	"github.com/CRM/Api-Gateway/service"
)

func New(svc service.Service) Handler {
	return Handler{
        service: svc,
	}
}

type Handler struct {
	service service.Service
}