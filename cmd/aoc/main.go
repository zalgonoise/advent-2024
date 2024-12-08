package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/zalgonoise/advent-2024/cmd/aoc/day01"
	"github.com/zalgonoise/advent-2024/cmd/aoc/day02"
	"github.com/zalgonoise/advent-2024/cmd/aoc/day03"
	"github.com/zalgonoise/advent-2024/cmd/aoc/day04"
	"github.com/zalgonoise/advent-2024/cmd/aoc/day05"
	"github.com/zalgonoise/advent-2024/cmd/aoc/day06"
)

var (
	errInvalidDay     = errors.New("invalid day")
	errUnsupportedDay = errors.New("unsupported day")
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	code, err := run(logger)
	if err != nil {
		logger.ErrorContext(context.Background(), "runtime error", slog.String("error", err.Error()))
		os.Exit(code)
	}

	os.Exit(code)
}

func run(logger *slog.Logger) (int, error) {
	fs := flag.NewFlagSet("main", flag.ExitOnError)

	day := fs.Int("day", 0, "the Advent of Code day to run")

	if err := fs.Parse(os.Args[1:3]); err != nil {
		return 1, err
	}

	if *day < 1 || *day > 31 {
		return 1, fmt.Errorf("%w: %d", errInvalidDay, *day)
	}

	ctx := context.Background()

	switch *day {
	case 1:
		return day01.Exec(ctx, logger, os.Args[3:])
	case 2:
		return day02.Exec(ctx, logger, os.Args[3:])
	case 3:
		return day03.Exec(ctx, logger, os.Args[3:])
	case 4:
		return day04.Exec(ctx, logger, os.Args[3:])
	case 5:
		return day05.Exec(ctx, logger, os.Args[3:])
	case 6:
		return day06.Exec(ctx, logger, os.Args[3:])
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedDay, *day)
	}
}
