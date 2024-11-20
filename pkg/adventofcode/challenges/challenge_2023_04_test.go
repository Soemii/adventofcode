package challenges

import "testing"

func TestChallenge202304(t *testing.T) {
	ChallengeTest(t, NewChallenge202304(), TestData{
		Expect:    "21558",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "10425665",
		NotExpect: nil,
		Finished:  true,
	})
}
