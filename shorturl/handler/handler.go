package handler

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return a http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, request, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, request)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /handler-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)

	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(pathUrls), fallback), nil
}

// TODO introduce interface for handler
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.Handler, error) {

	var pathUrls []pathUrl
	err := yaml.Unmarshal(jsonBytes, &pathUrls)

	return bytesHandler(pathUrls, err, fallback)
}

func bytesHandler(pathUrls []pathUrl, err error, fallback http.Handler) (http.Handler, error) {
	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(pathUrls), fallback), nil

}

func buildMap(pathsToUrl []pathUrl) (buildMap map[string]string) {
	buildMap = make(map[string]string)
	for _, path := range pathsToUrl {
		buildMap[path.Path] = path.URL
	}
	return
}

type pathUrl struct {
	Path string
	URL  string
}
