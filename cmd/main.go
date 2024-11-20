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
		),
	).Run()

}

func AsRoute(route interface{}) interface{} {
	return fx.Annotate(route, fx.As(new(web.Route)), fx.ResultTags(`group:"route"`))
}

func AsChallenge(challenge interface{}) interface{} {
	return fx.Annotate(challenge, fx.As(new(adventofcode.Challenge)), fx.ResultTags(`group:"challenge"`))
}

//func registerAllChallenges() map[string]adventofcode.Challenge {
//	challengeMap := make(map[string]adventofcode.Challenge)
//	//2015
//	addChallenge(challengeMap, challenges2.Challenge20150101{})
//	addChallenge(challengeMap, challenges2.Challenge20150102{})
//
//	addChallenge(challengeMap, challenges2.Challenge20150201{})
//	addChallenge(challengeMap, challenges2.Challenge20150202{})
//
//	addChallenge(challengeMap, challenges2.Challenge20150301{})
//	addChallenge(challengeMap, challenges2.Challenge20150302{})
//
//	addChallenge(challengeMap, challenges2.Challenge20150401{})
//	addChallenge(challengeMap, challenges2.Challenge20150402{})
//
//	addChallenge(challengeMap, challenges2.Challenge20150501{})
//	addChallenge(challengeMap, challenges2.Challenge20150502{})
//
//	//2023
//	addChallenge(challengeMap, challenges2.Challenge20230101{})
//	addChallenge(challengeMap, challenges2.Challenge20230102{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230201{})
//	addChallenge(challengeMap, challenges2.Challenge20230202{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230301{})
//	addChallenge(challengeMap, challenges2.Challenge20230302{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230401{})
//	addChallenge(challengeMap, challenges2.Challenge20230402{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230501{})
//	addChallenge(challengeMap, challenges2.Challenge20230502{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230601{})
//	addChallenge(challengeMap, challenges2.Challenge20230602{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230701{})
//	addChallenge(challengeMap, challenges2.Challenge20230702{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230801{})
//	addChallenge(challengeMap, challenges2.Challenge20230802{})
//
//	addChallenge(challengeMap, challenges2.Challenge20230901{})
//	addChallenge(challengeMap, challenges2.Challenge20230902{})
//
//	addChallenge(challengeMap, challenges2.Challenge20231001{})
//	addChallenge(challengeMap, challenges2.Challenge20231002{})
//
//	return challengeMap
//}
