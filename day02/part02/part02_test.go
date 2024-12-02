package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2024/day02"
	"github.com/zalgonoise/advent-2024/is"
)

const exampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart02(t *testing.T) {
	for _, testcase := range []struct {
		name   string
		input  string
		thresh int
		wants  int
	}{
		{
			name:   "Example",
			input:  exampleInput,
			thresh: 3,
			wants:  4,
		},
		{
			name:   "Challenge",
			input:  day02.Input,
			thresh: 3,
			wants:  618,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			values, err := ParseTable(testcase.input)
			is.NoError(t, err)

			filtered := Filter(values, testcase.thresh)

			is.Equal(t, testcase.wants, len(filtered))
		})
	}
}

func TestFilter(t *testing.T) {
	for _, testcase := range []struct {
		name   string
		input  []int
		thresh int
		wants  bool
	}{
		{
			name:   "Success/Ascending",
			input:  []int{1, 2, 3, 4, 5},
			thresh: 3,
			wants:  true,
		},
		{
			name:   "Success/Descending",
			input:  []int{5, 4, 3, 2, 1},
			thresh: 3,
			wants:  true,
		},
		{
			name:   "Success/Ascending/WithSwap/Index0",
			input:  []int{3, 2, 3, 4, 5},
			thresh: 3,
			wants:  true,
		},
		{
			name:   "Success/Ascending/WithSwap/Index1",
			input:  []int{1, 0, 3, 4, 5},
			thresh: 3,
			wants:  true,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			n := Filter([][]int{testcase.input}, testcase.thresh)
			if !testcase.wants {
				is.Equal(t, 0, len(n))

				return
			}

			is.Equal(t, 1, len(n))
		})
	}
}

func TestOrder(t *testing.T) {
	for _, testcase := range []struct {
		name     string
		input    []int
		order    int
		skip     bool
		wantSets int
	}{
		{
			name:     "Ascending",
			input:    []int{1, 2, 3, 4, 5},
			order:    1,
			skip:     false,
			wantSets: 1,
		},
		{
			name:     "Descending",
			input:    []int{5, 4, 3, 2, 1},
			order:    -1,
			skip:     false,
			wantSets: 1,
		},
		{
			name:     "Descending/WithSkip",
			input:    []int{5, 4, 8, 2, 1},
			order:    -1,
			skip:     false,
			wantSets: 1,
		},
		{
			name:     "Descending/SwapOnStart",
			input:    []int{3, 4, 3, 2, 1},
			order:    -1,
			skip:     false,
			wantSets: 1,
		},
		{
			name:     "Descending/SwapOnStart",
			input:    []int{3, 4, 3, 2, 1},
			order:    -1,
			skip:     false,
			wantSets: 1,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			skip, sets := checkOrder(testcase.input, true)

			if skip {
				is.Equal(t, 0, len(sets.items))

				return
			}
		})
	}
}
