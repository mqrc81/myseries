package web

import (
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/mqrc81/myseries/backend/myseries"
)

type UserHandler struct {
	store    myseries.Store
	sessions *scs.SessionManager
}

func (h *UserHandler) Profile() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
