package main

import (
	"net/http"
)

type Middleware []http.Handler

func (m *Middleware) Add(handler http.Handler) {
	*m = append(*m, handler)
}

func (m Middleware) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	// proccess middle ware
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	written bool
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	return &MiddlewareResponseWriter {
		ResponseWriter: w,
	}
}

func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (m Middleware) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	// wrap the supplied ResponseWriter
	mw := NewMiddlewareResponseWriter(w)
	for _, handler := range m {
		// call the handler with our middleresponsewriter
		handler.ServeHTTP(mw, r)
		if mw.written {
			return
		}
	}
	http.NotFound(w, r)
}


