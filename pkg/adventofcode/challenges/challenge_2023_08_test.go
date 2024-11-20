package challenges

import "testing"

func TestChallenge202308(t *testing.T) {
	ChallengeTest(t, NewChallenge202308(), TestData{
		Expect:    "19631",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "21003205388413",
		NotExpect: nil,
		Finished:  true,
	})
}
