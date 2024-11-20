package challenges

import "testing"

func TestChallenge20150106(t *testing.T) {
	ChallengeTest(t, NewChallenge201506(), TestData{
		Expect:    "569999",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "17836115",
		NotExpect: nil,
		Finished:  true,
	})
}
