package part01

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

var (
	ErrNoBreakpoint         = errors.New("no breakpoint found")
	ErrInvalidNumPriorities = errors.New("invalid number of priorities")
)

func Parse(s string) (map[int][]int, [][]int, error) {
	sets, orders, ok := strings.Cut(s, "\n\n")
	if !ok {
		return nil, nil, ErrNoBreakpoint
	}

	setLines := strings.Split(sets, "\n")
	priorities := make(map[int][]int, len(setLines))

	for i := range setLines {
		ns := strings.Split(setLines[i], "|")

		if len(ns) != 2 {
			return nil, nil, ErrInvalidNumPriorities
		}

		left, err := strconv.Atoi(ns[0])
		if err != nil {
			return nil, nil, err
		}

		right, err := strconv.Atoi(ns[1])
		if err != nil {
			return nil, nil, err
		}

		if _, exists := priorities[left]; !exists {
			priorities[left] = make([]int, 0, 64)
		}

		priorities[left] = append(priorities[left], right)
	}

	orderLines := strings.Split(orders, "\n")
	instructions := make([][]int, 0, len(orderLines))

	for i := range orderLines {
		ns := strings.Split(orderLines[i], ",")
		instruction := make([]int, 0, len(ns))

		for i := range ns {
			value, err := strconv.Atoi(ns[i])
			if err != nil {
				return nil, nil, err
			}

			instruction = append(instruction, value)
		}

		instructions = append(instructions, instruction)
	}

	return priorities, instructions, nil
}

func Filter(priorities map[int][]int, instructions [][]int) [][]int {
	valid := make([][]int, 0, len(instructions))

	for i := range instructions {
		if !checkInstruction(instructions[i], priorities) {
			continue
		}

		valid = append(valid, instructions[i])
	}

	return valid
}

func SumMiddleValue(instructions [][]int) int {
	var n int

	for i := range instructions {
		n += instructions[i][len(instructions[i])/2]
	}

	return n
}

func checkInstruction(ns []int, priorities map[int][]int) bool {
	for idx := 0; idx < len(ns); idx++ {
		after := priorities[ns[idx]]

		if idx > 0 && containsPrev(ns[:idx], after) {
			return false
		}
	}

	return true
}

func containsPrev(prev []int, after []int) bool {
	for ip := range prev {
		if slices.Contains(after, prev[ip]) {
			return true
		}
	}

	return false
}
