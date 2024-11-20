package challenges

import "testing"

func TestChallenge202310(t *testing.T) {
	ChallengeTest(t, NewChallenge202310(), TestData{
		Expect:    "6820",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "1108",
		NotExpect: nil,
		Finished:  false,
	})
}
