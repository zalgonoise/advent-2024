package part02

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Mul struct {
	a, b int
}

const (
	dont = "don't()"
	do   = "do()"
)

var (
	ErrInvalidPair = errors.New("invalid number of items in a pair")
)

var exp = regexp.MustCompile(`mul\(\d{1,3}?,\d{1,3}\)`)

func ParseValid(input string) []string {
	n := len(input)
	value := input

	valid := make([]string, 0, 64)

	for n > 0 {
		left, right, ok := strings.Cut(value, dont)
		if !ok {
			if left != "" {
				valid = append(valid, left)
			}

			break
		}

		valid = append(valid, left)
		left, right, ok = strings.Cut(right, do)
		if !ok {
			break
		}

		value = right
	}

	return valid
}

func ParseRegexp(input []string) ([]Mul, error) {
	muls := make([]Mul, 0, len(input)*8)

	for i := range input {
		results := exp.FindAllString(input[i], -1)

		for i := range results {
			value := results[i]
			// crop off mul, parens and comma
			_, value, _ = strings.Cut(value, "mul(")
			value, _, _ = strings.Cut(value, ")")

			numbers := strings.Split(value, ",")

			if len(numbers) != 2 {
				return nil, ErrInvalidPair
			}

			left, err := strconv.Atoi(numbers[0])
			if err != nil {
				return nil, err
			}

			right, err := strconv.Atoi(numbers[1])
			if err != nil {
				return nil, err
			}

			muls = append(muls, Mul{left, right})
		}
	}

	return muls, nil
}

func Sum(muls []Mul) int {
	var n int

	for i := range muls {
		n += muls[i].a * muls[i].b
	}

	return n
}
