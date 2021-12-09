package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"ogcapi"
	"strconv"

	"github.com/go-spatial/geom/encoding/geojson"
)

type FeatureCollection struct {
	geojson.FeatureCollection
	Name  string `json:"name"`
	Links []ogcapi.Link
}

func (featurecollection *FeatureCollection) AddLinks(links []ogcapi.Link) {
	featurecollection.Links = links
}

type Feature struct {
	geojson.Feature
	Links []ogcapi.Link
}

func (feature *Feature) AddLinks(links []ogcapi.Link) {
	feature.Links = links
}

type GeoJSON struct {
	features FeatureCollection
}

func (g *GeoJSON) GetCollectionFromGeoJSON() ogcapi.Collection {

	collection := ogcapi.Collection{}
	collection.Id = g.features.Name
	collection.Title = "Title of the " + g.features.Name
	collection.Description = "Description of the " + g.features.Name
	collection.ItemType = "feature"

	return collection
}

func (g *GeoJSON) GetFeatureCollection(params ogcapi.FeaturesParams) (ogcapi.ObjectModel, error) {

	features := []geojson.Feature{}

	for _, f := range g.features.Features {
		feature := geojson.Feature{}
		feature.ID = f.ID
		feature.Type = f.Type
		feature.Properties = f.Properties
		feature.Geometry = f.Geometry
		features = append(features, feature)
	}

	featurecollection := geojson.FeatureCollection{Features: features}

	return &FeatureCollection{FeatureCollection: featurecollection, Name: params.CollectionId}, nil
}

func (g *GeoJSON) GetFeature(collectionid, id string) (ogcapi.ObjectModel, error) {

	feature := geojson.Feature{}

	for _, f := range g.features.Features {

		if strconv.Itoa(int(*f.ID)) == id {

			feature.ID = f.ID
			feature.Type = f.Type
			feature.Properties = f.Properties
			feature.Geometry = f.Geometry
			return &Feature{Feature: feature}, nil
		}
	}

	return &Feature{Feature: feature}, errors.New("No feature found with id: " + id)

}

func (g *GeoJSON) ReadFile(path string) {

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
