package ogcapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ObjectModel interface {
	AddLinks([]Link)
}

// Feature Model
type FeatureDatasource interface {
	GetFeatureCollection(FeaturesParams) (ObjectModel, error)
	GetFeature(string, string) (ObjectModel, error)
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

func (e *Engine) GetCollections() (Collections, error) {

	collections := []Collection{}

	for _, collection := range e.Collections {
		collections = append(collections, collection)
	}

	return Collections{Collections: collections}, nil
}

func (e *Engine) GetCollection(id string) (Collection, error) {

	collection := e.Collections[id]

	return collection, nil
}
