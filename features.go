package ogcapi

import (
	"encoding/json"
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

	limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		// TODO return error
		log.Fatalf("Could not marshal collections, got error: %v", err)
	}

	param := FeaturesParams{
		CollectionId: s[len(s)-2],
		Limit:        limit,
	}

	data, err := json.Marshal(e.FeatureDatasource.GetFeatureCollection(param))
	if err != nil {
		log.Fatalf("Could not marshal collections, got error: %v", err)
	}

	w.Write(data)
}

func (e *Engine) ItemController(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")

	if f, err := e.FeatureDatasource.GetFeature(s[len(s)-3], chi.URLParam(r, "id")); err == nil {
		data, err := json.Marshal(f)
		if err != nil {
			log.Fatalf("Could not marshal collections, got error: %v", err)
		}
		w.Write(data)
	}
	// TODO, what now?
}
