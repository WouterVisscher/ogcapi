package main

import (
	"log"
	"net/http"
	"ogcapi"
)

func main() {

	geojson := GeoJSON{}
	geojson.Init("example.geojson")
	collection := geojson.GetCollection()

	e := ogcapi.Engine{FeatureDatasource: &geojson, Collections: map[string]ogcapi.Collection{geojson.features.Name: collection}}

	dispatchMux := http.NewServeMux()
	dispatchMux.Handle("/", e.GetHandler())

	log.Fatal(http.ListenAndServe(":10000", dispatchMux))
}
