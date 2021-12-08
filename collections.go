package ogcapi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const collectionKey = key("collection")

func (e *Engine) CollectionsHandler() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/collections", func(r chi.Router) {
		r.Get("/", e.ProcesCollections)
		r.Route("/{collection}", func(r chi.Router) {
			r.Use(CollectionCtx)
			r.Get("/", e.ProcesCollection)

			r.Handle("/items*", e.FeatureHandler())
		})
	})

	return r
}

func CollectionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), collectionKey, chi.URLParam(r, "collection"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (e *Engine) ProcesCollections(w http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(e.GetCollections())
	if err != nil {
		log.Fatalf("Could not marshal collections, got error: %v", err)
	}

	w.Write(data)
}

func (e *Engine) ProcesCollection(w http.ResponseWriter, r *http.Request) {
	collection := r.Context().Value(collectionKey).(string)

	data, err := json.Marshal(e.GetCollection(collection))
	if err != nil {
		log.Fatalf("Could not marshal collections, got error: %v", err)
	}

	w.Write(data)
}
