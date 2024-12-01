package part01

import (
	"bytes"
	"errors"
	"slices"
	"strconv"
	"strings"
)

var (
	ErrInputEmpty         = errors.New("input is empty")
	ErrUnexpectedTableLen = errors.New("unexpected table length")
)

func ParseTable(input string) ([2][]int, error) {
	lines := strings.Split(input, "\n")

	if len(lines) == 0 {
		return [2][]int{}, ErrInputEmpty
	}

	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for i := range lines {
		items, err := parseTableItem(lines[i])
		if err != nil {
			return [2][]int{}, err
		}

		left = append(left, items[0])
		right = append(right, items[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	return [2][]int{left, right}, nil
}

func GetDifference(input [2][]int) []int {
	out := make([]int, 0, len(input[0]))

	for i := range input[0] {
		if input[0][i] > input[1][i] {
			out = append(out, input[0][i]-input[1][i])

			continue
		}

		out = append(out, input[1][i]-input[0][i])
	}

	return out
}

func Sum(input []int) int {
	var n int

	for i := range input {
		n += input[i]
	}

	return n
}

func parseTableItem(item string) ([]int, error) {
	items := make([]int, 0, 2)

	buf := &bytes.Buffer{}

	for i := range item {
		switch {
		case item[i] == ' ' && buf.Len() == 0:
			continue
		case item[i] == ' ':
			n, err := strconv.Atoi(buf.String())
			if err != nil {
				return nil, err
			}

			items = append(items, n)

			buf.Reset()
		default:
			buf.WriteByte(item[i])
		}
	}

	if buf.Len() > 0 {
		n, err := strconv.Atoi(buf.String())
		if err != nil {
			return nil, err
		}

		items = append(items, n)
	}

	if len(items) != 2 {
		return nil, ErrUnexpectedTableLen
	}

	return items, nil
}
