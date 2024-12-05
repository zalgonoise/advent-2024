package part01

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidPair = errors.New("invalid number of items in a pair")
)

var exp = regexp.MustCompile(`mul\(\d{1,3}?,\d{1,3}\)`)

func ParseRegexp(input string) ([]Mul, error) {
	results := exp.FindAllString(input, -1)
	muls := make([]Mul, 0, len(results))

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

	return muls, nil
}
