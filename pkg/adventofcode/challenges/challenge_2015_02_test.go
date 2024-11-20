package challenges

import "testing"

func TestChallenge20150102(t *testing.T) {
	ChallengeTest(t, NewChallenge201502(), TestData{
		Expect:    "1606483",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "3842356",
		NotExpect: nil,
		Finished:  true,
	})
}
