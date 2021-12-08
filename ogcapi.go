package ogcapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Datasource interface {
	GetCollections() Collections
	GetCollection(string) Collection
	GetFeatureCollection(FeaturesParams) FeatureCollection
	GetFeature(string) Feature
}

type Engine struct {
	Collections map[string]Collection
	Datasource  Datasource
}

func (e *Engine) GetHandler() http.Handler {
	r := chi.NewRouter()

	r.Handle("/*", e.CoreHandler())
	r.Handle("/collections*", e.CollectionsHandler())

	return r
}
