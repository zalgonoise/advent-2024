package day06

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d6 "github.com/zalgonoise/advent-2024/day06"
	p1 "github.com/zalgonoise/advent-2024/day06/part01"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-06", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate this day's challenge. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d6.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 1),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		grid := p1.Parse(*input)

		res := p1.Move(grid)

		result := p1.SumResults(res)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:

		logger.WarnContext(ctx, "not yet implmented", attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
