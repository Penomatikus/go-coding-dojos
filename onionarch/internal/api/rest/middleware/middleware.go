package middleware

import (
	"fmt"
	"net/http"
)

// Convinient type for: "func (next http.Handler) http.Handler"
type Middleware func(http.Handler) http.Handler

// Chain chains middlewares (handlers calling handlers)
type Chain struct {
	middlewares []Middleware
}

// Provides a new Chain with an option to keep the order of appearance
func NewChain(preverseOrder bool, chain ...Middleware) Chain {
	c := Chain{middlewares: chain}
	if preverseOrder {
		c.preserveOrder()
	}

	return c
}

func (c *Chain) preserveOrder() {
	reverseChain := make([]Middleware, 0, len(c.middlewares))
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		reverseChain = append(reverseChain, c.middlewares[i])
	}
	c.middlewares = reverseChain
}

// Commit wraps each middleware with the next one in the chain where handler is the last commited one.
// Examples of struct field tags and their meanings:
//
//	// Results in callstack: m3, m2, m1, handler
//	c := NewMiddlewareChain(false, m1, m2, m3)
//	c.Commit(handler)
//
//	// Results in callstack: m1, m2, m3, handler
//	c := NewMiddlewareChain(true, m1, m2, m3)
//	c.Commit(handler)
func (c *Chain) Commit(handler http.Handler) http.Handler {
	for i := range c.middlewares {
		handler = c.middlewares[i](handler)
	}
	return handler
}

// Compose chains Middlewares in order of appearance.
//
//	// Results in callstack: m3, m2, m1, handler
//	Compose(m1, m2, m3)(handler)
//
// Note: Its the function only approach of Chain{}
func Compose(chain ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		for _, c := range chain {
			handler = c(handler)
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
