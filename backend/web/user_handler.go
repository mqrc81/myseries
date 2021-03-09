package web

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/mqrc81/myseries/backend/myseries"
)

var (
	usersProfileTemplate *template.Template
)

func init() {
	usersProfileTemplate = template.Must(template.ParseFiles(layout, htmlPath+"users_profile.html"))
}

type UserHandler struct {
	store    myseries.Store
	sessions *scs.SessionManager
}

func (h *UserHandler) Profile() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
