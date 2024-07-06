package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExampleCompose() {
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("1, ")
			next.ServeHTTP(w, r)
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("2, ")
			next.ServeHTTP(w, r)
		})
	}

	middleware3 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("3, ")
			next.ServeHTTP(w, r)
		})
	}

	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("4;")
	})

	req := httptest.NewRequest("POST", "/", bytes.NewReader([]byte{}))
	Compose(middleware1, middleware2, middleware3)(finalHandler).ServeHTTP(nil, req)
	// Output: 1, 2, 3, 4;
}
