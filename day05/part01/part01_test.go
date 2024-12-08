package part01

import (
	"github.com/zalgonoise/advent-2024/day05"
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
			input: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
			wants: 143,
		},
		{
			name:  "Challenge",
			input: day05.Input,
			wants: 0,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			priorities, instructions, err := Parse(testcase.input)
			is.NoError(t, err)

			filtered := Filter(priorities, instructions)

			is.Equal(t, testcase.wants, SumMiddleValue(filtered))
		})
	}
}
