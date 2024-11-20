package challenges

import "testing"

func TestChallenge20150103(t *testing.T) {
	ChallengeTest(t, NewChallenge201503(), TestData{
		Expect:    "2081",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "2341",
		NotExpect: nil,
		Finished:  true,
	})
}
