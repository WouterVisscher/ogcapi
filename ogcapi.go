package ogcapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Feature Model
type FeatureDatasource interface {
	GetFeatureCollection(FeaturesParams) (RawFeatureCollection, error)
	GetFeature(string, string) (RawFeature, error)
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

func (e *Engine) GetCollection(collectionId string) (Collection, error) {

	collection := e.Collections[collectionId]

	return collection, nil
}
