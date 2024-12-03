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
			wants:  711,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			values, err := ParseTable(testcase.input)
			is.NoError(t, err)

			filtered := Filter2(values, testcase.thresh)

			is.Equal(t, testcase.wants, len(filtered))
		})
	}
}

func TestOrder(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input []int
		order int
		skip  bool
	}{
		{
			name:  "Ascending",
			input: []int{1, 2, 3, 4, 5},
			order: 1,
			skip:  false,
		},
		{
			name:  "Descending",
			input: []int{5, 4, 3, 2, 1},
			order: -1,
			skip:  false,
		},
		{
			name:  "Descending/WithSkip",
			input: []int{5, 4, 8, 2, 1},
			order: -1,
			skip:  false,
		},
		{
			name:  "Descending/SwapOnStart",
			input: []int{3, 4, 3, 2, 1},
			order: -1,
			skip:  false,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			order, isUnsafe := getOrder(testcase.input)

			is.Equal(t, testcase.skip, isUnsafe)
			is.Equal(t, testcase.order, order)
		})
	}
}

func TestIsSafe_Ascending(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input []int
		wants bool
	}{
		{
			name:  "OK/NoInterval",
			input: []int{1, 2, 3, 4, 5},
			wants: true,
		},
		{
			name:  "OK/MaxInterval",
			input: []int{1, 4, 7, 10, 13},
			wants: true,
		},
		{
			name:  "OK/WithGap/OnEnd",
			input: []int{1, 4, 7, 10, 9},
			wants: true,
		},
		{
			name:  "OK/WithGap/OnMiddle",
			input: []int{1, 4, 7, 6, 8},
			wants: true,
		},
		{
			name:  "OK/WithGap/OnStart",
			input: []int{3, 2, 3, 4, 5},
			wants: true,
		},
		{
			name:  "OK/WithGap/OnEnd/Equal",
			input: []int{1, 4, 7, 10, 10},
			wants: true,
		},
		{
			name:  "OK/WithGap/OnMiddle/Equal",
			input: []int{1, 4, 7, 7, 8},
			wants: true,
		},
		{
			name:  "OK/WithGap/OnStart/Equal",
			input: []int{2, 2, 3, 4, 5},
			wants: true,
		},
		{
			name:  "OK/IntervalOverflow/Once/Start",
			input: []int{1, 6, 7, 8, 9},
			wants: true,
		},
		{
			name:  "OK/IntervalOverflow/Once/Middle",
			input: []int{5, 6, 13, 8, 9},
			wants: true,
		},
		{
			name:  "OK/IntervalOverflow/Once/End",
			input: []int{5, 6, 7, 8, 15},
			wants: true,
		},
		{
			name:  "Fail/IntervalOverflow/Start",
			input: []int{1, 7, 13, 14, 15},
			wants: false,
		},
		{
			name:  "Fail/IntervalOverflow/Middle",
			input: []int{1, 2, 7, 8, 9},
			wants: false,
		},
		{
			name:  "Fail/IntervalOverflow/End",
			input: []int{5, 6, 7, 13, 21},
			wants: false,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			is.Equal(t, testcase.wants, isSafe(testcase.input, 3))
		})
	}
}
