package main

import (
	"log"
	"net/http"
	"ogcapi"
)

func main() {

	geojson := GeoJSONDataSource("example.geojson")
	collection := GetCollectionFromGeoJSON(&geojson)

	renderers := make(map[string]func(http.ResponseWriter, interface{}))
	renderers["json"] = ogcapi.DefaultJSONRender

	e := ogcapi.Engine{
		FeatureDatasource: &geojson,
		Collections:       map[string]ogcapi.Collection{geojson.features.Name: collection},
		Renderers:         renderers,
	}

	dispatchMux := http.NewServeMux()
	dispatchMux.Handle("/", e.GetHandler())

	log.Fatal(http.ListenAndServe(":10000", dispatchMux))
}
