package httperr

import (
	
	"net/http"

	"github.com/go-chi/render"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusUnauthorized)
	render.JSON(w, r, render.M{
		"error": ErrUnauthorized.Error(),
		"message": msg,
	})
}