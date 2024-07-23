package middleware

import (
	"log"
	"net/http"
)

// Convinient type for "a middleware is just a handler calling another handler"
type Middleware func(http.Handler) http.Handler

// Compose chains Middlewares in order of appearance.
//
//	// Results in callstack: m1, m2, m3, handler
//	Compose(m1, m2, m3)(handler)
func Compose(chain ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		for i := len(chain) - 1; i >= 0; i-- {
			handler = chain[i](handler)
		}
		return handler
	}
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth Middleware (not implemented)")
		next.ServeHTTP(w, r)
	})
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("New " + r.Method + " request to " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Metrics Middleware (not implemented)")
		next.ServeHTTP(w, r)
	})
}
