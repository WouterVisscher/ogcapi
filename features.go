package ogcapi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const itemKey = key("item")

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

	r.Use(middleware.Logger)

	r.Route("/items", func(r chi.Router) {
		r.Get("/", e.ProcesItems)
		r.Route("/{item}", func(r chi.Router) {
			r.Use(ItemCtx)
			r.Get("/", e.ProcesItem)
		})
	})

	return r
}

func ItemCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), itemKey, chi.URLParam(r, "item"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (e *Engine) ProcesItems(w http.ResponseWriter, r *http.Request) {
	collection := r.Context().Value(collectionKey).(string)
	w.Write([]byte("hello items from " + collection))
}

func (e *Engine) ProcesItem(w http.ResponseWriter, r *http.Request) {
	collection := r.Context().Value(collectionKey).(string)
	item := r.Context().Value(itemKey).(string)

	w.Write([]byte("hello Item " + item + " from " + collection + " link: " + r.Host + r.RequestURI))
}
