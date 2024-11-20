package web

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewServeMux(routes []Route) (http.Handler, error) {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("path", r.URL.Path).Str("method", r.Method).Msg("New Request")
		mux.ServeHTTP(w, r)
	}), nil
}
