package web

import (
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/mqrc81/myseries/backend/myseries"
)

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
