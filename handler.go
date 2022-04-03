package urlshortener

import (
	"net/http"
  "fmt"
  "gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
    path := r.URL.Path
    elem, ok := pathsToUrls[path]
    if(ok){
      fmt.Println(elem)
      http.Redirect(w, r, elem, http.StatusFound)
      return
    }
    fallback.ServeHTTP(w, r)
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
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type pathUrl struct{
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse the yaml.
	pathUrls, err := parseYaml(yml)
	if(err != nil){
		return nil, err
	}
	// Convert yaml to map (to use map handler)
	pathUrlMap := yamlToMap(pathUrls)
	// Return the mapp handler
	return MapHandler(pathUrlMap, fallback), nil
}

func parseYaml(yamlData []byte)([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlData, &pathUrls)
	if(err != nil){
		return nil, err
	}
	return pathUrls, nil
}

func yamlToMap(pathUrls []pathUrl) (map[string]string){
	pathUrlMap := make(map[string]string)
	for _, p := range pathUrls{
		pathUrlMap[p.Path]=p.URL
	}
	return pathUrlMap
}
