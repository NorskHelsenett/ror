package httpserver

import (
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func MiddlewareLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rlog.Debugc(r.Context(), "Request", rlog.String("method", r.Method), rlog.String("url", r.URL.String()))
		f(w, r)
	}
}

func MiddlewareJsonContentType(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}
