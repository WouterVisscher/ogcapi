package ogcapi

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (e *Engine) CollectionsHandler() http.Handler {

	r := chi.NewRouter()

	r.Route("/collections", func(r chi.Router) {
		r.Get("/", e.CollectionsController)

		for collectionId, collection := range e.Collections {
			r.Route("/"+collectionId, func(r chi.Router) {
				r.Get("/", e.CollectionController)

				// TODO maybe switch when tiles/styles/maps/... are added
				if collection.ItemType == "feature" {
					r.Handle("/items*", e.FeatureHandler())
				}
			})
		}
	})

	return r
}

func (e *Engine) CollectionsController(w http.ResponseWriter, r *http.Request) {

	c, err := e.GetCollections()
	if err != nil {
		// TODO, what now?
		// Send client error msg that getting collections went wrong
		log.Fatalf("not collections found, got error: %v", err)
	}

	key := "json"

	e.GetRenderer(key)(w, c)
}

func (e *Engine) CollectionController(w http.ResponseWriter, r *http.Request) {

	// TODO may think about something beter...
	// but path is already validated by the router
	s := strings.Split(r.URL.Path, "/")

	c, err := e.GetCollection(s[len(s)-1])
	if err == nil {
		// TODO, what now?
		// Send client error msg that collection could not be retrieved, while it was defined
		log.Fatalf("not collectionc found, got error: %v", err)
	}

	key := "json"

	e.GetRenderer(key)(w, c)
}
