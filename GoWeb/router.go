package main

import (
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]map[string]http.HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {

	m, ok := r.handlers[method]
	if !ok {
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}

	m[pattern] = h

}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	for pattern, handler := range r.handlers[req.Method] {
		if ok, _ := match(pattern, req.URL.Path); ok {
			handler(w, req)
			return
		}
	}

	http.NotFound(w, req)
	return
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {

		switch {
		case patterns[i] == paths[i]:
		case len(patterns[i]) > 0 && string(patterns[i][0]) == ":":
			params[patterns[i][1:]] = paths[i]
		default:
			return false, nil
		}
	}

	return true, params

}
