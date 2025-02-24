package middleware

import (
	l "cds-graphql-poc/graph/loaders"
	"context"
	"net/http"
)

// Loaders Middleware injects data loaders into the context
func Loaders(
	loaders *l.Loaders,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), l.GetLoadersKey(), loaders))
		next.ServeHTTP(w, r)
	})
}
