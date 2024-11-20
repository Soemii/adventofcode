package challenges

import "testing"

func TestChallenge202301(t *testing.T) {
	ChallengeTest(t, NewChallenge202301(), TestData{
		Expect:    "55971",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "54719",
		NotExpect: nil,
		Finished:  true,
	})
}
