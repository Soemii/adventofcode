package challenges

import "testing"

func TestChallenge20150105(t *testing.T) {
	ChallengeTest(t, NewChallenge201505(), TestData{
		Expect:    "238",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "69",
		NotExpect: nil,
		Finished:  true,
	})
}
