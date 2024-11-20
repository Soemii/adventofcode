package challenges

import "testing"

func TestChallenge202303(t *testing.T) {
	ChallengeTest(t, NewChallenge202303(), TestData{
		Expect:    "520019",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "75519888",
		NotExpect: nil,
		Finished:  true,
	})
}
