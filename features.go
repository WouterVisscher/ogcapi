package ogcapi

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type FeaturesParams struct {
	CollectionId string
	Limit        int
	Offset       int
	ContentType  string
	Bbox         [4]float64
	Datetime     string
}

func (e *Engine) FeatureHandler() http.Handler {

	r := chi.NewRouter()

	r.Route("/items", func(r chi.Router) {
		r.Get("/", e.ItemsController)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", e.ItemController)
		})
	})

	return r
}

func (e *Engine) ItemsController(w http.ResponseWriter, r *http.Request) {

	s := strings.Split(r.URL.Path, "/")
	w.Write([]byte("hello items from " + s[len(s)-2]))
}

func (e *Engine) ItemController(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")

	w.Write([]byte("hello Item " + chi.URLParam(r, "id") + " from " + s[len(s)-3] + " link: " + r.Host + r.RequestURI))
}
