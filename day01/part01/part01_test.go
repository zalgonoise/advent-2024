package part01

import (
	"github.com/zalgonoise/advent-2024/day01"
	"testing"

	"github.com/zalgonoise/advent-2024/is"
)

const exampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart01(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Example",
			input: exampleInput,
			wants: 11,
		},
		{
			name:  "Challenge",
			input: day01.Input,
			wants: 1222801,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			values, err := ParseTable(testcase.input)
			is.NoError(t, err)

			diff := GetDifference(values)
			res := Sum(diff)

			is.Equal(t, testcase.wants, res)
		})
	}
}
