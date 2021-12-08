package ogcapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (e *Engine) CoreHandler() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", ProcesLandingPage)
	r.Get("/api", ProcesAPI)
	r.Get("/conformance", ProcesConformance)

	return r
}

func ProcesConformance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello conformance"))
}

func ProcesAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello api"))
}

func ProcesLandingPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello landingpage"))
}
