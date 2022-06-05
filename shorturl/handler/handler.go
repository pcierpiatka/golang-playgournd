package handler

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"strings"
)

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
func FileHandler(fileName string, fallback http.Handler) (http.Handler, error) {

	extension := strings.Split(fileName, ".")
	fileType := extension[1]

	paths, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if fileType == "json" {
		return jsonHandler(paths, fallback)
	} else if fileType == "yml" {
		return yamlHandler(paths, fallback)
	} else {
		fmt.Printf("Unsupported file type %s", fileType)
		os.Exit(1)
		return nil, nil
	}
}

func yamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)

	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(pathUrls), fallback), nil
}

// TODO introduce interface for handler
func jsonHandler(jsonBytes []byte, fallback http.Handler) (http.Handler, error) {

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

type Ble interface {
	handle(bytes []byte, fallback http.Handler) (http.HandlerFunc, error)
}
