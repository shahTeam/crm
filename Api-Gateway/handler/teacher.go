package handler

import (
	"context"
	"net/http"

	"github.com/CRM/Api-Gateway/pkg/httperr"
	"github.com/CRM/Api-Gateway/request"
	"github.com/go-chi/render"
)

func (h Handler) RegisterTeacher(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterTeacherRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}
teacher, err := h.service.Teacher.RegisterTeacher(context.Background(), req)
if err != nil {
	httperr.Handle(w, r, err)
	return
}

render.Status(r, http.StatusCreated)
render.JSON(w, r, render.M{
	"teacher": teacher,
	"token": auth
})
	
}