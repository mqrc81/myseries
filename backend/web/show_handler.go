package web

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/mqrc81/myseries/backend/myseries"
)

var (
	showsListTemplate, showsShowTemplate *template.Template
)

func init() {
	showsListTemplate = template.Must(template.ParseFiles(layout, htmlPath+"shows_list.html"))
	showsShowTemplate = template.Must(template.ParseFiles(layout, htmlPath+"shows_show.html"))
}

type ShowHandler struct {
	store    myseries.Store
	sessions *scs.SessionManager
}

func (h *ShowHandler) List() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}

func (h *ShowHandler) Show() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
