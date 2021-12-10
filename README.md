# GOGCAPI

## Example

Run from `/example`

```go
go run .
```

## Assumptions

1. Collections are a core mechanic therefor they are preprepared
1. Collections are generated/driven by configuration. Therefor quite static which
   should reflect in the code
1. (Collections != Collection) != (FeatureCollections != Feature) Thinks might look
   the same, but aren't
1. Feature data that is moved around in it's own structs based on the go-spatial
   geom package
1. html and the json/geojson output are part of the core mechanic of this package
1. the html a template/file served that call the json for populating the page
   with data
