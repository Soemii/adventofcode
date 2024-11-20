package web

import (
	"errors"
	"github.com/rs/zerolog/log"

	"go.uber.org/fx"
	"net/http"
)

func NewServer(lc fx.Lifecycle, handler http.Handler) *http.Server {
	server := &http.Server{
		Handler: handler,
		Addr:    ":9090",
	}
	start := func() {
		go func() {
			//log.Info().Msgf("Listen on %s", server.Addr)
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Error().Err(err).Caller()
			}
		}()
	}
	stop := func() error {
		return server.Close()
	}
	lc.Append(fx.StartStopHook(start, stop))
	return server
}
