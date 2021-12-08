package main

import (
	"encoding/json"
	"errors"
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

func (g *GeoJSON) GetCollection() ogcapi.Collection {

	collection := ogcapi.Collection{}
	collection.Id = g.features.Name
	collection.Title = "Title of the " + g.features.Name
	collection.Description = "Description of the " + g.features.Name
	collection.ItemType = "feature"

	// self := ogcapi.Link{}

	return collection
}

func (g *GeoJSON) GetFeatureCollection(params ogcapi.FeaturesParams) ogcapi.FeatureCollection {

	fc := []*ogcapi.Feature{}

	for _, f := range g.features.Features {
		feature := ogcapi.Feature{}
		feature.ID = f.Properties["id"]
		feature.Type = f.Type
		feature.Properties = f.Properties
		feature.Geometry = f.Geometry
		fc = append(fc, &feature)
	}

	return ogcapi.FeatureCollection{Features: fc, Type: "FeatureCollection"}
}

func (g *GeoJSON) GetFeature(collectionid, id string) (ogcapi.Feature, error) {

	feature := ogcapi.Feature{}

	for _, f := range g.features.Features {

		if f.Properties["id"] == id {

			feature.ID = f.Properties["id"]
			feature.Type = f.Type
			feature.Properties = f.Properties
			feature.Geometry = f.Geometry
			return feature, nil
		}
	}

	return feature, errors.New("No feature found with id: " + id)

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
