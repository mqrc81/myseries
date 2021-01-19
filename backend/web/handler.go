package web

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/mqrc81/myseries/myseries"
)

var (
	Templates = make(map[string]*template.Template)
)

func init() {
	// path := "frontend/templates/"
	// layout := path + "layout.html"
	// Templates["home"] = template.Must(template.ParseFiles(layout, path+"home.html"))
}

func NewHandler(store myseries.Store) *Handler {

	handler := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	handler.Use(middleware.Logger)

	handler.Get("/", handler.Home())

	return handler
}

type Handler struct {
	*chi.Mux
	store myseries.Store
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

		if err := Templates["home"].Execute(res, data{
			Genres: genres,
		}); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
