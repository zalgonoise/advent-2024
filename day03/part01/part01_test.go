package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2024/day03"
	"github.com/zalgonoise/advent-2024/is"
)

const exampleInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestPart01_Regexp(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "OK/Snippet",
			input: `mul(2,4)`,
			wants: 8,
		},
		{
			name:  "OK/Snippet",
			input: `mul(5,5)`,
			wants: 25,
		},
		{
			name:  "OK/Snippet",
			input: `mul(11,8)`,
			wants: 88,
		},
		{
			name:  "OK/Snippet",
			input: `mul(8,5)`,
			wants: 40,
		},
		{
			name:  "Fail/Snippet",
			input: `mul[3,7]`,
			wants: 0,
		},
		{
			name:  "Fail/Snippet",
			input: `mul(32,64]`,
			wants: 0,
		},
		{
			name:  "Example",
			input: exampleInput,
			wants: 161,
		},
		{
			name:  "Challenge",
			input: day03.Input,
			wants: 165225049,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			muls, err := ParseRegexp(testcase.input)
			is.NoError(t, err)

			result := Sum(muls)

			is.Equal(t, testcase.wants, result)
		})
	}
}

func TestPart01(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "OK/Snippet",
			input: `mul(2,4)`,
			wants: 8,
		},
		{
			name:  "OK/Snippet",
			input: `mul(5,5)`,
			wants: 25,
		},
		{
			name:  "OK/Snippet",
			input: `mul(11,8)`,
			wants: 88,
		},
		{
			name:  "OK/Snippet",
			input: `mul(8,5)`,
			wants: 40,
		},
		{
			name:  "Fail/Snippet",
			input: `mul[3,7]`,
			wants: 0,
		},
		{
			name:  "Fail/Snippet",
			input: `mul(32,64]`,
			wants: 0,
		},
		{
			name:  "Example",
			input: exampleInput,
			wants: 161,
		},
		{
			name:  "Challenge",
			input: day03.Input,
			wants: 164352721,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			muls := Parse(testcase.input)
			result := Sum(muls)

			is.Equal(t, testcase.wants, result)
		})
	}
}
