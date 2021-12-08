package ogcapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Feature Model
type FeatureDatasource interface {
	GetFeatureCollection(FeaturesParams) FeatureCollection
	GetFeature(string) Feature
}

type Engine struct {
	Collections       map[string]Collection
	FeatureDatasource FeatureDatasource
}

func (e *Engine) GetHandler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Handle("/*", e.CoreHandler())
	r.Handle("/collections*", e.CollectionsHandler())

	return r
}

func (e *Engine) GetCollections() Collections {

	collections := []Collection{}

	for _, collection := range e.Collections {
		collections = append(collections, collection)
	}

	return Collections{Collections: collections}
}

func (e *Engine) GetCollection(id string) Collection {

	collection := e.Collections[id]

	return collection
}
