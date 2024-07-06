package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func ExampleCompose() {
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("1, ")
			fmt.Fprint(w, "1, ")
			next.ServeHTTP(w, r)
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("2, ")
			fmt.Fprint(w, "2, ")
			next.ServeHTTP(w, r)
		})
	}

	middleware3 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("3, ")
			fmt.Fprint(w, "3, ")
			next.ServeHTTP(w, r)
		})
	}

	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("4;")
		fmt.Fprint(w, "4;")
	})

	req := httptest.NewRequest("POST", "/", bytes.NewReader([]byte{}))
	rec := httptest.NewRecorder()
	Compose(middleware1, middleware2, middleware3)(finalHandler).ServeHTTP(rec, req)
	data, _ := io.ReadAll(rec.Body)
	fmt.Print(" " + string(data))
	// Output: 1, 2, 3, 4; 1, 2, 3, 4;
}
