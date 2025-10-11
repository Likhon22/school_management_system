package middlewares

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func (mw *Middleware) Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check of the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		gz := gzip.NewWriter(w)
		defer gz.Close()
		//wrap the ResponseWriter
		newWriter := &gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gz,
		}
		next.ServeHTTP(newWriter, r)

	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer    *gzip.Writer
	headerSet bool
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	if !g.headerSet {
		g.Header().Set("Content-Encoding", "gzip")
		g.headerSet = true
	}
	return g.Writer.Write(b)
}
func (g *gzipResponseWriter) WriteHeader(statusCode int) {
	if !g.headerSet {
		g.Header().Set("Content-Encoding", "gzip")
		g.headerSet = true
	}
	g.ResponseWriter.WriteHeader(statusCode)
}
