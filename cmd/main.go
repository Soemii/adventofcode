package main

import (
	"github.com/Soemii/AdventOfCode/internal/inbound/web"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode/challenges"
	"github.com/ipfans/fxlogger"
	"github.com/rs/zerolog"
	"net/http"

	"os"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.WithLogger(fxlogger.WithZerolog(zerolog.New(os.Stdout))),
		fx.Invoke(func(server *http.Server) {}),
		fx.Provide(
			web.NewServer,
			fx.Annotate(web.NewServeMux, fx.ParamTags(`group:"route"`)),
			fx.Annotate(web.NewChallengeHandler, fx.As(new(web.Route)), fx.ResultTags(`group:"route"`), fx.ParamTags(`group:"challenge"`)),
			fx.Annotate(web.NewIndexHandler, fx.As(new(web.Route)), fx.ResultTags(`group:"route"`), fx.ParamTags(`group:"challenge"`)),
			AsRoute(web.NewSubmitRoute),

			AsChallenge(challenges.NewChallenge201501),
			AsChallenge(challenges.NewChallenge201502),
			AsChallenge(challenges.NewChallenge201503),
			AsChallenge(challenges.NewChallenge201504),
			AsChallenge(challenges.NewChallenge201505),
			AsChallenge(challenges.NewChallenge201506),

			AsChallenge(challenges.NewChallenge202301),
			AsChallenge(challenges.NewChallenge202302),
			AsChallenge(challenges.NewChallenge202303),
			AsChallenge(challenges.NewChallenge202304),
			AsChallenge(challenges.NewChallenge202305),
			AsChallenge(challenges.NewChallenge202306),
			AsChallenge(challenges.NewChallenge202307),
			AsChallenge(challenges.NewChallenge202308),
			AsChallenge(challenges.NewChallenge202309),
			AsChallenge(challenges.NewChallenge202310),
		),
	).Run()

}

func AsRoute(route interface{}) interface{} {
	return fx.Annotate(route, fx.As(new(web.Route)), fx.ResultTags(`group:"route"`))
}

func AsChallenge(challenge interface{}) interface{} {
	return fx.Annotate(challenge, fx.As(new(adventofcode.Challenge)), fx.ResultTags(`group:"challenge"`))
}
