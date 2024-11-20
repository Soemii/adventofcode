package challenges

import (
	"fmt"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"os"
	"strconv"
	"testing"
)

type TestData struct {
	Expect    string
	NotExpect []func(string string) bool
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
		for _, fun := range dataFirst.NotExpect {
			if fun(result) {
				t.Errorf("ExecuteFirst() not expected return value = %v", result)
				return
			}
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
		for _, fun := range dataFirst.NotExpect {
			if fun(result) {
				t.Errorf("ExecuteSecond() not expected return value = %v", result)
				return
			}
		}
		if dataSecond.Expect != "" && dataSecond.Expect != result {
			t.Errorf("ExecuteSecond() unexpected return value = %v", result)
		}
	})
}

func LowerThan(i int) func(string string) bool {
	return func(string string) bool {
		atoi, _ := strconv.Atoi(string)
		return atoi >= i
	}
}

func BiggerThan(i int) func(string string) bool {
	return func(string string) bool {
		atoi, _ := strconv.Atoi(string)
		return atoi <= i
	}
}
