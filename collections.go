package ogcapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (e *Engine) CollectionsHandler() http.Handler {

	r := chi.NewRouter()

	r.Route("/collections", func(r chi.Router) {
		r.Get("/", e.CollectionsController)

		for id, collection := range e.Collections {
			r.Route("/"+id, func(r chi.Router) {
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

	data, err := json.Marshal(e.GetCollections())
	if err != nil {
		log.Fatalf("Could not marshal collections, got error: %v", err)
	}

	w.Write(data)
}

func (e *Engine) CollectionController(w http.ResponseWriter, r *http.Request) {

	// TODO may think about something beter...
	// but path is already validated by the router
	s := strings.Split(r.URL.Path, "/")

	data, err := json.Marshal(e.GetCollection(s[len(s)-1]))
	if err != nil {
		log.Fatalf("Could not marshal collections, got error: %v", err)
	}

	w.Write(data)
}
