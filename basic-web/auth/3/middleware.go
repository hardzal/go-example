package main

import "net/http"

// USERNAME constants
const USERNAME = "batman"

// PASSWORD constants
const PASSWORD = "secret"

// CustomeMux struct
type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

// RegisterMiddleware func
func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}

// MiddlewareAuth func
func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("something went wrong"))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte("wrong username/password"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// MiddlewareAllowOnlyGET func
func MiddlewareAllowOnlyGET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
