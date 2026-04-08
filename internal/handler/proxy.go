package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// NewUpstreamProxy returns a Gin handler that reverse-proxies to the given origin
// (e.g. https://thesimpsonsapi.com), preserving path and query string.
func NewUpstreamProxy(origin *url.URL) gin.HandlerFunc {
	rp := httputil.NewSingleHostReverseProxy(origin)
	// Incoming requests keep Host from the client (e.g. localhost). The upstream
	// CDN expects Host to match thesimpsonsapi.com; otherwise it may reject with 403.
	director := rp.Director
	rp.Director = func(req *http.Request) {
		director(req)
		req.Host = origin.Host
	}
	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(`{"error":"upstream unavailable"}`))
	}
	return func(c *gin.Context) {
		rp.ServeHTTP(c.Writer, c.Request)
	}
}
