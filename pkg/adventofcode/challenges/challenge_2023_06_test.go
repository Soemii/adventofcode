package challenges

import "testing"

func TestChallenge202306(t *testing.T) {
	ChallengeTest(t, NewChallenge202306(), TestData{
		Expect:    "1660968",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "26499773",
		NotExpect: nil,
		Finished:  true,
	})
}
