package middleware

import (
	"fmt"
	"net/http"
)

// Convinient type for: "func (next http.Handler) http.Handler"
type Middleware func(http.Handler) http.Handler

// Compose chains Middlewares in order of appearance.
//
//	// Results in callstack: m3, m2, m1, handler
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
		fmt.Print("\nAuth Middleware")
		next.ServeHTTP(w, r)
	})
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("\nLogging Middleware")
		next.ServeHTTP(w, r)
	})
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("\nMetrics Middleware")
		next.ServeHTTP(w, r)
	})
}
