package part01

import (
	"fmt"
	"github.com/zalgonoise/advent-2024/grid"
	"slices"
	"strings"
)

const (
	up = iota
	down
	left
	right
)

func Parse(input string) grid.Grid[byte] {
	lines := strings.Split(input, "\n")

	g := grid.Grid[byte]{
		Height: len(lines),
		Width:  len(lines[0]),
		Values: make([][]byte, 0, len(lines)),
	}

	for i := range lines {
		g.Values = append(g.Values, []byte(lines[i]))
	}

	return g
}

func printGrid(g grid.Grid[byte], start, cursor grid.Coord, dir int) {
	v := make([][]byte, 0, len(g.Values))
	for i := range g.Values {
		vals := make([]byte, len(g.Values[i]))
		copy(vals, g.Values[i])
		v = append(v, vals)
	}

	v[start.Y][start.X] = '.'

	switch dir {
	case up:
		v[cursor.Y][cursor.X] = '^'
	case down:
		v[cursor.Y][cursor.X] = 'v'
	case left:
		v[cursor.Y][cursor.X] = '<'
	case right:
		v[cursor.Y][cursor.X] = '>'
	}

	sb := &strings.Builder{}

	for i := range v {
		sb.Write(v[i])
		sb.WriteByte('\n')
	}

	fmt.Println(sb.String())
}

func Move(g grid.Grid[byte]) map[int][]int {
	starts := grid.FindAll(g, '^')
	if len(starts) != 1 {
		return nil
	}

	pos := starts[0]
	dir := up
	cache := make(map[int][]int) // map[Y][]Xs
	cache[pos.Y] = make([]int, 0, len(g.Values[0]))
	cache[pos.Y] = append(cache[pos.Y], pos.X)

	for {
		switch dir {
		case up:
			next := pos.Up()
			upValue, ok := g.Value(next)
			if !ok {
				return cache
			}

			if upValue == '#' {
				dir = right

				continue
			}

			pos = next
			if cache[pos.Y] == nil {
				cache[pos.Y] = make([]int, 0, len(g.Values[0]))
			}

			if !slices.Contains(cache[pos.Y], pos.X) {
				cache[pos.Y] = append(cache[pos.Y], pos.X)
			}

		case down:
			next := pos.Down()
			upValue, ok := g.Value(next)
			if !ok {
				return cache
			}

			if upValue == '#' {
				dir = left

				continue
			}

			pos = next
			if cache[pos.Y] == nil {
				cache[pos.Y] = make([]int, 0, len(g.Values[0]))
			}

			if !slices.Contains(cache[pos.Y], pos.X) {
				cache[pos.Y] = append(cache[pos.Y], pos.X)
			}

		case left:
			next := pos.Left()
			upValue, ok := g.Value(next)
			if !ok {
				return cache
			}

			if upValue == '#' {
				dir = up

				continue
			}

			pos = next
			if cache[pos.Y] == nil {
				cache[pos.Y] = make([]int, 0, len(g.Values[0]))
			}

			if !slices.Contains(cache[pos.Y], pos.X) {
				cache[pos.Y] = append(cache[pos.Y], pos.X)
			}
		case right:
			next := pos.Right()
			upValue, ok := g.Value(next)
			if !ok {
				return cache
			}

			if upValue == '#' {
				dir = down

				continue
			}

			pos = next
			if cache[pos.Y] == nil {
				cache[pos.Y] = make([]int, 0, len(g.Values[0]))
			}

			if !slices.Contains(cache[pos.Y], pos.X) {
				cache[pos.Y] = append(cache[pos.Y], pos.X)
			}
		}
	}
}

func SumResults(res map[int][]int) int {
	var n int

	for _, v := range res {
		n += len(v)
	}

	return n
}
