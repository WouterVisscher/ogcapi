package main

import (
	"log"
	"net/http"
	"ogcapi"
)

func main() {

	geojson := GeoJSON{}
	geojson.ReadFile("example.geojson")
	collection := geojson.GetCollectionFromGeoJSON()

	e := ogcapi.Engine{FeatureDatasource: &geojson, Collections: map[string]ogcapi.Collection{geojson.features.Name: collection}}

	dispatchMux := http.NewServeMux()
	dispatchMux.Handle("/", e.GetHandler())

	log.Fatal(http.ListenAndServe(":10000", dispatchMux))
}
