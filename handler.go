package urlshort

import (
	"log"
	"net/http"

	Yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		url, prs := pathsToUrls[p]
		if !prs {
			fallback.ServeHTTP(w, r)
			return
		}
		http.RedirectHandler(url, http.StatusPermanentRedirect).ServeHTTP(w, r)
	})
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var paths redirects
	err := Yaml.Unmarshal(yml, &paths)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
		return nil, err
	}
	pathMap := make(map[string]string)
	for _, el := range paths {
		pathMap[el.Path] = el.Url
	}
	return MapHandler(pathMap, fallback), nil
}

type redirect struct {
	Path string
	Url  string
}

type redirects []redirect
