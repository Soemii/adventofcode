package challenges

import "testing"

func TestChallenge20150104(t *testing.T) {
	ChallengeTest(t, NewChallenge201504(), TestData{
		Expect:    "346386",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "9958218",
		NotExpect: nil,
		Finished:  true,
	})
}
