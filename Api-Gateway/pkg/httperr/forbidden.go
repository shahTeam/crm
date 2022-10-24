package httperr

import (
	"net/http"

	"github.com/go-chi/render"
)

func Forbidden(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusForbidden)
	render.JSON(w,r, render.M{
		"error": ErrForbidden.Error(),
		"msg": msg,
	})
}