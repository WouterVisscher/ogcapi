package main

import (
	"errors"
	"ogcapi"
	"strconv"

	"github.com/go-spatial/geom/encoding/geojson"
)

type FeatureCollection struct {
	geojson.FeatureCollection
	Name string
}

type GeoJSON struct {
	features FeatureCollection
}

// Model: GetFeatureCollection
func (g *GeoJSON) GetFeatureCollection(params ogcapi.FeaturesParams) (ogcapi.RawFeatureCollection, error) {

	features := []geojson.Feature{}

	for _, f := range g.features.Features {
		feature := geojson.Feature{}
		feature.ID = f.ID
		feature.Type = f.Type
		feature.Properties = f.Properties
		feature.Geometry = f.Geometry
		features = append(features, feature)
	}

	return ogcapi.RawFeatureCollection{FeatureCollection: geojson.FeatureCollection{Features: features}, Name: g.features.Name}, nil
}

// Model: GetFeature
func (g *GeoJSON) GetFeature(collectionid, id string) (ogcapi.RawFeature, error) {

	feature := geojson.Feature{}

	for _, f := range g.features.Features {

		if strconv.Itoa(int(*f.ID)) == id {

			feature.ID = f.ID
			feature.Type = f.Type
			feature.Properties = f.Properties
			feature.Geometry = f.Geometry
			return ogcapi.RawFeature{Feature: feature}, nil
		}
	}

	return ogcapi.RawFeature{}, errors.New("No feature found with id: " + id)
}

// XML View: GetFeatureCollection

// XML View: GetFeature
