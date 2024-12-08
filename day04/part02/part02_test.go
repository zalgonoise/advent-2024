package part02

import (
	"github.com/zalgonoise/advent-2024/day04"
	"github.com/zalgonoise/advent-2024/is"
	"testing"
)

func TestPart01(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name: "Example",
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			wants: 9,
		},
		{
			name: "Small",
			input: `MAS
SAM
MMX`,
			wants: 0,
		},
		{
			name:  "Challenge",
			input: day04.Input,
			wants: 0,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			grid := Parse(testcase.input)

			is.Equal(t, testcase.wants, len(FindAll(grid)))
		})
	}
}
