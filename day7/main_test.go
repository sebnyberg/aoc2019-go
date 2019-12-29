package day7_test

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/aoc2019/day7"
	"github.com/stretchr/testify/require"
)

func getThrust(program []int, phases []int) (int, error) {
	var (
		output int
		err    error
	)
	for _, p := range phases {
		output, err = day7.RunProgram(program, []int{p, output}, true)
		if err != nil {
			return 0, err
		}
	}

	return output, nil
}

func Test_Day7(t *testing.T) {
	tcs := []struct {
		program        []int
		phases         []int
		shouldErr      bool
		expectedOutput int
	}{
		{
			program:        []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			phases:         []int{4, 3, 2, 1, 0},
			shouldErr:      false,
			expectedOutput: 43210,
		},
		{
			program:        []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			phases:         []int{0, 1, 2, 3, 4},
			shouldErr:      false,
			expectedOutput: 54321,
		},
		{
			program: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			phases:         []int{1, 0, 4, 3, 2},
			shouldErr:      false,
			expectedOutput: 65210,
		},
	}

	for idx, tc := range tcs {
		t.Run(fmt.Sprintf("test_%v", idx), func(t *testing.T) {
			output, err := getThrust(tc.program, tc.phases)
			require.Nil(t, err)
			require.Equal(t, tc.expectedOutput, output)
		})
	}

}
