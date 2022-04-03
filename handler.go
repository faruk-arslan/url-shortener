package urlshortener

import (
	"net/http"
  "fmt"
  // "gopkg.in/yaml.v3"
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
type StructPath struct {
	Path string `yaml:"path"`
}

type StructUrl struct {
	// Embedded structs are not treated as embedded in YAML by default. To do that,
	// add the ",inline" annotation below
	StructPath `yaml:",inline"`
	Url       string `yaml:"url"`
}

// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
//   return func (w http.ResponseWriter, r *http.Request){
//     var url StructUrl
//     err := yaml.Unmarshal([]byte(yml), &url)
//     if err != nil {
//       fallback.ServeHTTP(w, r)
//     }
//     fmt.Println(url.Path)
//   	fmt.Println(url.Url)
//     http.Redirect(w, r, url.Url, http.StatusFound)
//     return
//   }
//
//
// }
