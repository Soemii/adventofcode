package challenges

import (
	"testing"
)

func TestChallenge20150101(t *testing.T) {
	ChallengeTest(t, NewChallenge201501(), TestData{
		Expect:    "74",
		NotExpect: nil,
		Finished:  true,
	}, TestData{
		Expect:    "1795",
		NotExpect: nil,
		Finished:  true,
	})
}
