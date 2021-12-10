package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"ogcapi"
)

func GetCollectionFromGeoJSON(g *GeoJSON) ogcapi.Collection {

	collection := ogcapi.Collection{}
	collection.Id = g.features.Name
	collection.Title = "Title of the " + g.features.Name
	collection.Description = "Description of the " + g.features.Name
	collection.ItemType = "feature"

	return collection
}

func GeoJSONDataSource(path string) GeoJSON {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file from path (%v) with error: %v", path, err)
	}

	gjson := FeatureCollection{}

	err = json.Unmarshal(bytes, &gjson)
	if err != nil {
		log.Fatalf("Could not unmarshal file from path (%v) with error: %v", path, err)
	}

	return GeoJSON{features: gjson}
}
