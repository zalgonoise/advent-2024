package part01

import (
	"github.com/zalgonoise/advent-2024/day06"
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
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			wants: 41,
		},
		{
			name:  "Challenge",
			input: day06.Input,
			wants: 0,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			grid := Parse(testcase.input)
			res := Move(grid)

			is.Equal(t, testcase.wants, SumResults(res))
		})
	}
}
