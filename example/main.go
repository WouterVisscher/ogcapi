package main

import (
	"log"
	"net/http"
	"ogcapi"
)

func main() {

	geojson := GeoJSON{}
	geojson.Init("example.geojson")

	e := ogcapi.Engine{Datasource: &geojson}

	dispatchMux := http.NewServeMux()
	dispatchMux.Handle("/", e.GetHandler())

	log.Fatal(http.ListenAndServe(":10000", dispatchMux))
}
