package challenges

import "testing"

func TestChallenge202307(t *testing.T) {
	ChallengeTest(t, NewChallenge202307(), TestData{
		Expect:    "247961593",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "248750699",
		NotExpect: nil,
		Finished:  true,
	})
}
