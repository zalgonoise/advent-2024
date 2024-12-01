package part02

import (
	"bytes"
	"errors"
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

	return [2][]int{left, right}, nil
}

func CalculateOccurrences(input [2][]int) []int {
	counts := make([]int, 0, len(input))

	for i := range input[0] {
		var count int

		for idx := range input[1] {
			if input[0][i] != input[1][idx] {
				continue
			}

			count++
		}

		counts = append(counts, input[0][i]*count)
	}

	return counts
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
