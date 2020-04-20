package handlers

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

//YAMLStruct ..
type YAMLStruct struct {
	PATH string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler ...
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLParser ...
func YAMLParser(data []byte, fallback http.Handler) (http.Handler, error) {
	var paths []YAMLStruct
	err := yaml.Unmarshal(data, &paths)
	if err != nil {
		return nil, err
	}
	pathMap := make(map[string]string)
	for _, obj := range paths {
		pathMap[obj.PATH] = obj.URL
	}
	return MapHandler(pathMap, fallback), nil
}
