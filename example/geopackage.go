package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"ogcapi"

	"github.com/go-spatial/geom/encoding/geojson"
)

// FeatureCollection describes a geoJSON collection feature
type FeatureCollection struct {
	geojson.FeatureCollection
	Name string `json:"name"`
}

type GeoJSON struct {
	features FeatureCollection
}

func (g *GeoJSON) GetCollections() ogcapi.Collections {

	collection := ogcapi.Collection{}
	collection.Id = g.features.Name
	collection.Title = "Title of the " + g.features.Name
	collection.Description = "Description of the " + g.features.Name
	collection.ItemType = "feature"

	self := ogcapi.Link{}

	collections := ogcapi.Collections{Collections: []ogcapi.Collection{collection}, Links: []ogcapi.Link{self}}

	return collections
}

func (g *GeoJSON) GetCollection(c string) ogcapi.Collection {

	collection := ogcapi.Collection{}
	collection.Id = g.features.Name
	collection.Title = "Title of the " + g.features.Name
	collection.Description = "Description of the " + g.features.Name
	collection.ItemType = "feature"

	return collection
}

func (g *GeoJSON) GetFeatureCollection(params ogcapi.FeaturesParams) ogcapi.FeatureCollection {
	return ogcapi.FeatureCollection{}
}

func (g *GeoJSON) GetFeature(id string) ogcapi.Feature {
	return ogcapi.Feature{}
}

func (g *GeoJSON) Init(path string) {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file from path (%v) with error: %v", path, err)
	}

	gjson := FeatureCollection{}

	err = json.Unmarshal(bytes, &gjson)
	if err != nil {
		log.Fatalf("Could not unmarshal file from path (%v) with error: %v", path, err)
	}

	g.features = gjson
}
