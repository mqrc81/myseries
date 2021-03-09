package web

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"

	"github.com/mqrc81/myseries/backend/myseries"
)

const (
	staticPath = "../frontend/static"
	htmlPath   = "../frontend/html/templates/"
	layout     = "../frontend/html/layout.html"
)

var (
	homeTemplate *template.Template
)

func init() {
	homeTemplate = template.Must(template.ParseFiles(layout, htmlPath+"home.html"))
}

func NewHandler(store myseries.Store, sessions *scs.SessionManager, csrfKey []byte) *Handler {

	h := &Handler{
		Mux:      chi.NewMux(),
		store:    store,
		sessions: sessions,
	}

	auth := AuthHandler{store, sessions}
	series := ShowHandler{store, sessions}

	h.Use(middleware.Logger)
	h.Use(csrf.Protect(csrfKey, csrf.Secure(false)))
	h.Use(sessions.LoadAndSave)
	// h.Use(h.withUser)

	h.Get("/", h.Home())

	h.Route("/series", func(r chi.Router) {
		r.Get("/", series.List())
		r.Get("/{show_id}", series.Show())
	})

	h.Post("/login", auth.Login())

	return h
}

type Handler struct {
	*chi.Mux
	store    myseries.Store
	sessions *scs.SessionManager
}

func (h *Handler) Home() http.HandlerFunc {

	type data struct {
		Genres []myseries.Genre
	}

	return func(res http.ResponseWriter, req *http.Request) {
		genres, err := h.store.GetGenres()
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = homeTemplate.Execute(res, data{Genres: genres}); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
}
