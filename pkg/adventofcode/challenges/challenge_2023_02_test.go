package challenges

import "testing"

func TestChallenge202302(t *testing.T) {
	ChallengeTest(t, NewChallenge202302(), TestData{
		Expect:    "2369",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "66363",
		NotExpect: nil,
		Finished:  true,
	})
}
