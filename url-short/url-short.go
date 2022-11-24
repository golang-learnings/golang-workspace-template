package main

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return a http.HandleFunc (which also implements http.Handler) that will attempt to map any paths (keys in the map) to their corresponding URL (values that each key in the map points to, in the string format).
// If the Path is not provided in the map, then the fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. parse YAML
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	// 2. Convert YAML array into map

	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	// 3. return map handler using map
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
