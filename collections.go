package ogcapi

import (
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

	if c, err := e.GetCollections(); err == nil {
		w.Write(JSONMarshaller(&c))
	}
	// TODO, what now?
	// Send client error msg that getting collections went wrong
}

func (e *Engine) CollectionController(w http.ResponseWriter, r *http.Request) {

	// TODO may think about something beter...
	// but path is already validated by the router
	s := strings.Split(r.URL.Path, "/")

	if c, err := e.GetCollection(s[len(s)-1]); err == nil {
		w.Write(JSONMarshaller(&c))
	}
	// TODO, what now?
	// Send client error msg that collection could not be retrieved, while it was defined
}
