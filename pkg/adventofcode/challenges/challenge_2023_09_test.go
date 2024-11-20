package challenges

import "testing"

func TestChallenge202309(t *testing.T) {
	ChallengeTest(t, NewChallenge202309(), TestData{
		Expect:    "1877825184",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "1108",
		NotExpect: nil,
		Finished:  true,
	})
}
