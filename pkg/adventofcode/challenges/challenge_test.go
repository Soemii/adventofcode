package challenges

import (
	"fmt"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"os"
	"slices"
	"testing"
)

type TestData struct {
	Expect    string
	NotExpect []string
	Finished  bool
}

func ChallengeTest(t *testing.T, challenge adventofcode.Challenge, dataFirst TestData, dataSecond TestData) {
	path := fmt.Sprintf("testdata/%d/%d.txt", challenge.GetYear(), challenge.GetDay())
	file, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("Could not read file: %s", path)
	}
	input := string(file)
	t.Run(fmt.Sprintf("%d-%d-First", challenge.GetYear(), challenge.GetDay()), func(t *testing.T) {
		if !dataFirst.Finished {
			return
		}
		result, err := challenge.ExecuteFirst(input)
		if err != nil {
			t.Errorf("ExecuteFirst() error = %v", err)
			return
		}
		if slices.Contains(dataFirst.NotExpect, result) {
			t.Errorf("ExecuteFirst() not expected return value = %v", result)
			return
		}
		if dataFirst.Expect != "" && dataFirst.Expect != result {
			t.Errorf("ExecuteFirst() unexpected return value = %v", result)
			return
		}
	})
	t.Run(fmt.Sprintf("%d-%d-Second", challenge.GetYear(), challenge.GetDay()), func(t *testing.T) {
		if !dataSecond.Finished {
			return
		}
		result, err := challenge.ExecuteSecond(input)
		if err != nil {
			t.Errorf("ExecuteSecond() error = %v", err)
		}
		if slices.Contains(dataSecond.NotExpect, result) {
			t.Errorf("ExecuteSecond() not expected return value = %v", result)
		}
		if dataSecond.Expect != "" && dataSecond.Expect != result {
			t.Errorf("ExecuteSecond() unexpected return value = %v", result)
		}
	})
}
