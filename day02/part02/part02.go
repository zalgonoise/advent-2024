package part02

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

var ErrInputEmpty = errors.New("input is empty")

func ParseTable(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")

	if len(lines) == 0 {
		return nil, ErrInputEmpty
	}

	set := make([][]int, 0, len(lines))

	for i := range lines {
		items, err := parseTableItem(lines[i])
		if err != nil {
			return nil, err
		}

		set = append(set, items)
	}

	return set, nil
}

func Filter2(sets [][]int, thresh int) []int {
	indices := make([]int, 0, len(sets))

	for i := range sets {
		if isSafe(sets[i], thresh) {
			indices = append(indices, i)
		}
	}

	return indices
}

func isSafe(set []int, thresh int) bool {
	order, isUnsafe := getOrder(set)
	if isUnsafe {
		return false
	}

	switch order {
	case 1:
		return isAnySafe(breakdownSet(set), thresh)
	case -1:
		rev := make([]int, len(set))
		copy(rev, set)
		slices.Reverse(rev)

		return isAnySafe(breakdownSet(rev), thresh)
	}

	return false
}

func isAnySafe(sets [][]int, thresh int) bool {
	for i := range sets {
		if isSafeAscending(sets[i], thresh) {
			return true
		}
	}

	return false
}

func isSafeAscending(set []int, thresh int) bool {
	for i := 1; i < len(set); i++ {
		if set[i] < set[i-1] {
			return false
		}

		if set[i]-set[i-1] > thresh {
			return false
		}
	}

	return true
}

func breakdownSet(set []int) [][]int {
	sets := make([][]int, 0, len(set)+1)
	sets = append(sets, set)

	for i := range set {
		n := make([]int, 0, len(set)-1)

		n = append(n, set[:i]...)
		n = append(n, set[i+1:]...)

		sets = append(sets, n)
	}

	return sets
}

func getOrder(set []int) (int, bool) {
	var (
		asc  int
		desc int

		order = make([]int, len(set)-1)
	)

	for i := 1; i < len(set); i++ {
		switch {
		case set[i] > set[i-1]:
			order[i-1] = 1
			asc++
		case set[i] < set[i-1]:
			order[i-1] = -1
			desc++
		}
	}

	switch {
	case asc > 1 && desc > 1:
		return 0, true
	case asc > 1:
		return 1, false
	case desc > 1:
		return -1, false
	default:
		return 0, true
	}
}

func parseTableItem(item string) ([]int, error) {
	values := strings.Split(item, " ")

	output := make([]int, 0, len(values))

	for i := range values {
		value, err := strconv.Atoi(values[i])
		if err != nil {
			return nil, err
		}

		output = append(output, value)
	}

	return output, nil
}
