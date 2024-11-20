package web

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewSubmitRoute() Route {
	return submitRoute{}
}

type submitRoute struct{}

func (submitRoute) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	path := request.FormValue("path")
	http.Redirect(writer, request, path, http.StatusTemporaryRedirect)
	log.Info().Str("path", path).Msg("Redirect submit")
}

func (submitRoute) Pattern() string {
	return "POST /submit"
}
