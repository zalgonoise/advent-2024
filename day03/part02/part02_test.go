package part02

import (
	"github.com/zalgonoise/advent-2024/day03"
	"github.com/zalgonoise/advent-2024/is"
	"testing"
)

const exampleInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart01_Regexp(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Example",
			input: exampleInput,
			wants: 48,
		},
		{
			name:  "Challenge",
			input: day03.Input,
			wants: 108830766,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			valid := ParseValid(testcase.input)
			muls, err := ParseRegexp(valid)
			is.NoError(t, err)

			result := Sum(muls)

			is.Equal(t, testcase.wants, result)
		})
	}
}
