package part02

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
		skip, s := checkOrder(set[i], true)

		switch {
		case skip:
			continue
		case s.order < 0:
			if filterDecreasing(set[i], thresh, s.canSkip) {
				indices = append(indices, i)
			}
		case s.order > 0:
			if filterIncreasing(set[i], thresh, s.canSkip) {
				indices = append(indices, i)
			}
		}
	}

	return indices
}

type Set struct {
	canSkip bool
	order   int
	items   []int
}

func checkOrder(set []int, alreadySkipped bool) (skip bool, out Set) {
	var order int

	if len(set) < 2 {
		return true, Set{}
	}

	for i := 1; i < len(set)-1; i++ {
		switch {
		case set[i] > set[i-1]:
			if order < 0 {
				if !alreadySkipped {
					return true, Set{}
				}

				alreadySkipped = false

				prev, cur, _ := cropByIndex(set, i)

				prevSkip, prevOut := checkOrder(prev, false)
				curSkip, curOut := checkOrder(cur, false)

				switch {
				case !prevSkip && !curSkip:
					return true, Set{}
				case !prevSkip:
					return false, prevOut
				case !curSkip:
					return false, curOut
				default:
					return true, Set{}
				}
			}

			order = 1
		case set[i] < set[i-1]:
			if order > 0 {
				if !alreadySkipped {
					return true, Set{}
				}

				alreadySkipped = false

				prev, cur, _ := cropByIndex(set, i)

				prevSkip, prevOut := checkOrder(prev, false)
				curSkip, curOut := checkOrder(cur, false)

				switch {
				case !prevSkip && !curSkip:
					return true, Set{}
				case !prevSkip:
					return false, prevOut
				case !curSkip:
					return false, curOut
				default:
					return true, Set{}
				}
			}

			order = -1
		default:
		}
	}

	return false, Set{canSkip: alreadySkipped, order: order, items: set}
}

func filterIncreasing(set []int, thresh int, canSkip bool) bool {
	for i := 1; i < len(set); i++ {
		switch {
		case set[i] < set[i-1]:
			if !canSkip {
				return false
			}

			if i == 1 {
				switch {
				case set[i] < set[i-1]:
					prev, cur, _ := cropByIndex(set, i)

					if !filterDecreasing(prev, thresh, false) &&
						!filterDecreasing(cur, thresh, false) {
						return false
					}

					continue
				case set[i] > set[i-1]:
					prev, cur, _ := cropByIndex(set, i)

					if !filterIncreasing(prev, thresh, false) &&
						!filterIncreasing(cur, thresh, false) {
						return false
					}

					continue
				default:
					return false
				}
			}
		case set[i]-set[i-1] == 0:
			if !canSkip {
				return false
			}

			prev, cur, next := cropByIndex(set, i)

			if !filterIncreasing(prev, thresh, false) &&
				!filterIncreasing(cur, thresh, false) &&
				!filterIncreasing(next, thresh, false) {
				return false
			}
		case set[i]-set[i-1] > thresh:
			if !canSkip {
				return false
			}

			prev, cur, _ := cropByIndex(set, i)

			if !filterIncreasing(prev, thresh, false) &&
				!filterIncreasing(cur, thresh, false) {
				return false
			}
		}
	}

	return true
}

func filterDecreasing(set []int, thresh int, canSkip bool) bool {
	for i := 1; i < len(set); i++ {
		switch {
		// swapped order
		// [3,2,3] --> ASC:[2,3], DESC[3,2]
		case set[i-1] < set[i]:
			if !canSkip {
				return false
			}

			if i == 1 {
				switch {
				case set[i] < set[i-1]:
					prev, cur, _ := cropByIndex(set, i)

					if !filterDecreasing(prev, thresh, false) &&
						!filterDecreasing(cur, thresh, false) {
						return false
					}

					continue
				case set[i] > set[i-1]:
					prev, cur, _ := cropByIndex(set, i)

					if !filterIncreasing(prev, thresh, false) &&
						!filterIncreasing(cur, thresh, false) {
						return false
					}

					continue
				default:
					return false
				}
			}

		// values are the same
		case set[i-1]-set[i] == 0:
			if !canSkip {
				return false
			}

			prev, cur, next := cropByIndex(set, i)
			if !filterDecreasing(prev, thresh, false) &&
				!filterDecreasing(cur, thresh, false) &&
				!filterDecreasing(next, thresh, false) {
				return false
			}
		// over threshold
		case set[i-1]-set[i] > thresh:
			if !canSkip {
				return false
			}

			prev, cur, _ := cropByIndex(set, i)

			if !filterDecreasing(prev, thresh, false) &&
				!filterDecreasing(cur, thresh, false) {
				return false
			}
		}
	}

	return true
}

func cropByIndex(set []int, index int) (prev, cur, next []int) {
	prev = make([]int, 0, len(set)-1)
	prev = append(prev, set[:index-1]...)
	prev = append(prev, set[index:]...)

	cur = make([]int, 0, len(set)-1)
	cur = append(cur, set[:index]...)
	cur = append(cur, set[index+1:]...)

	if index == len(set)-1 {
		return prev, cur, set[:index]
	}

	next = make([]int, 0, len(set)-1)
	next = append(next, set[:index+1]...)
	next = append(next, set[index+2:]...)

	return prev, cur, next
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
