package ogcapi

import (
	"encoding/json"
	"log"

	"github.com/go-spatial/geom/encoding/geojson"
)

type Collections struct {
	Collections []Collection `json:"collections"`
	Links       []Link       `json:"links"`
}

type Collection struct {
	/* identifier of the collection used, for example, in URIs	*/
	Id string `json:"id"`
	/* The extent of the features in the collection. In the Core only spatial and temporal
	extents are specified. Extensions may add additional members to represent other
	extents, for example, thermal or pressure ranges.	*/
	Extent *Extent `json:"extent,omitempty"`
	/* a description of the features in the collection	*/
	Description string `json:"description,omitempty"`
	/* human readable title of the collection	*/
	Title string `json:"title,omitempty"`
	Links []Link `json:"links"`
	/* indicator about the type of the items in the collection (the default value is 'feature').	*/
	ItemType string `json:"itemType,omitempty"`
}

type FeatureCollection struct {
	NumberReturned int64      `json:"numberReturned,omitempty"`
	TimeStamp      string     `json:"timeStamp,omitempty"`
	Type           string     `json:"type"`
	Features       []*Feature `json:"features"`
	Links          []Link     `json:"links,omitempty"`
	NumberMatched  int64      `json:"numberMatched,omitempty"`
	Crs            string     `json:"crs,omitempty"`
	Offset         int64      `json:"-"`
	Next           bool
}

type Feature struct {
	// overwrite ID in geojson.Feature so strings are also allowed as id
	ID interface{} `json:"id,omitempty"`
	geojson.Feature
	// Added Links in de document
	Links []Link `json:"links,omitempty"`
}

type Extent struct {
	/* The spatial extent of the features in the collection.	*/
	Spatial *Spatial `json:"spatial,omitempty"`
	/* The temporal extent of the features in the collection.	*/
	Temporal *Temporal `json:"temporal,omitempty"`
}

type Spatial struct {
	/* One or more bounding boxes that describe the spatial extent of the dataset.
	   In the Core only a single bounding box is supported. Extensions may support
	   additional areas. If multiple areas are provided, the union of the bounding
	   boxes describes the spatial extent.	*/
	Bbox [][]float64 `json:"bbox,omitempty"`
	/* Coordinate reference system of the coordinates in the spatial extent
	   (property `bbox`). The default reference system is WGS 84 longitude/latitude.
	   In the Core this is the only supported coordinate reference system.
	   Extensions may support additional coordinate reference systems and add
	   additional enum values.	*/
	Crs string `json:"crs,omitempty"`
}

type Temporal struct {
	/* Coordinate reference system of the coordinates in the temporal extent
	   (property `interval`). The default reference system is the Gregorian calendar.
	   In the Core this is the only supported temporal coordinate reference system.
	   Extensions may support additional temporal coordinate reference systems and add
	   additional enum values.	*/
	Trs string `json:"trs,omitempty"`
	/* One or more time intervals that describe the temporal extent of the dataset.
	   The value `null` is supported and indicates an open time interval.
	   In the Core only a single time interval is supported. Extensions may support
	   multiple intervals. If multiple intervals are provided, the union of the
	   intervals describes the temporal extent.	*/
	Interval [][]string `json:"interval,omitempty"`
}

type Link struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Rel      string `json:"rel,omitempty"`
	Length   int64  `json:"length,omitempty"`
	Hreflang string `json:"hreflang,omitempty"`
	Href     string `json:"href"`
}

func JSONMarshaller(i interface{}) []byte {

	if data, err := json.Marshal(i); err == nil {
		return data
	} else {
		log.Fatalf("Could not marshal collections, got error: %v", err)
		// TODO build nice Exception message for client
		return nil
	}
}
