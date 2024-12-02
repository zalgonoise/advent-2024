package part01

import (
	"errors"
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

func Filter(set [][]int, thresh int) []int {
	indices := make([]int, 0, len(set))

	for i := range set {
		switch {
		case len(set[i]) < 2:
			continue
		case set[i][0] > set[i][1]:
			if filterDecreasing(set[i], thresh) {
				indices = append(indices, i)
			}
		case set[i][0] < set[i][1]:
			if filterIncreasing(set[i], thresh) {
				indices = append(indices, i)
			}
		default:
			continue
		}
	}

	return indices
}

func filterIncreasing(set []int, thresh int) bool {
	for i := 1; i < len(set); i++ {
		if set[i] <= set[i-1] {
			return false
		}

		if set[i]-set[i-1] > thresh {
			return false
		}
	}

	return true
}

func filterDecreasing(set []int, thresh int) bool {
	for i := 1; i < len(set); i++ {
		if set[i] >= set[i-1] {
			return false
		}

		if set[i-1]-set[i] > thresh {
			return false
		}
	}

	return true
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
