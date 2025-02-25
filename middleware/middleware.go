package middleware

import (
	l "cds-graphql-poc/graph/loaders"
	c "cds-graphql-poc/middleware/contextcache"
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

// Cache Middleware injects the cache into the context
func Cache(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add all caches to the context
		r = r.WithContext(c.WithCompanyCache(r.Context()))
		r = r.WithContext(c.WithInvestmentCache(r.Context()))
		r = r.WithContext(c.WithSummaryKPICache(r.Context()))

		next.ServeHTTP(w, r)
	})
}
