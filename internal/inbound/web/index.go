package web

import (
	"fmt"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
)

func NewIndexHandler(challenges []adventofcode.Challenge) Route {
	return indexHandler{
		challenges: challenges,
	}
}

type indexHandler struct {
	challenges []adventofcode.Challenge
}

func (i indexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	challengeMap := make(map[string]string)
	for _, challenge := range i.challenges {
		challengeMap[fmt.Sprintf("/challenge/%d/%d", challenge.GetYear(), challenge.GetDay())] = fmt.Sprintf("%d - %d", challenge.GetYear(), challenge.GetDay())
	}
	err := template.Must(template.ParseFS(uiFiles, "ui/*.html")).ExecuteTemplate(writer, "index.html", challengeMap)
	if err != nil {
		log.Error().Err(err)
	}
}

func (indexHandler) Pattern() string {
	return "GET /"
}
