package web

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
	"strconv"
)

func NewChallengeHandler(challenges []adventofcode.Challenge) Route {
	return challengeHandler{
		challenges,
	}
}

type challengeHandler struct {
	challenges []adventofcode.Challenge
}

func (c challengeHandler) FindChallengeByDayAndYear(day int, year int) adventofcode.Challenge {
	for _, challenge := range c.challenges {
		if challenge.GetYear() == year && challenge.GetDay() == day {
			return challenge
		}
	}
	return nil
}

func (c challengeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	temp := template.Must(template.ParseFS(uiFiles, "ui/*.html"))

	year, err := strconv.Atoi(request.PathValue("year"))
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to convert year")
		err = temp.ExecuteTemplate(writer, "error.html", err.Error())
		if err != nil {
			log.Error().Err(err).Caller().Msg("Failed to render error.html")
		}
		return
	}
	day, err := strconv.Atoi(request.PathValue("day"))
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to convert day")
		err = temp.ExecuteTemplate(writer, "error.html", err.Error())
		if err != nil {
			log.Error().Err(err).Caller().Msg("Failed to render error.html")
		}
		return
	}
	challenge := c.FindChallengeByDayAndYear(day, year)
	if challenge == nil {
		log.Error().Msgf("Challenge does not exists")
		err = temp.ExecuteTemplate(writer, "error.html", "Challenge does not exists")
		if err != nil {
			log.Error().Err(err).Caller().Msg("Failed to render error.html")
		}
		return
	}

	err = request.ParseForm()
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to parse form")
		err = temp.ExecuteTemplate(writer, "error.html", err.Error())
		if err != nil {
			log.Error().Err(err).Caller().Msg("Failed to render error.html")
		}
		return
	}
	input := request.FormValue("input")
	first, err := challenge.ExecuteFirst(input)
	if err != nil {
		err = temp.ExecuteTemplate(writer, "error.html", err.Error())
		if err != nil {
			log.Error().Err(err).Caller().Msg("Failed to render error.html")
		}
		return
	}

	second, err := challenge.ExecuteSecond(input)
	if err != nil {
		err = temp.ExecuteTemplate(writer, "error.html", err.Error())
		if err != nil {
			log.Error().Err(err).Caller().Msg("Failed to render error.html")
		}
		return
	}

	err = temp.ExecuteTemplate(writer, "result.html", struct {
		First  string
		Second string
	}{
		First:  first,
		Second: second,
	})
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to render result.html")
	}
	return
}

func (challengeHandler) Pattern() string {
	return "POST /challenge/{year}/{day}"
}
