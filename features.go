package ogcapi

import (
	"log"
	"net/http"
	"strconv"
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

	param := FeaturesParams{
		CollectionId: s[len(s)-2],
	}

	value := r.URL.Query()["limit"]

	if value != nil {
		limit, err := strconv.Atoi(value[0])
		if err != nil {
			// TODO return error
			// Invalid input/query parameters
			log.Fatalf("Could not marshal collections, got error: %v", err)
		}
		param.Limit = limit
	}

	if fc, err := e.FeatureDatasource.GetFeatureCollection(param); err == nil {
		w.Write(JSONMarshaller(fc))
	}
	// TODO, what now?
	// Send client error msg that features could not be retrieved
}

func (e *Engine) ItemController(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")

	if f, err := e.FeatureDatasource.GetFeature(s[len(s)-3], chi.URLParam(r, "id")); err == nil {
		w.Write(JSONMarshaller(f))
	}
	// TODO, what now?
	// Send client error msg that feature could not be retrieved
}
