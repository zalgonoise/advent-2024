package day02

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d2 "github.com/zalgonoise/advent-2024/day02"
	p1 "github.com/zalgonoise/advent-2024/day02/part01"
	p2 "github.com/zalgonoise/advent-2024/day02/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-02", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate this day's challenge. An empty string uses the generated input")
	thresh := fs.Int("thresh", 3, "threshold")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d2.Input
	}

	if *thresh <= 0 {
		*thresh = 3
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 1),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		table, err := p1.ParseTable(*input)
		if err != nil {
			logger.ErrorContext(ctx, "execution failed", slog.String("error", err.Error()), attr)
		}

		result := len(p1.Filter(table, *thresh))

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		table, err := p2.ParseTable(*input)
		if err != nil {
			logger.ErrorContext(ctx, "execution failed", slog.String("error", err.Error()), attr)
		}

		result := len(p2.Filter2(table, *thresh))

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
