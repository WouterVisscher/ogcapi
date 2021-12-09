package ogcapi

// GeoJSON View

import (
	"encoding/json"
	"log"
)

func JSONMarshaller(i interface{}) []byte {

	if data, err := json.Marshal(i); err == nil {
		return data
	} else {
		log.Fatalf("Could not marshal collections, got error: %v", err)
		// TODO build nice Exception message for client
		return nil
	}
}
