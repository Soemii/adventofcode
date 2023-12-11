package main

import (
	"flag"
	"fmt"
	"github.com/Soemii/AdventOfCode/internal/challenges"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"log"
	"os"
)

func main() {
	year := flag.Int("year", 2023, "")
	day := flag.Int("day", 1, "")
	challengeId := flag.Int("challenge", 1, "")
	flag.Parse()
	challengeMap := registerAllChallenges()
	challenge, ok := challengeMap[ChallengeParamsTpString(*year, *day, *challengeId)]
	if !ok {
		log.Fatalln("Challenge is not registered")
		return
	}
	file, readErr := os.ReadFile(fmt.Sprintf("./rawinput/%v/day%v.txt", *year, *day))
	if readErr != nil {
		log.Fatalln(readErr)
	}
	executeErr := challenge.Execute(string(file))
	if executeErr != nil {
		log.Fatalln(executeErr)
		return
	}
}

func registerAllChallenges() map[string]adventofcode.Challenge {
	challengeMap := make(map[string]adventofcode.Challenge)
	//2015
	addChallenge(challengeMap, challenges.Challenge20150101{})
	addChallenge(challengeMap, challenges.Challenge20150102{})

	addChallenge(challengeMap, challenges.Challenge20150201{})
	addChallenge(challengeMap, challenges.Challenge20150202{})

	addChallenge(challengeMap, challenges.Challenge20150301{})
	addChallenge(challengeMap, challenges.Challenge20150302{})

	addChallenge(challengeMap, challenges.Challenge20150401{})
	addChallenge(challengeMap, challenges.Challenge20150402{})

	addChallenge(challengeMap, challenges.Challenge20150501{})
	addChallenge(challengeMap, challenges.Challenge20150502{})

	//2023
	addChallenge(challengeMap, challenges.Challenge20230101{})
	addChallenge(challengeMap, challenges.Challenge20230102{})

	addChallenge(challengeMap, challenges.Challenge20230201{})
	addChallenge(challengeMap, challenges.Challenge20230202{})

	addChallenge(challengeMap, challenges.Challenge20230301{})
	addChallenge(challengeMap, challenges.Challenge20230302{})

	addChallenge(challengeMap, challenges.Challenge20230401{})
	addChallenge(challengeMap, challenges.Challenge20230402{})

	addChallenge(challengeMap, challenges.Challenge20230501{})
	addChallenge(challengeMap, challenges.Challenge20230502{})

	addChallenge(challengeMap, challenges.Challenge20230601{})
	addChallenge(challengeMap, challenges.Challenge20230602{})

	addChallenge(challengeMap, challenges.Challenge20230701{})
	addChallenge(challengeMap, challenges.Challenge20230702{})

	addChallenge(challengeMap, challenges.Challenge20230801{})
	addChallenge(challengeMap, challenges.Challenge20230802{})

	addChallenge(challengeMap, challenges.Challenge20230901{})
	addChallenge(challengeMap, challenges.Challenge20230902{})

	addChallenge(challengeMap, challenges.Challenge20231001{})
	addChallenge(challengeMap, challenges.Challenge20231002{})

	return challengeMap
}

func addChallenge(challenges map[string]adventofcode.Challenge, challenge adventofcode.Challenge) {
	_, ok := challenges[ChallengeToString(challenge)]
	if ok {
		log.Printf("Challenge %v already exists", ChallengeToString(challenge))
		return
	}
	challenges[ChallengeToString(challenge)] = challenge
}

func ChallengeToString(challenge adventofcode.Challenge) string {
	return ChallengeParamsTpString(challenge.GetYear(), challenge.GetDay(), challenge.GetChallenge())
}

func ChallengeParamsTpString(year, day, challenge int) string {
	return fmt.Sprintf("%v:%v:%v", year, day, challenge)
}
