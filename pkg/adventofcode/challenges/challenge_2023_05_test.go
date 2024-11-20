package challenges

import "testing"

func TestChallenge202305(t *testing.T) {
	ChallengeTest(t, NewChallenge202305(), TestData{
		Expect:    "346433842",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "60294664",
		NotExpect: nil,
		Finished:  true,
	})
}
