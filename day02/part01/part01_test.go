package part01

import (
	"github.com/zalgonoise/advent-2024/day02"
	"testing"

	"github.com/zalgonoise/advent-2024/is"
)

const exampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart01(t *testing.T) {
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
			wants:  2,
		},
		{
			name:   "Challenge",
			input:  day02.Input,
			thresh: 3,
			wants:  585,
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
